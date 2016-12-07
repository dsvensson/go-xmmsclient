package xmmsclient

import (
	"bytes"
	"encoding/binary"
)

type listProducer func(buffer *bytes.Buffer) error

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

func serializeAnyList(buffer *bytes.Buffer, length int, restrict uint32, producer listProducer) error {
	err := binary.Write(buffer, binary.BigEndian, restrict)
	if err != nil {
		return err
	}

	err = binary.Write(buffer, binary.BigEndian, uint32(length))
	if err != nil {
		return err
	}

	return producer(buffer)
}

func serializeList(list XmmsList, buffer *bytes.Buffer) error {
	return serializeAnyList(buffer, len(list), TypeNone,
		func(buffer *bytes.Buffer) error {
			for _, entry := range list {
				err := serializeXmmsValue(entry, buffer)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
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

func serializeColl(coll XmmsColl, buffer *bytes.Buffer) error {
	err := binary.Write(buffer, binary.BigEndian, coll.Type)
	if err != nil {
		return err
	}

	err = serializeDict(coll.Attributes, buffer)
	if err != nil {
		return err
	}

	err = serializeAnyList(buffer, len(coll.IdList), TypeInt64,
		func(buffer *bytes.Buffer) error {
			for _, id := range coll.IdList {
				err := binary.Write(buffer, binary.BigEndian, id)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
	if err != nil {
		return err
	}

	err = serializeAnyList(buffer, len(coll.Operands), TypeColl,
		func(buffer *bytes.Buffer) error {
			for _, operand := range coll.Operands {
				err := serializeColl(operand, buffer)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
	if err != nil {
		return err
	}

	return nil
}

func serializeXmmsValue(value XmmsValue, buffer *bytes.Buffer) (err error) {
	switch value.(type) {
	case XmmsInt:
		err = binary.Write(buffer, binary.BigEndian, TypeInt64)
		if err == nil {
			return serializeInt(value.(XmmsInt), buffer)
		}
	case XmmsString:
		err = binary.Write(buffer, binary.BigEndian, TypeString)
		if err == nil {
			return serializeString([]byte(value.(XmmsString)), buffer)
		}
	case XmmsError:
		err = binary.Write(buffer, binary.BigEndian, TypeError)
		if err == nil {
			return serializeString([]byte(value.(XmmsError)), buffer)
		}
	case XmmsDict:
		err = binary.Write(buffer, binary.BigEndian, TypeDict)
		if err == nil {
			return serializeDict(value.(XmmsDict), buffer)
		}
	case XmmsList:
		err = binary.Write(buffer, binary.BigEndian, TypeList)
		if err == nil {
			return serializeList(value.(XmmsList), buffer)
		}
	case XmmsColl:
		err = binary.Write(buffer, binary.BigEndian, TypeColl)
		if err == nil {
			return serializeColl(value.(XmmsColl), buffer)
		}
	}
	return nil
}
