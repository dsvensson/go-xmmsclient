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

func parseHeader(header *bytes.Buffer) (objId uint32, cmdId uint32, sequenceNr uint32, length uint32, err error) {
	err = binary.Read(header, binary.BigEndian, &objId)
	if err != nil {
		return
	}

	err = binary.Read(header, binary.BigEndian, &cmdId)
	if err != nil {
		return
	}

	err = binary.Read(header, binary.BigEndian, &sequenceNr)
	if err != nil {
		return
	}

	err = binary.Read(header, binary.BigEndian, &length)
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
	var header = make([]byte, 16)
	for {
		read, err := io.ReadFull(c.conn, header)
		if err != nil {
			// TODO: Post to error channel
			fmt.Println("error reading socket")
			continue
		}

		if read != len(header) {
			fmt.Println("nothing to read")
			continue
		}

		_, _, sequenceNr, length, err := parseHeader(bytes.NewBuffer(header))

		payload := make([]byte, length)
		c.conn.Read(payload)
		value, _ := DeserializeXmmsValue(bytes.NewBuffer(payload))
		c.inbound <- reply{sequenceNr, value}
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
			binary.Write(c.conn, binary.BigEndian, msg.objectId)
			if err != nil {
				continue
			}
			binary.Write(c.conn, binary.BigEndian, msg.commandId)
			if err != nil {
				continue
			}
			binary.Write(c.conn, binary.BigEndian, msg.context.sequenceNr)
			if err != nil {
				continue
			}
			binary.Write(c.conn, binary.BigEndian, uint32(len(payload.Bytes())))
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
			registry[reply.sequenceNr].result <- reply.value
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
func Dial(url string) (client Client, err error) {
	addr, err := net.ResolveTCPAddr("tcp", url)
	if err != nil {
		return
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return
	}

	client = Client{
		conn: conn,
	}

	client.inbound = make(chan reply)
	client.outbound = make(chan message)
	client.registry = make(chan context)

	go client.router()
	go client.reader()
	go client.writer()

	client.clientId = client.MainHello("hello-from-go")

	return
}
