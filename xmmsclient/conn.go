package xmmsclient

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"net"
)

type context struct {
	result     chan XmmsValue
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
	objectId  uint32
	commandId uint32
	broadcast bool
	args      XmmsValue
	result    chan XmmsValue
}

type reply struct {
	sequenceNr uint32
	value      XmmsValue
}

type Client struct {
	conn       io.ReadWriteCloser
	sequenceNr uint32
	clientName string
	clientId   int64

	inbound  chan reply
	outbound chan message
	registry chan context
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

func (c *Client) reader() {
	var buffer = make([]byte, 16)
	for {
		read, err := io.ReadFull(c.conn, buffer)
		if err != nil {
			// TODO: Post to error channel
			fmt.Println("error reading socket")
			continue
		}

		if read != len(buffer) {
			fmt.Println("nothing to read")
			continue
		}

		header, err := parseHeader(bytes.NewBuffer(buffer))

		payload := make([]byte, header.length)
		c.conn.Read(payload)
		value, _ := DeserializeXmmsValue(bytes.NewBuffer(payload))
		c.inbound <- reply{header.sequenceNr, value}
	}
}

func (c *Client) writer() {
	for msg := range c.outbound {
		var payload bytes.Buffer

		sequenceNr := c.nextSequenceNr()
		c.registry <- context{msg.result, sequenceNr, msg.broadcast}

		err := SerializeXmmsValue(msg.args, &payload)
		if err != nil {
			continue
		}

		header := header{
			objectId:   msg.objectId,
			commandId:  msg.commandId,
			sequenceNr: sequenceNr,
			length:     uint32(len(payload.Bytes())),
		}

		err = writeHeader(c.conn, &header)
		if err != nil {
			continue
		}

		payload.WriteTo(c.conn)
		if err != nil {
			continue
		}
	}
}

func (c *Client) router() {
	var registry = make(map[uint32](context))

	for {
		select {
		case ctx := <-c.registry:
			registry[ctx.sequenceNr] = ctx
		case reply := <-c.inbound:
			ctx := registry[reply.sequenceNr]
			ctx.result <- reply.value
			if !ctx.broadcast {
				delete(registry, ctx.sequenceNr)
			}
		}
	}
}

func (c *Client) dispatch(objectId uint32, commandId uint32, args XmmsValue) chan XmmsValue {
	var result = make(chan XmmsValue)
	c.outbound <- message{
		objectId:  objectId,
		commandId: commandId,
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

	c.conn, err = net.DialTCP("tcp", nil, addr)
	if err != nil {
		return err
	}

	c.inbound = make(chan reply)
	c.outbound = make(chan message)
	c.registry = make(chan context)

	go c.router()
	go c.reader()
	go c.writer()

	c.clientId = int64(c.MainHello(24, c.clientName).(XmmsInt)) // TODO: err

	return nil
}

func (c *Client) ClientId() (int64, error) {
	if c.clientId == -1 {
		return -1, errors.New("Client Id not initialized.")
	}
	return int64(c.clientId), nil
}

func NewClient(name string) *Client {
	client := Client{
		clientId:   -1,
		clientName: name,
	}

	return &client
}
