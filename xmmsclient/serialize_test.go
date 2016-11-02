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

func TestSerializeString(t *testing.T) {
	var expected = []byte{
		0x00, 0x00, 0x00, 0x03,
		0x00, 0x00, 0x00, 0x05,
		0x74, 0x65, 0x73, 0x74, 0x00,
	}
	var buffer bytes.Buffer

	var err = SerializeXmmsValue(XmmsString("test"), &buffer)
	if err != nil {
		t.Fatal(err)
	}

	checkBuffer(t, expected, buffer.Bytes())
}

func TestSerializeList(t *testing.T) {
	var expected = []byte{
		0x00, 0x00, 0x00, 0x06,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a,
		0x00, 0x00, 0x00, 0x03,
		0x00, 0x00, 0x00, 0x05,
		0x74, 0x65, 0x73, 0x74, 0x00,
	}
	var buffer bytes.Buffer

	var err = SerializeXmmsValue(NewXmmsList(XmmsInt(42), XmmsString("test")), &buffer)
	if err != nil {
		t.Fatal(err)
	}

	checkBuffer(t, expected, buffer.Bytes())
}

func TestSerializeDict(t *testing.T) {
	var expected = []byte{
		0x00, 0x00, 0x00, 0x07,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x04,
		0x69, 0x6e, 0x74, 0x00,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a,
	}
	var buffer bytes.Buffer

	var err = SerializeXmmsValue(XmmsDict{"int": XmmsInt(42)}, &buffer)
	if err != nil {
		t.Fatal(err)
	}

	checkBuffer(t, expected, buffer.Bytes())
}
