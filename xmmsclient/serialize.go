package xmmsclient

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
)

type ioWriter func(w io.Writer) error

func write(data interface{}) ioWriter {
	return func(w io.Writer) error {
		return binary.Write(w, binary.BigEndian, data)
	}
}

func compose(funs ...ioWriter) ioWriter {
	return func(w io.Writer) error {
		for _, fun := range funs {
			err := fun(w)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func serializeInt(i XmmsInt) ioWriter {
	return write(i)
}

func splitFloat(v float64) (int32, int32) {
	mant, exp := math.Frexp(v)
	if v > 0 {
		return int32(mant * math.MaxInt32), int32(exp)
	}
	return int32(mant * math.Abs(math.MinInt32)), int32(exp)
}

func serializeFloat(i XmmsFloat) ioWriter {
	mant, exp := splitFloat(float64(i))
	return compose(
		write(mant),
		write(exp),
	)
}

func serializeString(s []byte) ioWriter {
	return compose(
		write(uint32(len(s)+1)),
		write(s),
		write(byte(0)),
	)
}

func serializeAnyList(length int, restrict uint32, producer ioWriter) ioWriter {
	return compose(
		write(restrict),
		write(uint32(length)),
		producer,
	)
}

func serializeList(list XmmsList) ioWriter {
	return serializeAnyList(len(list), TypeNone,
		func(w io.Writer) error {
			for _, entry := range list {
				err := serializeXmmsValue2(entry)(w)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func serializeStringList(list XmmsStrings) ioWriter {
	return serializeAnyList(len(list), TypeNone,
		func(w io.Writer) error {
			for _, entry := range list {
				err := serializeXmmsValue2(XmmsString(entry))(w)
				if err != nil {
					return err
				}
			}
			return nil
		},
	)
}

func serializeDict(dict XmmsDict) ioWriter {
	return func(w io.Writer) error {
		err := binary.Write(w, binary.BigEndian, uint32(len(dict)))
		if err != nil {
			return err
		}

		for k, v := range dict {
			err := compose(
				serializeString([]byte(k)),
				serializeXmmsValue2(v),
			)(w)
			if err != nil {
				return err
			}
		}

		return nil
	}
}

func serializeIdList(ids []int) ioWriter {
	return func(w io.Writer) error {
		for _, id := range ids {
			err := serializeInt(XmmsInt(id))(w)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func serializeOperands(operands []XmmsColl) ioWriter {
	return func(w io.Writer) error {
		for _, operand := range operands {
			err := serializeColl(operand)(w)
			if err != nil {
				return err
			}
		}
		return nil
	}
}

func serializeColl(coll XmmsColl) ioWriter {
	return compose(
		write(coll.Type),
		serializeDict(coll.Attributes),
		serializeAnyList(len(coll.IDList), TypeInt64, serializeIdList(coll.IDList)),
		serializeAnyList(len(coll.Operands), TypeColl, serializeOperands(coll.Operands)),
	)
}

func serializeXmmsValue2(value XmmsValue) ioWriter {
	var typ uint32
	var serializer ioWriter

	switch value.(type) {
	case XmmsInt:
		typ, serializer = TypeInt64, serializeInt(value.(XmmsInt))
	case XmmsFloat:
		typ, serializer = TypeFloat, serializeFloat(value.(XmmsFloat))
	case XmmsString:
		typ, serializer = TypeString, serializeString([]byte(value.(XmmsString)))
	case XmmsError:
		typ, serializer = TypeError, serializeString([]byte(value.(XmmsError)))
	case XmmsDict:
		typ, serializer = TypeDict, serializeDict(value.(XmmsDict))
	case XmmsList:
		typ, serializer = TypeList, serializeList(value.(XmmsList))
	case XmmsStrings:
		typ, serializer = TypeList, serializeStringList(value.(XmmsStrings))
	case XmmsColl:
		typ, serializer = TypeColl, serializeColl(value.(XmmsColl))
	default:
		return func(w io.Writer) error {
			return errors.New("Unknown type")
		}
	}
	return compose(
		write(typ),
		serializer,
	)
}

func serializeXmmsValue(w io.Writer, value XmmsValue) error {
	return serializeXmmsValue2(value)(w)
}
