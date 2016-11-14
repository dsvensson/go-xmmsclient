package xmmsclient

type XmmsValue interface {
	isXmmsValue()
}
type XmmsInt int64
type XmmsFloat float64
type XmmsError string
type XmmsString string
type XmmsDict map[string]XmmsValue
type XmmsList []XmmsValue
type XmmsRestrictedList struct {
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

func (l XmmsRestrictedList) isXmmsValue() {}

func (l XmmsColl) isXmmsValue() {}
