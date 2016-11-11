package xmmsclient

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"net"
)

type result struct {
	value XmmsValue
	err   error
}

type context struct {
	result     chan result
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
	result    chan result
}

type reply struct {
	sequenceNr uint32
	value      XmmsValue
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

func (c *Client) reader(conn *net.TCPConn, inbound chan reply) {
	var buffer = make([]byte, 16)

	for {
		_, err := io.ReadFull(conn, buffer)
		if err != nil {
			return
		}

		header, err := parseHeader(bytes.NewBuffer(buffer))

		payload := make([]byte, header.length)

		_, err = io.ReadFull(conn, payload)
		if err != nil {
			return
		}

		value, _ := DeserializeXmmsValue(bytes.NewBuffer(payload))

		inbound <- reply{header.sequenceNr, value}
	}
}

func (c *Client) writer(conn *net.TCPConn, outbound chan message) {
writer:
	for {
		select {
		case msg := <-outbound:
			var payload bytes.Buffer

			err := SerializeXmmsValue(msg.args, &payload)
			if err != nil {
				break writer
			}

			msg.header.length = uint32(len(payload.Bytes()))

			err = writeHeader(conn, &msg.header)
			if err != nil {
				break writer
			}

			payload.WriteTo(conn)
			if err != nil {
				break writer
			}
		case <-c.shutdownIO:
			return
		}
	}

	// TODO: Probably a better place for this.
	conn.Close()
}

func (c *Client) router(inbound chan reply, outbound chan message) {
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
			if error, ok := reply.value.(XmmsError); ok {
				ctx.result <- result{nil, errors.New(string(error))}
			} else {
				ctx.result <- result{reply.value, nil}
			}
			if !ctx.broadcast {
				delete(registry, ctx.sequenceNr)
			}
		case <-c.shutdownRegistry:
			break router
		}
	}

	for _, v := range registry {
		v.result <- result{nil, io.EOF}
	}
}

func (c *Client) dispatch(objectId uint32, commandId uint32, args XmmsValue) chan result {
	var result = make(chan result)
	c.registry <- message{
		header: header{
			objectId:  objectId,
			commandId: commandId,
		},
		broadcast: false,
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

	inbound := make(chan reply)
	outbound := make(chan message)

	go c.reader(conn, inbound)
	go c.writer(conn, outbound)
	go c.router(inbound, outbound)

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
