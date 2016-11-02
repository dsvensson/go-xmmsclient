package xmmsclient

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"testing"
)

func checkBuffer(t *testing.T, expected []byte, actual []byte) {
	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("\nwant:\n%s\nhave:\n%s", hex.Dump(expected), hex.Dump(actual))
	}
}

func TestSerializeInt(t *testing.T) {
	var expected = []byte{
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a,
	}
	var buffer bytes.Buffer

	var err = SerializeXmmsValue(XmmsInt(42), &buffer)
	if err != nil {
		t.Fatal(err)
	}

	checkBuffer(t, expected, buffer.Bytes())
}
