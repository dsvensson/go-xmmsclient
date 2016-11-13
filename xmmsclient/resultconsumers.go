// auto-generated
package xmmsclient

type genericConsumerType struct {
	value XmmsValue
	err   error
}

type genericConsumer struct {
	result chan genericConsumerType
}

func newGenericConsumer() genericConsumer {
	return genericConsumer{make(chan genericConsumerType)}
}

func (r *genericConsumer) post(value XmmsValue, err error) {
	if err != nil {
		r.result <- genericConsumerType{value, err}
	} else {
		r.result <- genericConsumerType{value, err}
	}
}

type intConsumerType struct {
	value XmmsInt
	err   error
}

type intConsumer struct {
	result chan intConsumerType
}

func newIntConsumer() intConsumer {
	return intConsumer{make(chan intConsumerType)}
}

func (r *intConsumer) post(value XmmsValue, err error) {
	if err != nil {
		r.result <- intConsumerType{0, err}
	} else {
		r.result <- intConsumerType{value.(XmmsInt), err}
	}
}

type stringConsumerType struct {
	value XmmsString
	err   error
}

type stringConsumer struct {
	result chan stringConsumerType
}

func newStringConsumer() stringConsumer {
	return stringConsumer{make(chan stringConsumerType)}
}

func (r *stringConsumer) post(value XmmsValue, err error) {
	if err != nil {
		r.result <- stringConsumerType{"", err}
	} else {
		r.result <- stringConsumerType{value.(XmmsString), err}
	}
}

type listConsumerType struct {
	value XmmsList
	err   error
}

type listConsumer struct {
	result chan listConsumerType
}

func newListConsumer() listConsumer {
	return listConsumer{make(chan listConsumerType)}
}

func (r *listConsumer) post(value XmmsValue, err error) {
	if err != nil {
		r.result <- listConsumerType{XmmsList{}, err}
	} else {
		r.result <- listConsumerType{value.(XmmsList), err}
	}
}

type dictConsumerType struct {
	value XmmsDict
	err   error
}

type dictConsumer struct {
	result chan dictConsumerType
}

func newDictConsumer() dictConsumer {
	return dictConsumer{make(chan dictConsumerType)}
}

func (r *dictConsumer) post(value XmmsValue, err error) {
	if err != nil {
		r.result <- dictConsumerType{XmmsDict{}, err}
	} else {
		r.result <- dictConsumerType{value.(XmmsDict), err}
	}
}

type collConsumerType struct {
	value XmmsColl
	err   error
}

type collConsumer struct {
	result chan collConsumerType
}

func newCollConsumer() collConsumer {
	return collConsumer{make(chan collConsumerType)}
}

func (r *collConsumer) post(value XmmsValue, err error) {
	if err != nil {
		r.result <- collConsumerType{XmmsColl{}, err}
	} else {
		r.result <- collConsumerType{value.(XmmsColl), err}
	}
}
