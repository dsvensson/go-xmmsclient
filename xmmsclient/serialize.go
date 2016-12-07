package xmmsclient

import (
	"bytes"
	"encoding/binary"
)

type listProducer func(buffer *bytes.Buffer) error

func serializeInt(buffer *bytes.Buffer, i XmmsInt) error {
	return binary.Write(buffer, binary.BigEndian, i)
}

func serializeString(buffer *bytes.Buffer, s []byte) error {
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

func serializeList(buffer *bytes.Buffer, list XmmsList) error {
	return serializeAnyList(buffer, len(list), TypeNone,
		func(buffer *bytes.Buffer) error {
			for _, entry := range list {
				err := serializeXmmsValue(buffer, entry)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func serializeDict(buffer *bytes.Buffer, dict XmmsDict) error {
	err := binary.Write(buffer, binary.BigEndian, uint32(len(dict)))
	if err != nil {
		return err
	}

	for k, v := range dict {
		err = serializeString(buffer, []byte(k))
		if err != nil {
			return err
		}

		err = serializeXmmsValue(buffer, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func serializeColl(buffer *bytes.Buffer, coll XmmsColl) error {
	err := binary.Write(buffer, binary.BigEndian, coll.Type)
	if err != nil {
		return err
	}

	err = serializeDict(buffer, coll.Attributes)
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
				err := serializeColl(buffer, operand)
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

func serializeXmmsValue(buffer *bytes.Buffer, value XmmsValue) (err error) {
	switch value.(type) {
	case XmmsInt:
		err = binary.Write(buffer, binary.BigEndian, TypeInt64)
		if err == nil {
			return serializeInt(buffer, value.(XmmsInt))
		}
	case XmmsString:
		err = binary.Write(buffer, binary.BigEndian, TypeString)
		if err == nil {
			return serializeString(buffer, []byte(value.(XmmsString)))
		}
	case XmmsError:
		err = binary.Write(buffer, binary.BigEndian, TypeError)
		if err == nil {
			return serializeString(buffer, []byte(value.(XmmsError)))
		}
	case XmmsDict:
		err = binary.Write(buffer, binary.BigEndian, TypeDict)
		if err == nil {
			return serializeDict(buffer, value.(XmmsDict))
		}
	case XmmsList:
		err = binary.Write(buffer, binary.BigEndian, TypeList)
		if err == nil {
			return serializeList(buffer, value.(XmmsList))
		}
	case XmmsColl:
		err = binary.Write(buffer, binary.BigEndian, TypeColl)
		if err == nil {
			return serializeColl(buffer, value.(XmmsColl))
		}
	}
	return nil
}
