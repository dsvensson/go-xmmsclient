package xmmsclient

import (
	"bytes"
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/assert"
)

func roundtrip(value XmmsValue, test func(XmmsValue) bool) bool {
	var buffer bytes.Buffer
	if err := serializeXmmsValue(&buffer, value); err == nil {
		res, err := deserializeXmmsValue(&buffer)
		return err == nil && test(res)
	}
	return false
}

func TestXmmsInt64(t *testing.T) {
	props := gopter.NewProperties(gopter.DefaultTestParameters())
	props.Property("Roundtrip XmmsInt64", prop.ForAll(
		func(v int64) bool {
			return roundtrip(XmmsInt(v), func(res XmmsValue) bool {
				return XmmsInt(v) == res
			})
		},
		gen.Int64(),
	))
	props.TestingRun(t)
}

func TestXmmsString(t *testing.T) {
	props := gopter.NewProperties(gopter.DefaultTestParameters())
	props.Property("Roundtrip XmmsString", prop.ForAll(
		func(v string) bool {
			return roundtrip(XmmsString(v), func(res XmmsValue) bool {
				return XmmsString(v) == res
			})
		},
		gen.AnyString(),
	))
	props.TestingRun(t)
}

func TestXmmsFloat(t *testing.T) {
	props := gopter.NewProperties(gopter.DefaultTestParameters())
	props.Property("Roundtrip XmmsFloat", prop.ForAll(
		func(v float64) bool {
			return roundtrip(XmmsFloat(v), func(res XmmsValue) bool {
				// TODO: Is this rounding error correct?
				return assert.Equal(t, float32(v), float32(res.(XmmsFloat)))
			})
		},
		gen.Float64(),
	))
	props.TestingRun(t)
}

func TestXmmsList(t *testing.T) {
	props := gopter.NewProperties(gopter.DefaultTestParameters())
	props.Property("Roundtrip XmmsList", prop.ForAll(
		func(v []string) bool {
			var buffer bytes.Buffer
			if err := serializeXmmsValue(&buffer, XmmsStrings(v)); err == nil {
				res, err := tryDeserializeStringList(&buffer)
				return err == nil && assert.Equal(t, v, res)
			}
			return false
		},
		gen.SliceOf(gen.AnyString()),
	))
	props.TestingRun(t)
}
