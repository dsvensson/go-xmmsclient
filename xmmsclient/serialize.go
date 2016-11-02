package xmmsclient

import (
	"bytes"
	"encoding/binary"
)

func serializeInt(i XmmsInt, buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, i)
}

func serializeString(s XmmsString, buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, uint32(len(s)+1))
	binary.Write(buffer, binary.BigEndian, []byte(s))
	binary.Write(buffer, binary.BigEndian, byte(0))
}

func serializeList(l XmmsList, buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, l.Restrict)
	binary.Write(buffer, binary.BigEndian, uint32(len(l.Entries)))
	if l.Restrict != TypeNone {
		// TODO: serialize restricted types
	} else {
		for _, entry := range l.Entries {
			SerializeXmmsValue(entry, buffer)
		}
	}
}

func serializeDict(dict XmmsDict, buffer *bytes.Buffer) {
	binary.Write(buffer, binary.BigEndian, uint32(len(dict)))
	for k, v := range dict {
		serializeString(XmmsString(k), buffer)
		SerializeXmmsValue(v, buffer)
	}
}

func SerializeXmmsValue(value XmmsValue, buffer *bytes.Buffer) {
	switch value.(type) {
	case XmmsInt:
		binary.Write(buffer, binary.BigEndian, TypeInt64)
		serializeInt(value.(XmmsInt), buffer)
	case XmmsString:
		binary.Write(buffer, binary.BigEndian, TypeString)
		serializeString(value.(XmmsString), buffer)
	case XmmsDict:
		binary.Write(buffer, binary.BigEndian, TypeDict)
		serializeDict(value.(XmmsDict), buffer)
	case XmmsList:
		binary.Write(buffer, binary.BigEndian, TypeList)
		serializeList(value.(XmmsList), buffer)
	}
}
