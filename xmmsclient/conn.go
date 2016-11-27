package xmmsclient

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
)

type context struct {
	result     chan []byte
	sequenceNr uint32
	broadcast  bool
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
	result    chan []byte
}

type reply struct {
	sequenceNr uint32
	payload    []byte
}

type Client struct {
	sequenceNr uint32
	clientName string
	clientId   int

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

func writeHeader(w io.ReadWriteCloser, hdr *header) error {
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

func (c *Client) reader(conn *net.TCPConn, inbound chan reply, errors chan error) {
	var buffer = make([]byte, 16)

	for {
		_, err := io.ReadFull(conn, buffer)
		if err != nil {
			errors <- err
			break
		}

		header, err := parseHeader(bytes.NewBuffer(buffer))
		if err != nil {
			errors <- err
			break
		}

		payload := make([]byte, header.length)

		_, err = io.ReadFull(conn, payload)
		if err != nil {
			errors <- err
			break
		}

		inbound <- reply{header.sequenceNr, payload}
	}
}

func (c *Client) writer(conn *net.TCPConn, outbound chan message, errors chan error) {
writer:
	for {
		select {
		case msg := <-outbound:
			var payload bytes.Buffer

			err := serializeXmmsValue(msg.args, &payload)
			if err != nil {
				errors <- err
				break writer
			}

			msg.header.length = uint32(len(payload.Bytes()))

			err = writeHeader(conn, &msg.header)
			if err != nil {
				errors <- err
				break writer
			}

			payload.WriteTo(conn)
			if err != nil {
				errors <- err
				break writer
			}
		case <-c.shutdownIO:
			break writer
		}
	}

	// TODO: Probably a better place for this.
	conn.Close()
}

func errorToBytes(err string) []byte {
	var payload bytes.Buffer
	serializeXmmsValue(XmmsError(err), &payload)
	return payload.Bytes()
}

func (c *Client) router(inbound chan reply, outbound chan message, errors chan error) {
	var registry = make(map[uint32](context))

router:
	for {
		select {
		case msg := <-c.registry:
			msg.header.sequenceNr = c.nextSequenceNr()
			registry[msg.header.sequenceNr] = context{
				msg.result,
				msg.header.sequenceNr,
				msg.broadcast,
			}
			outbound <- msg
		case reply := <-inbound:
			ctx := registry[reply.sequenceNr]
			go func() {
				ctx.result <- reply.payload
			}()
			if !ctx.broadcast {
				delete(registry, ctx.sequenceNr)
			}
		case err := <-errors:
			payload := errorToBytes(err.Error())
			for _, v := range registry {
				v.result <- payload
			}
			break router
		case <-c.shutdownRegistry:
			payload := errorToBytes(io.EOF.Error())
			for _, v := range registry {
				v.result <- payload
			}
			break router
		}
	}
}

func (c *Client) dispatch(objectId uint32, commandId uint32, args XmmsValue) chan []byte {
	result := make(chan []byte)
	c.registry <- message{
		header: header{
			objectId:  objectId,
			commandId: commandId,
		},
		broadcast: objectId == 0,
		args:      args,
		result:    result,
	}
	return result
}

func (c *Client) Dial(url string) error {
	addr, err := net.ResolveTCPAddr("tcp", url)
	if err != nil {
		return err
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return err
	}

	c.shutdownRegistry = make(chan bool)
	c.shutdownIO = make(chan bool)
	c.registry = make(chan message)

	errors := make(chan error)
	inbound := make(chan reply)
	outbound := make(chan message)

	go c.reader(conn, inbound, errors)
	go c.writer(conn, outbound, errors)
	go c.router(inbound, outbound, errors)

	clientId, err := c.MainHello(24, c.clientName)
	if err != nil {
		return err
	}

	c.clientId = int(clientId)

	return nil
}

func (c *Client) ClientId() (int, error) {
	if c.clientId == -1 {
		return -1, errors.New("Client Id not initialized.")
	}
	return c.clientId, nil
}

func (c *Client) Close() {
	c.shutdownRegistry <- true
	c.shutdownIO <- true
	c.clientId = -1
}

func NewClient(name string) *Client {
	client := Client{
		clientId:   -1,
		clientName: name,
	}

	return &client
}
