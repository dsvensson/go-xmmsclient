package xmmsclient

import (
	"bytes"
	"encoding/binary"
)

func serializeInt(i XmmsInt, buffer *bytes.Buffer) error {
	return binary.Write(buffer, binary.BigEndian, i)
}

func serializeString(s []byte, buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.BigEndian, uint32(len(s)+1))
	if err != nil {
		return err
	}

	err = binary.Write(buffer, binary.BigEndian, s)
	if err != nil {
		return err
	}

	err = binary.Write(buffer, binary.BigEndian, byte(0))
	if err != nil {
		return err
	}

	return nil
}

func serializeList(l XmmsList, buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.BigEndian, TypeNone)
	if err != nil {
		return err
	}

	err = binary.Write(buffer, binary.BigEndian, uint32(len(l)))
	if err != nil {
		return err
	}

	for _, entry := range l {
		err = serializeXmmsValue(entry, buffer)
		if err != nil {
			return err
		}
	}

	return nil
}

func serializeDict(dict XmmsDict, buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.BigEndian, uint32(len(dict)))
	if err != nil {
		return err
	}

	for k, v := range dict {
		err = serializeString([]byte(k), buffer)
		if err != nil {
			return err
		}

		err = serializeXmmsValue(v, buffer)
		if err != nil {
			return err
		}
	}

	return nil
}

func serializeXmmsValue(value XmmsValue, buffer *bytes.Buffer) (err error) {
	switch value.(type) {
	case XmmsInt:
		err = binary.Write(buffer, binary.BigEndian, TypeInt64)
		if err == nil {
			serializeInt(value.(XmmsInt), buffer)
		}
	case XmmsString:
		err = binary.Write(buffer, binary.BigEndian, TypeString)
		if err == nil {
			serializeString([]byte(value.(XmmsString)), buffer)
		}
	case XmmsError:
		err = binary.Write(buffer, binary.BigEndian, TypeError)
		if err == nil {
			serializeString([]byte(value.(XmmsError)), buffer)
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
	return nil
}
