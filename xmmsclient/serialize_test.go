package xmmsclient

import (
	"bytes"
	"reflect"
	"testing"
)

func TestSerializeInt(t *testing.T) {
	var expected = []byte{
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a,
	}
	var buffer bytes.Buffer

	var err = SerializeXmmsValue(XmmsInt(42), &buffer)
	if err != nil {
		t.Error(err)
	}

	if !reflect.DeepEqual(expected, buffer.Bytes()) {
		t.Errorf("\n\twant %+v\n\thave %+v", expected, buffer.Bytes())
	}
}
