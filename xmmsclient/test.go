package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"net"
)

const (
	IpcVersion int64 = 24
)

const (
	ObjectSignal   uint32 = 0
	ObjectMain     uint32 = 1
	ObjectPlaylist uint32 = 2
	ObjectConfig   uint32 = 3
	ObjectPlayback uint32 = 4
	ObjectMedialib uint32 = 5
	/* ... */
)
const (
	CommandMainHello       uint32 = 32
	CommandMainQuit        uint32 = 33
	CommandMainListPlugins uint32 = 34
	CommandMainStats       uint32 = 35
	/* ... */
)

const (
	TypeNone      uint32 = 0
	TypeError     uint32 = 1
	TypeInt64     uint32 = 2
	TypeString    uint32 = 3
	TypeColl      uint32 = 4
	TypeBin       uint32 = 5
	TypeList      uint32 = 6
	TypeDict      uint32 = 7
	TypeBitBuffer uint32 = 8
	TypeFloat     uint32 = 9
)

type XmmsValue interface {
	isXmmsValue()
	serialize(buffer *bytes.Buffer)
}

type XmmsInt int64

func (i XmmsInt) isXmmsValue() {}
func (i XmmsInt) serialize(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, i)
}

type XmmsString string

func (s XmmsString) isXmmsValue() {}
func (s XmmsString) serialize(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, uint32(len(s)+1))
	binary.Write(buffer, binary.BigEndian, []byte(s))
	binary.Write(buffer, binary.BigEndian, byte(0))
}

type XmmsList struct {
	entries  []XmmsValue
	restrict uint32
}

func NewXmmsList(entries ...XmmsValue) XmmsList {
	return XmmsList{entries: entries, restrict: TypeNone}
}
func (l XmmsList) isXmmsValue() {}
func (l XmmsList) serialize(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, l.restrict)
	binary.Write(buffer, binary.BigEndian, uint32(len(l.entries)))
	if l.restrict != TypeNone {
		for _, entry := range l.entries {
			entry.serialize(buffer)
		}
	} else {
		for _, entry := range l.entries {
			SerializeXmmsValue(entry, buffer)
		}
	}
}

type XmmsDict map[string]XmmsValue

func (d XmmsDict) isXmmsValue() {}
func (d XmmsDict) serialize(buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, uint32(len(d)))
	for k, v := range d {
		XmmsString(k).serialize(buffer)
		SerializeXmmsValue(v, buffer)
	}
}

func SerializeXmmsValue(value XmmsValue, buffer *bytes.Buffer) {
	switch value.(type) {
	case XmmsInt:
		binary.Write(buffer, binary.BigEndian, TypeInt64)
		value.(XmmsInt).serialize(buffer)
	case XmmsString:
		binary.Write(buffer, binary.BigEndian, TypeString)
		value.(XmmsString).serialize(buffer)
	case XmmsDict:
		binary.Write(buffer, binary.BigEndian, TypeDict)
		value.(XmmsDict).serialize(buffer)
	case XmmsList:
		binary.Write(buffer, binary.BigEndian, TypeList)
		value.(XmmsList).serialize(buffer)
	}
}

func DeserializeXmmsValue(buffer *bytes.Buffer) (result XmmsValue, err error) {
	var valueType uint32
	binary.Read(buffer, binary.BigEndian, &valueType)
	switch valueType {
	case TypeInt64:
		var value XmmsInt
		binary.Read(buffer, binary.BigEndian, &value)
		result = value
	case TypeString:
		var length uint32
		binary.Read(buffer, binary.BigEndian, &length)
		var data = make([]byte, length)
		binary.Read(buffer, binary.BigEndian, &data)
		result = XmmsString(data)
	case TypeList:
		var restrict uint32
		binary.Read(buffer, binary.BigEndian, &restrict) // TODO: respect restrict
		var length uint32
		binary.Read(buffer, binary.BigEndian, &length)
		var value = NewXmmsList()
		for i := uint32(0); i < length; i++ {
			var entry, _ = DeserializeXmmsValue(buffer)
			value.entries = append(value.entries, entry)
		}
		result = value
	case TypeDict:
		var length uint32
		binary.Read(buffer, binary.BigEndian, &length)
		var value = XmmsDict{}
		for i := uint32(0); i < length; i++ {
			var strlength uint32
			binary.Read(buffer, binary.BigEndian, &strlength)
			var data = make([]byte, strlength)
			binary.Read(buffer, binary.BigEndian, &data)
			var eval, _ = DeserializeXmmsValue(buffer)
			value[string(data)] = eval
		}
		result = value
	}
	return
}

type XmmsMessage struct {
	objectId  uint32
	commandId uint32
	args      XmmsValue
}

func send_message(conn *net.TCPConn, message XmmsMessage) {
	var cookie = uint32(0x31337)

	binary.Write(conn, binary.BigEndian, message.objectId)
	binary.Write(conn, binary.BigEndian, message.commandId)
	binary.Write(conn, binary.BigEndian, cookie)

	var payload bytes.Buffer
	SerializeXmmsValue(message.args, &payload)
	binary.Write(conn, binary.BigEndian, uint32(len(payload.Bytes())))
	payload.WriteTo(conn)
}

func recv_message(conn *net.TCPConn) {
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
	var value, _ = DeserializeXmmsValue(buffer)
	fmt.Println(value)
}

func send_hello(conn *net.TCPConn, client_name string) {
	send_message(conn, XmmsMessage{
		objectId:  ObjectMain,
		commandId: CommandMainHello,
		args:      NewXmmsList(XmmsInt(IpcVersion), XmmsString(client_name)),
	})
	recv_message(conn)
}

func send_list_plugins(conn *net.TCPConn, client_name string) {
	send_message(conn, XmmsMessage{
		objectId:  ObjectMain,
		commandId: CommandMainListPlugins,
		args:      NewXmmsList(XmmsInt(0)),
	})
	recv_message(conn)
}

func main() {
	var addr, _ = net.ResolveTCPAddr("tcp", "localhost:xmms2")
	var conn, _ = net.DialTCP("tcp", nil, addr)
	send_hello(conn, "hello-from-go")
	send_list_plugins(conn, "hello-from-go")
}
