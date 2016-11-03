package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/dsvensson/go-xmmsclient/xmmsclient"
	"io"
	"log"
	"net"
)

type XmmsMessage struct {
	objectId  uint32
	commandId uint32
	args      xmmsclient.XmmsValue
	cookie    uint32
}

type XmmsReply struct {
	cookie uint32
	value  xmmsclient.XmmsValue
}

func send_hello(registry chan XmmsContext, outqueue chan XmmsMessage, client_name string) xmmsclient.XmmsValue {
	var cookie = uint32(0x31336)
	var ctx = XmmsContext{make(chan xmmsclient.XmmsValue), cookie}
	registry <- ctx
	outqueue <- XmmsMessage{
		objectId:  xmmsclient.ObjectMain,
		commandId: xmmsclient.CommandMainHello,
		args: xmmsclient.NewXmmsList(
			xmmsclient.XmmsInt(xmmsclient.IpcVersion),
			xmmsclient.XmmsString(client_name)),
		cookie: cookie}
	return <-ctx.result
}

func send_list_plugins(registry chan XmmsContext, outqueue chan XmmsMessage) xmmsclient.XmmsValue {
	var cookie = uint32(0x31337)
	var ctx = XmmsContext{make(chan xmmsclient.XmmsValue), cookie}
	registry <- ctx
	outqueue <- XmmsMessage{
		objectId:  xmmsclient.ObjectMain,
		commandId: xmmsclient.CommandMainListPlugins,
		args:      xmmsclient.NewXmmsList(xmmsclient.XmmsInt(0)),
		cookie:    cookie}
	return <-ctx.result
}

func connect(dest string) (conn *net.TCPConn, err error) {
	var addr, _ = net.ResolveTCPAddr("tcp", dest)
	conn, err = net.DialTCP("tcp", nil, addr)
	return
}

type XmmsContext struct {
	result chan xmmsclient.XmmsValue
	cookie uint32
}

func main() {
	var conn, err = connect("localhost:xmms2")
	if err != nil {
		log.Fatal("Could not connect")
		return
	}

	var registryChan = make(chan XmmsContext)
	var inqueue = make(chan XmmsReply)
	var outqueue = make(chan XmmsMessage)

	// inbound
	go func() {
		var header = make([]byte, 16)
		for {
			fmt.Println("Starting to read stuff")
			var n, err = io.ReadFull(conn, header)
			if err != nil {
				fmt.Println("error reading socket")
				continue
			}
			if n == 0 {
				fmt.Println("nothing to read")
				continue
			}

			var headerBuffer = bytes.NewBuffer(header)
			var objectId uint32
			var commandId uint32
			var cookie uint32
			var length uint32

			binary.Read(headerBuffer, binary.BigEndian, &objectId)
			binary.Read(headerBuffer, binary.BigEndian, &commandId)
			binary.Read(headerBuffer, binary.BigEndian, &cookie)
			binary.Read(headerBuffer, binary.BigEndian, &length)

			fmt.Printf("Got => obj=%d, cmd=%d, cookie=0x%x, len=%d\n", objectId, commandId, cookie, length)

			var payload = make([]byte, length)
			conn.Read(payload)
			fmt.Println(hex.Dump(payload))
			var buffer = bytes.NewBuffer(payload)
			var value, _ = xmmsclient.DeserializeXmmsValue(buffer)
			fmt.Println(value)
			inqueue <- XmmsReply{cookie, value}
		}
	}()

	// result registry
	go func() {
		var registry = make(map[uint32](XmmsContext))

		for {
			select {
			case ctx := <-registryChan:
				fmt.Println("Registered context")
				fmt.Println(ctx)
				registry[ctx.cookie] = ctx
			case reply := <-inqueue:
				registry[reply.cookie].result <- reply.value
			}
		}
	}()

	// outbound
	go func() {
		for {
			fmt.Println("here we go again", len(outqueue))
			select {
			case msg := <-outqueue:
				fmt.Println("Sending message", msg)
				binary.Write(conn, binary.BigEndian, msg.objectId)
				binary.Write(conn, binary.BigEndian, msg.commandId)
				binary.Write(conn, binary.BigEndian, msg.cookie)

				var payload bytes.Buffer
				xmmsclient.SerializeXmmsValue(msg.args, &payload)
				binary.Write(conn, binary.BigEndian, uint32(len(payload.Bytes())))
				payload.WriteTo(conn)
				fmt.Println("done sending")
			}
			fmt.Println("here we go")
		}

	}()

	var result = send_hello(registryChan, outqueue, "hello-from-go")
	fmt.Println("Got result", result)

	var list = send_list_plugins(registryChan, outqueue)
	fmt.Println(list)

	select {}
}
