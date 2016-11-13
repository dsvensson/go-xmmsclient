package xmmsclient

import (
	"errors"
)

type XmmsValue interface {
	isXmmsValue()
}
type XmmsInt int64
type XmmsFloat float64
type XmmsError string
type XmmsString string
type XmmsDict map[string]XmmsValue
type XmmsList struct {
	Entries  []XmmsValue
	Restrict uint32
}
type XmmsColl struct {
	Type       uint32
	Operands   XmmsList
	Attributes XmmsDict
	IdList     XmmsList
}

func (i XmmsInt) isXmmsValue() {}

func (i XmmsFloat) isXmmsValue() {}

func (s XmmsError) isXmmsValue() {}

func (s XmmsString) isXmmsValue() {}

func (d XmmsDict) isXmmsValue() {}

func (l XmmsList) isXmmsValue() {}

func (l XmmsColl) isXmmsValue() {}

func NewXmmsList(entries ...XmmsValue) XmmsList {
	return XmmsList{Entries: entries, Restrict: TypeNone}
}

func valueAsInt(raw XmmsValue) (XmmsInt, error) {
	if value, ok := raw.(XmmsInt); ok {
		return value, nil
	}
	return 0, errors.New("Bad type")
}

func valueAsString(raw XmmsValue) (XmmsString, error) {
	if value, ok := raw.(XmmsString); ok {
		return value, nil
	}
	return "", errors.New("Bad type")
}

func valueAsList(raw XmmsValue) (XmmsList, error) {
	if value, ok := raw.(XmmsList); ok {
		return value, nil
	}
	return XmmsList{}, errors.New("Bad type")
}

func valueAsDict(raw XmmsValue) (XmmsDict, error) {
	if value, ok := raw.(XmmsDict); ok {
		return value, nil
	}
	return XmmsDict{}, errors.New("Bad type")
}
