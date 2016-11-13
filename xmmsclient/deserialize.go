package xmmsclient

import (
	"bytes"
	"encoding/binary"
)

func deserializeInt(buffer *bytes.Buffer) (value XmmsInt, err error) {
	err = binary.Read(buffer, binary.BigEndian, &value)
	return
}

func deserializeRawString(buffer *bytes.Buffer) (value string, err error) {
	var length uint32
	err = binary.Read(buffer, binary.BigEndian, &length)
	if err != nil {
		return
	}
	var data = make([]byte, length)
	err = binary.Read(buffer, binary.BigEndian, &data)
	if err != nil {
		return
	}
	value = string(data[:length-1])

	return
}

func deserializeString(buffer *bytes.Buffer) (value XmmsString, err error) {
	data, err := deserializeRawString(buffer)
	if err == nil {
		value = XmmsString(data)
	}
	return
}

func deserializeError(buffer *bytes.Buffer) (value XmmsError, err error) {
	data, err := deserializeRawString(buffer)
	if err == nil {
		value = XmmsError(data)
	}
	return
}

func deserializeList(buffer *bytes.Buffer) (value XmmsList, err error) {
	var restrict uint32
	err = binary.Read(buffer, binary.BigEndian, &restrict) // TODO: respect restrict
	if err != nil {
		return
	}
	var length uint32
	err = binary.Read(buffer, binary.BigEndian, &length)
	if err != nil {
		return
	}
	var list = NewXmmsList()
	if restrict != TypeNone {
		for i := uint32(0); i < length; i++ {
			var entry XmmsValue
			entry, err = DeserializeXmmsValueType(restrict, buffer)
			list.Entries = append(list.Entries, entry)
		}
	} else {
		for i := uint32(0); i < length; i++ {
			var entry XmmsValue
			entry, err = DeserializeXmmsValue(buffer)
			list.Entries = append(list.Entries, entry)
		}
	}
	value = list
	return
}

func deserializeDict(buffer *bytes.Buffer) (value XmmsDict, err error) {
	var length uint32
	err = binary.Read(buffer, binary.BigEndian, &length)
	if err != nil {
		return
	}
	var dict = XmmsDict{}
	for i := uint32(0); i < length; i++ {
		var entry XmmsValue
		var key string

		key, err = deserializeRawString(buffer)
		if err != nil {
			return
		}

		entry, err = DeserializeXmmsValue(buffer)
		if err != nil {
			return
		}

		dict[string(key)] = entry
	}
	value = dict
	return
}

func DeserializeXmmsValueType(valueType uint32, buffer *bytes.Buffer) (result XmmsValue, err error) {
	switch valueType {
	case TypeInt64:
		result, err = deserializeInt(buffer)
	case TypeError:
		result, err = deserializeError(buffer)
	case TypeString:
		result, err = deserializeString(buffer)
	case TypeList:
		result, err = deserializeList(buffer)
	case TypeDict:
		result, err = deserializeDict(buffer)
	}
	return
}

func DeserializeXmmsValue(buffer *bytes.Buffer) (XmmsValue, error) {
	var valueType uint32
	binary.Read(buffer, binary.BigEndian, &valueType)
	return DeserializeXmmsValueType(valueType, buffer)
}
