package xmmsclient

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"sync"
)

type context struct {
	result    chan reply
	broadcast bool
}

type header struct {
	objectId   uint32
	commandId  uint32
	sequenceNr uint32
	length     uint32
}

type message struct {
	header    header
	broadcast bool
	args      XmmsValue
	result    chan reply
}

type reply struct {
	sequenceNr uint32
	payload    []byte
	err        error
}

type Client struct {
	sync.RWMutex

	sequenceNr uint32
	clientName string

	shutdownRegistry chan bool
	shutdownIO       chan bool
	registry         chan message
}

func parseHeader(buf *bytes.Buffer) (*header, error) {
	var hdr header

	err := binary.Read(buf, binary.BigEndian, &hdr.objectId)
	if err != nil {
		return nil, err
	}

	err = binary.Read(buf, binary.BigEndian, &hdr.commandId)
	if err != nil {
		return nil, err
	}

	err = binary.Read(buf, binary.BigEndian, &hdr.sequenceNr)
	if err != nil {
		return nil, err
	}

	err = binary.Read(buf, binary.BigEndian, &hdr.length)
	if err != nil {
		return nil, err
	}

	return &hdr, nil
}

func writeHeader(w io.Writer, hdr *header) error {
	err := binary.Write(w, binary.BigEndian, hdr.objectId)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, hdr.commandId)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, hdr.sequenceNr)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, hdr.length)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) nextSequenceNr() uint32 {
	c.sequenceNr += 1
	return c.sequenceNr
}

func (c *Client) reader(r io.Reader, inbound chan reply) {
	buffer := make([]byte, 16)

	for {
		_, err := io.ReadFull(r, buffer)
		if err != nil {
			inbound <- reply{err: err}
			break
		}

		header, err := parseHeader(bytes.NewBuffer(buffer))
		if err != nil {
			inbound <- reply{err: err}
			break
		}

		payload := make([]byte, header.length)

		_, err = io.ReadFull(r, payload)
		if err != nil {
			inbound <- reply{err: err}
			break
		}

		inbound <- reply{sequenceNr: header.sequenceNr, payload: payload}
	}
}

func (c *Client) writer(w io.Writer, outbound chan message, errors chan error) {
writer:
	for {
		select {
		case msg := <-outbound:
			var payload bytes.Buffer

			err := serializeXmmsValue(&payload, msg.args)
			if err != nil {
				errors <- err
				break writer
			}

			msg.header.length = uint32(len(payload.Bytes()))

			err = writeHeader(w, &msg.header)
			if err != nil {
				errors <- err
				break writer
			}

			payload.WriteTo(w)
			if err != nil {
				errors <- err
				break writer
			}
		case <-c.shutdownIO:
			break writer
		}
	}
}

func (c *Client) shutdownRouter(registry map[uint32](context), err error) {
	// Reference command channel as it will be nullified
	channel := c.registry

	// Grab the RW-lock, close and nullify the command channel
	// reference which will allow the draining loop to exit.
	go func() {
		c.Lock()
		c.registry = nil
		close(channel)
		c.Unlock()
	}()

	// Terminate all active subscriptions
	for _, v := range registry {
		v.result <- reply{err: err}
	}

	// Drain trailing requests
	for msg := range channel {
		msg.result <- reply{err: io.EOF}
	}
}

func (c *Client) router(inbound chan reply, outbound chan message, errors chan error) {
	registry := make(map[uint32](context))

	for {
		select {
		case msg := <-c.registry:
			msg.header.sequenceNr = c.nextSequenceNr()
			registry[msg.header.sequenceNr] = context{
				msg.result,
				msg.broadcast,
			}
			outbound <- msg
		case reply := <-inbound:
			if reply.err != nil {
				c.shutdownRouter(registry, reply.err)
				return
			}

			ctx := registry[reply.sequenceNr]

			go func() {
				ctx.result <- reply
			}()

			if !ctx.broadcast {
				delete(registry, reply.sequenceNr)
			}
		case err := <-errors:
			c.shutdownRouter(registry, err)
			return
		case <-c.shutdownRegistry:
			c.shutdownRouter(registry, io.EOF)
			return
		}
	}
}

func (c *Client) dispatch(objectId uint32, commandId uint32, args XmmsValue) chan reply {
	c.RLock()
	defer c.RUnlock()

	result := make(chan reply, 1)
	if c.registry == nil {
		result <- reply{err: io.EOF}
	} else {
		c.registry <- message{
			header: header{
				objectId:  objectId,
				commandId: commandId,
			},
			broadcast: objectId == 0,
			args:      args,
			result:    result,
		}
	}
	return result
}

func (c *Client) sendHello() (int, error) {
	result := make(chan reply)

	c.registry <- message{
		header: header{
			objectId:  1,
			commandId: 32,
		},
		broadcast: false,
		args:      XmmsList{XmmsInt(IpcVersion), XmmsString(c.clientName)},
		result:    result,
	}

	reply := <-result
	if reply.err != nil {
		return -1, reply.err
	}

	buffer := bytes.NewBuffer(reply.payload)

	value, err := tryDeserialize(buffer)
	if err != nil {
		return -1, err
	}

	clientId, ok := value.(XmmsInt)
	if !ok {
		return -1, errors.New("Bad reply from server")
	}

	return int(clientId), nil
}

func (c *Client) Dial(url string) (int, error) {
	c.Lock()
	defer c.Unlock()

	addr, err := net.ResolveTCPAddr("tcp", url)
	if err != nil {
		return -1, err
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return -1, err
	}

	c.shutdownRegistry = make(chan bool)
	c.shutdownIO = make(chan bool)
	c.registry = make(chan message)

	errors := make(chan error)
	inbound := make(chan reply)
	outbound := make(chan message)

	go c.reader(conn, inbound)
	go c.writer(conn, outbound, errors)
	go c.router(inbound, outbound, errors)

	clientId, err := c.sendHello()
	if err != nil {
		c.Close()
		return -1, err
	}

	return clientId, nil
}

func (c *Client) Close() {
	c.shutdownRegistry <- true
	c.shutdownIO <- true
}

func NewClient(name string) *Client {
	client := Client{
		clientName: name,
	}

	return &client
}
