package xmmsclient

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPropDictToDictGood(t *testing.T) {
	propdict := XmmsDict{
		"artist": XmmsDict{
			"plugin/flac": XmmsString("a"),
			"server":      XmmsString("b"),
		},
	}

	dict, err := PropDictToDictDefault(propdict)
	require.NoError(t, err)
	require.Equal(t, XmmsString("b"), dict["artist"])
}

func TestPropDictToDictWrongTypes(t *testing.T) {
	value, err := PropDictToDictDefault(XmmsDict{"foo": XmmsString("bar")})
	require.Nil(t, value)
	require.Error(t, err)
}
