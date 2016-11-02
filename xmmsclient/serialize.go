package xmmsclient

import (
	"bytes"
	"encoding/binary"
)

func serializeInt(i XmmsInt, buffer *bytes.Buffer) (err error) {
	err = binary.Write(buffer, binary.BigEndian, i)
	return
}

func serializeString(s XmmsString, buffer *bytes.Buffer) (err error) {
	err = binary.Write(buffer, binary.BigEndian, uint32(len(s)+1))
	if err != nil {
		return
	}

	err = binary.Write(buffer, binary.BigEndian, []byte(s))
	if err != nil {
		return
	}

	err = binary.Write(buffer, binary.BigEndian, byte(0))
	if err != nil {
		return
	}

	return
}

func serializeList(l XmmsList, buffer *bytes.Buffer) (err error) {
	err = binary.Write(buffer, binary.BigEndian, l.Restrict)
	if err != nil {
		return
	}

	err = binary.Write(buffer, binary.BigEndian, uint32(len(l.Entries)))
	if err != nil {
		return
	}

	if l.Restrict != TypeNone {
		// TODO: serialize restricted types
	} else {
		for _, entry := range l.Entries {
			err = SerializeXmmsValue(entry, buffer)
			if err != nil {
				return
			}
		}
	}

	return
}

func serializeDict(dict XmmsDict, buffer *bytes.Buffer) (err error) {
	err = binary.Write(buffer, binary.BigEndian, uint32(len(dict)))
	if err != nil {
		return
	}

	for k, v := range dict {
		err = serializeString(XmmsString(k), buffer)
		if err != nil {
			return
		}

		SerializeXmmsValue(v, buffer)
	}

	return
}

func SerializeXmmsValue(value XmmsValue, buffer *bytes.Buffer) (err error) {
	switch value.(type) {
	case XmmsInt:
		err = binary.Write(buffer, binary.BigEndian, TypeInt64)
		if err == nil {
			serializeInt(value.(XmmsInt), buffer)
		}
	case XmmsString:
		err = binary.Write(buffer, binary.BigEndian, TypeString)
		if err == nil {
			serializeString(value.(XmmsString), buffer)
		}
	case XmmsDict:
		err = binary.Write(buffer, binary.BigEndian, TypeDict)
		if err == nil {
			serializeDict(value.(XmmsDict), buffer)
		}
	case XmmsList:
		err = binary.Write(buffer, binary.BigEndian, TypeList)
		if err == nil {
			serializeList(value.(XmmsList), buffer)
		}
	}
	return
}
