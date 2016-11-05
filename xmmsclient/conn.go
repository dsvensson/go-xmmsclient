package xmmsclient

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync/atomic"
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
	args      XmmsValue
	context   context
}

type reply struct {
	sequenceNr uint32
	value      XmmsValue
}

type Client struct {
	conn       io.ReadWriteCloser
	sequenceNr uint32
	clientId   XmmsValue // TODO: uint32

	inbound  chan reply
	outbound chan message
	registry chan context
}

func parseHeader(buf *bytes.Buffer) (hdr header, err error) {
	err = binary.Read(buf, binary.BigEndian, &hdr.objectId)
	if err != nil {
		return
	}

	err = binary.Read(buf, binary.BigEndian, &hdr.commandId)
	if err != nil {
		return
	}

	err = binary.Read(buf, binary.BigEndian, &hdr.sequenceNr)
	if err != nil {
		return
	}

	err = binary.Read(buf, binary.BigEndian, &hdr.length)
	if err != nil {
		return
	}

	return
}

func writeHeader(w io.ReadWriteCloser, hdr *header) (err error) {
	err = binary.Write(w, binary.BigEndian, hdr.objectId)
	if err != nil {
		return
	}

	err = binary.Write(w, binary.BigEndian, hdr.commandId)
	if err != nil {
		return
	}

	err = binary.Write(w, binary.BigEndian, hdr.sequenceNr)
	if err != nil {
		return
	}

	err = binary.Write(w, binary.BigEndian, hdr.length)
	if err != nil {
		return
	}

	return
}

func (c *Client) nextSequenceNr() (sequenceNr uint32) {
	sequenceNr = atomic.AddUint32(&c.sequenceNr, 1)
	return
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
	for {
		select {
		case msg := <-c.outbound:
			var payload bytes.Buffer

			c.registry <- msg.context

			err := SerializeXmmsValue(msg.args, &payload)
			if err != nil {
				continue
			}

			header := header{
				objectId:   msg.objectId,
				commandId:  msg.commandId,
				sequenceNr: msg.context.sequenceNr,
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

// TODO: Needs a better home, should be generated.
func (c *Client) MainHello(client_name string) XmmsValue {
	context := context{
		result:     make(chan XmmsValue),
		sequenceNr: c.nextSequenceNr()}
	c.outbound <- message{
		objectId:  ObjectMain,
		commandId: CommandMainHello,
		context:   context,
		args:      NewXmmsList(XmmsInt(IpcVersion), XmmsString(client_name))}
	return <-context.result
}

// TODO: Needs a better home, should be generated.
func (c *Client) MainListPlugins() XmmsValue {
	context := context{
		result:     make(chan XmmsValue),
		sequenceNr: c.nextSequenceNr()}
	c.outbound <- message{
		args:      NewXmmsList(XmmsInt(0)),
		objectId:  ObjectMain,
		commandId: CommandMainListPlugins,
		context:   context}
	return <-context.result
}

// TODO: Probably something else that creates a new Client rather than Dial.
func Dial(url string, name string) (client Client, err error) {
	addr, err := net.ResolveTCPAddr("tcp", url)
	if err != nil {
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return
	}

	client = Client{
		conn:     conn,
		inbound:  make(chan reply),
		outbound: make(chan message),
		registry: make(chan context),
	}

	go client.router()
	go client.reader()
	go client.writer()

	client.clientId = client.MainHello(name)

	return
}
