package xmmsclient

import (
	"encoding/binary"
	"io"
)

type listProducer func(w io.Writer) error

func serializeInt(w io.Writer, i XmmsInt) error {
	return binary.Write(w, binary.BigEndian, i)
}

func serializeString(w io.Writer, s []byte) error {
	err := binary.Write(w, binary.BigEndian, uint32(len(s)+1))
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, s)
	if err != nil {
		return err
	}

	return binary.Write(w, binary.BigEndian, byte(0))
}

func serializeAnyList(w io.Writer, length int, restrict uint32, producer listProducer) error {
	err := binary.Write(w, binary.BigEndian, restrict)
	if err != nil {
		return err
	}

	err = binary.Write(w, binary.BigEndian, uint32(length))
	if err != nil {
		return err
	}

	return producer(w)
}

func serializeList(w io.Writer, list XmmsList) error {
	return serializeAnyList(w, len(list), TypeNone,
		func(w io.Writer) error {
			for _, entry := range list {
				err := serializeXmmsValue(w, entry)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func serializeStringList(w io.Writer, list XmmsStrings) error {
	return serializeAnyList(w, len(list), TypeNone,
		func(w io.Writer) error {
			for _, entry := range list {
				err := serializeXmmsValue(w, XmmsString(entry))
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func serializeDict(w io.Writer, dict XmmsDict) error {
	err := binary.Write(w, binary.BigEndian, uint32(len(dict)))
	if err != nil {
		return err
	}

	for k, v := range dict {
		err = serializeString(w, []byte(k))
		if err != nil {
			return err
		}

		err = serializeXmmsValue(w, v)
		if err != nil {
			return err
		}
	}

	return nil
}

func serializeColl(w io.Writer, coll XmmsColl) error {
	err := binary.Write(w, binary.BigEndian, coll.Type)
	if err != nil {
		return err
	}

	err = serializeDict(w, coll.Attributes)
	if err != nil {
		return err
	}

	err = serializeAnyList(w, len(coll.IDList), TypeInt64,
		func(w io.Writer) error {
			for _, id := range coll.IDList {
				innerErr := binary.Write(w, binary.BigEndian, int64(id))
				if innerErr != nil {
					return innerErr
				}
			}
			return nil
		},
	)
	if err != nil {
		return err
	}

	return serializeAnyList(w, len(coll.Operands), TypeColl,
		func(w io.Writer) error {
			for _, operand := range coll.Operands {
				innerErr := serializeColl(w, operand)
				if innerErr != nil {
					return innerErr
				}
			}
			return nil
		},
	)
}

func serializeXmmsValue(w io.Writer, value XmmsValue) (err error) {
	switch value.(type) {
	case XmmsInt:
		err = binary.Write(w, binary.BigEndian, TypeInt64)
		if err == nil {
			return serializeInt(w, value.(XmmsInt))
		}
	case XmmsString:
		err = binary.Write(w, binary.BigEndian, TypeString)
		if err == nil {
			return serializeString(w, []byte(value.(XmmsString)))
		}
	case XmmsError:
		err = binary.Write(w, binary.BigEndian, TypeError)
		if err == nil {
			return serializeString(w, []byte(value.(XmmsError)))
		}
	case XmmsDict:
		err = binary.Write(w, binary.BigEndian, TypeDict)
		if err == nil {
			return serializeDict(w, value.(XmmsDict))
		}
	case XmmsList:
		err = binary.Write(w, binary.BigEndian, TypeList)
		if err == nil {
			return serializeList(w, value.(XmmsList))
		}
	case XmmsStrings:
		err = binary.Write(w, binary.BigEndian, TypeList)
		if err == nil {
			return serializeStringList(w, value.(XmmsStrings))
		}
	case XmmsColl:
		err = binary.Write(w, binary.BigEndian, TypeColl)
		if err == nil {
			return serializeColl(w, value.(XmmsColl))
		}
	}
	return nil
}
