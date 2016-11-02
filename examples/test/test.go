package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/dsvensson/go-xmmsclient/xmmsclient"
	"net"
)

type XmmsMessage struct {
	objectId  uint32
	commandId uint32
	args      xmmsclient.XmmsValue
}

func send_message(conn *net.TCPConn, message XmmsMessage) {
	var cookie = uint32(0x31337)

	binary.Write(conn, binary.BigEndian, message.objectId)
	binary.Write(conn, binary.BigEndian, message.commandId)
	binary.Write(conn, binary.BigEndian, cookie)

	var payload bytes.Buffer
	xmmsclient.SerializeXmmsValue(message.args, &payload)
	binary.Write(conn, binary.BigEndian, uint32(len(payload.Bytes())))
	payload.WriteTo(conn)
}

func recv_message(conn *net.TCPConn) *xmmsclient.XmmsValue {
	var objectId uint32
	var commandId uint32
	var cookie uint32
	var length uint32
	binary.Read(conn, binary.BigEndian, &objectId)
	binary.Read(conn, binary.BigEndian, &commandId)
	binary.Read(conn, binary.BigEndian, &cookie)
	binary.Read(conn, binary.BigEndian, &length)

	fmt.Printf("obj=%d, cmd=%d, cookie=0x%x, len=%d\n", objectId, commandId, cookie, length)

	var payload = make([]byte, length)
	conn.Read(payload)
	fmt.Println(hex.Dump(payload))
	var buffer = bytes.NewBuffer(payload)
	var value, _ = xmmsclient.DeserializeXmmsValue(buffer)
	return &value
}

func send_hello(conn *net.TCPConn, client_name string) *xmmsclient.XmmsValue {
	send_message(conn, XmmsMessage{
		objectId:  xmmsclient.ObjectMain,
		commandId: xmmsclient.CommandMainHello,
		args: xmmsclient.NewXmmsList(
			xmmsclient.XmmsInt(xmmsclient.IpcVersion),
			xmmsclient.XmmsString(client_name))})
	return recv_message(conn)
}

func send_list_plugins(conn *net.TCPConn, client_name string) *xmmsclient.XmmsValue {
	send_message(conn, XmmsMessage{
		objectId:  xmmsclient.ObjectMain,
		commandId: xmmsclient.CommandMainListPlugins,
		args:      xmmsclient.NewXmmsList(xmmsclient.XmmsInt(0))})
	return recv_message(conn)
}

func main() {
	var addr, _ = net.ResolveTCPAddr("tcp", "localhost:xmms2")
	var conn, _ = net.DialTCP("tcp", nil, addr)
	send_hello(conn, "hello-from-go")

	var list = send_list_plugins(conn, "hello-from-go")
	fmt.Println(*list)
}
