package xmmsclient

type XmmsValue interface {
	isXmmsValue()
}
type XmmsInt int64
type XmmsString string
type XmmsDict map[string]XmmsValue
type XmmsList struct {
	Entries  []XmmsValue
	Restrict uint32
}

func (i XmmsInt) isXmmsValue() {}

func (s XmmsString) isXmmsValue() {}

func (d XmmsDict) isXmmsValue() {}

func (l XmmsList) isXmmsValue() {}

func NewXmmsList(entries ...XmmsValue) XmmsList {
	return XmmsList{Entries: entries, Restrict: TypeNone}
}
