package xmmsclient

import (
	"testing"
)

func TestPropDictToDictGood(t *testing.T) {
	propdict := XmmsDict{
		"artist": XmmsDict{
			"plugin/flac": XmmsString("a"),
			"server":      XmmsString("b"),
		},
	}

	dict, err := PropDictToDictDefault(propdict)
	if err != nil {
		t.Fatal("Error:", err)
	}

	if dict["artist"] != XmmsString("b") {
		t.Fatal("Wrong match:", dict["artist"])
	}
}

func TestPropDictToDictWrongTypes(t *testing.T) {
	_, err := PropDictToDictDefault(XmmsDict{"foo": XmmsString("bar")})
	if err == nil {
		t.Fatal("Error:", err)
	}
}
