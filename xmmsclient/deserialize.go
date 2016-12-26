package xmmsclient

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
)

type listConsumer func(value XmmsValue)

func deserializeInt(r io.Reader) (value XmmsInt, err error) {
	err = binary.Read(r, binary.BigEndian, &value)
	return
}

func deserializeFloat(r io.Reader) (value XmmsFloat, err error) {
	var mantissaInt int32
	var exponent int32
	var mantissa float64

	err = binary.Read(r, binary.BigEndian, &mantissaInt)
	if err != nil {
		return
	}

	err = binary.Read(r, binary.BigEndian, &exponent)
	if err != nil {
		return
	}

	if mantissaInt > 0 {
		mantissa = float64(mantissaInt) / float64(math.MaxInt32)
	} else {
		mantissa = float64(mantissaInt) / math.Abs(math.MinInt32)
	}

	value = XmmsFloat(math.Ldexp(mantissa, int(exponent)))

	return
}

func deserializeRawString(r io.Reader) (value string, err error) {
	var length uint32
	err = binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return
	}
	data := make([]byte, length)
	err = binary.Read(r, binary.BigEndian, &data)
	if err != nil {
		return
	}
	value = string(data[:length-1])

	return
}

func deserializeString(r io.Reader) (value XmmsString, err error) {
	data, err := deserializeRawString(r)
	if err == nil {
		value = XmmsString(data)
	}
	return
}

func deserializeError(r io.Reader) (value XmmsError, err error) {
	data, err := deserializeRawString(r)
	if err == nil {
		value = XmmsError(data)
	}
	return
}

func deserializeAnyList(r io.Reader, consumer listConsumer) error {
	var restrict uint32
	err := binary.Read(r, binary.BigEndian, &restrict) // TODO: respect restrict
	if err != nil {
		return err
	}
	var length uint32
	err = binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return err
	}
	if restrict != TypeNone {
		for i := uint32(0); i < length; i++ {
			entry, err := deserializeXmmsValueOfType(restrict, r)
			if err != nil {
				return err
			}
			consumer(entry)
		}
	} else {
		for i := uint32(0); i < length; i++ {
			entry, err := deserializeXmmsValue(r)
			if err != nil {
				return err
			}
			consumer(entry)
		}
	}
	return nil
}

func deserializeList(r io.Reader) (value XmmsList, err error) {
	list := XmmsList{}
	err = deserializeAnyList(r, func(value XmmsValue) {
		list = append(list, value)
	})
	value = list
	return
}

func deserializeDict(r io.Reader) (value XmmsDict, err error) {
	var length uint32
	err = binary.Read(r, binary.BigEndian, &length)
	if err != nil {
		return
	}
	dict := XmmsDict{}
	for i := uint32(0); i < length; i++ {
		var entry XmmsValue
		var key string

		key, err = deserializeRawString(r)
		if err != nil {
			return
		}

		entry, err = deserializeXmmsValue(r)
		if err != nil {
			return
		}

		dict[key] = entry
	}
	value = dict
	return
}

func deserializeColl(r io.Reader) (result XmmsColl, err error) {
	err = binary.Read(r, binary.BigEndian, &result.Type)
	if err != nil {
		return
	}

	result.Attributes, err = deserializeDict(r)
	if err != nil {
		return
	}

	err = deserializeAnyList(r, func(raw XmmsValue) {
		if value, ok := raw.(XmmsInt); ok {
			result.IDList = append(result.IDList, int(value))
		}
	})
	if err != nil {
		return
	}

	err = deserializeAnyList(r, func(raw XmmsValue) {
		if value, ok := raw.(XmmsColl); ok {
			result.Operands = append(result.Operands, value)
		}
	})

	return
}

func deserializeXmmsValueOfType(valueType uint32, r io.Reader) (result XmmsValue, err error) {
	switch valueType {
	case TypeInt64:
		result, err = deserializeInt(r)
	case TypeFloat:
		result, err = deserializeFloat(r)
	case TypeError:
		result, err = deserializeError(r)
	case TypeString:
		result, err = deserializeString(r)
	case TypeList:
		result, err = deserializeList(r)
	case TypeDict:
		result, err = deserializeDict(r)
	case TypeColl:
		result, err = deserializeColl(r)
	}
	return
}

func deserializeXmmsValue(r io.Reader) (XmmsValue, error) {
	var valueType uint32
	err := binary.Read(r, binary.BigEndian, &valueType)
	if err != nil {
		return nil, err
	}
	return deserializeXmmsValueOfType(valueType, r)
}

func tryDeserialize(r io.Reader) (XmmsValue, error) {
	value, err := deserializeXmmsValue(r)
	if err != nil {
		return nil, err
	}

	errorMessage, ok := value.(XmmsError)
	if ok {
		return nil, fmt.Errorf("%s", errorMessage)
	}

	return value, nil
}

func tryDeserializeList(r io.Reader, consumer listConsumer) error {
	var valueType uint32

	err := binary.Read(r, binary.BigEndian, &valueType)
	if err != nil {
		return err
	}

	switch valueType {
	case TypeError:
		value, err := deserializeError(r)
		if err != nil {
			return err
		}
		return fmt.Errorf("%s", value)
	case TypeList:
		return deserializeAnyList(r, consumer)
	default:
		return fmt.Errorf("Trying to parse non-list as list (%v)", valueType)
	}
}

func tryDeserializeIntList(r io.Reader) ([]int, error) {
	var list []int
	err := tryDeserializeList(r, func(raw XmmsValue) {
		if value, ok := raw.(XmmsInt); ok {
			list = append(list, int(value))
		}
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func tryDeserializeStringList(r io.Reader) ([]string, error) {
	var list []string
	err := tryDeserializeList(r, func(raw XmmsValue) {
		if value, ok := raw.(XmmsString); ok {
			list = append(list, string(value))
		}
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}

func tryDeserializeDictList(r io.Reader) ([]XmmsDict, error) {
	var list []XmmsDict
	err := tryDeserializeList(r, func(raw XmmsValue) {
		if value, ok := raw.(XmmsDict); ok {
			list = append(list, value)
		}
	})
	if err != nil {
		return nil, err
	}
	return list, nil
}
