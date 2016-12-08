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
	expected := []byte{
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a,
	}
	var buffer bytes.Buffer

	err := serializeXmmsValue(&buffer, XmmsInt(42))
	if err != nil {
		t.Fatal(err)
	}

	checkBuffer(t, expected, buffer.Bytes())
}

func TestSerializeString(t *testing.T) {
	expected := []byte{
		0x00, 0x00, 0x00, 0x03,
		0x00, 0x00, 0x00, 0x05,
		0x74, 0x65, 0x73, 0x74, 0x00,
	}
	var buffer bytes.Buffer

	err := serializeXmmsValue(&buffer, XmmsString("test"))
	if err != nil {
		t.Fatal(err)
	}

	checkBuffer(t, expected, buffer.Bytes())
}

func TestSerializeError(t *testing.T) {
	expected := []byte{
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x05,
		0x74, 0x65, 0x73, 0x74, 0x00,
	}
	var buffer bytes.Buffer

	err := serializeXmmsValue(&buffer, XmmsError("test"))
	if err != nil {
		t.Fatal(err)
	}

	checkBuffer(t, expected, buffer.Bytes())
}

func TestSerializeList(t *testing.T) {
	expected := []byte{
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

	err := serializeXmmsValue(&buffer, XmmsList{XmmsInt(42), XmmsString("test")})
	if err != nil {
		t.Fatal(err)
	}

	checkBuffer(t, expected, buffer.Bytes())
}

func TestSerializeDict(t *testing.T) {
	expected := []byte{
		0x00, 0x00, 0x00, 0x07,
		0x00, 0x00, 0x00, 0x01,
		0x00, 0x00, 0x00, 0x04,
		0x69, 0x6e, 0x74, 0x00,
		0x00, 0x00, 0x00, 0x02,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x2a,
	}
	var buffer bytes.Buffer

	err := serializeXmmsValue(&buffer, XmmsDict{"int": XmmsInt(42)})
	if err != nil {
		t.Fatal(err)
	}

	checkBuffer(t, expected, buffer.Bytes())
}

func TestSerializeColl(t *testing.T) {
	expected := []byte{
		0x00, 0x00, 0x00, 0x04, /* XMMSV_TYPE_COLL */
		0x00, 0x00, 0x00, 0x06, /* XMMS_COLLECTION_TYPE_MATCH */
		0x00, 0x00, 0x00, 0x01, /* number of attributes*/

		0x00, 0x00, 0x00, 0x05, /* attr[0] key length */
		0x73, 0x65, 0x65, 0x64, /* attr[0] key "seed" */
		0x00, /*                               "\0"   */

		0x00, 0x00, 0x00, 0x02, /* attr[0] value type  */
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x7a, 0x69, /* attr[0] value 31337 */

		0x00, 0x00, 0x00, 0x02, /* idlist: restrict type XMMSV_TYPE_INT64 */
		0x00, 0x00, 0x00, 0x00, /* idlist: count */
		0x00, 0x00, 0x00, 0x04, /* operands: restrict type XMMSV_TYPE_COLL */
		0x00, 0x00, 0x00, 0x01, /* operands: number of operands */

		0x00, 0x00, 0x00, 0x01, /* operand[0]: coll type (_UNIVERSE) */

		0x00, 0x00, 0x00, 0x00, /* number of attributes*/
		0x00, 0x00, 0x00, 0x02, /* idlist: restrict type XMMSV_TYPE_INT64 */
		0x00, 0x00, 0x00, 0x00, /* idlist: count */
		0x00, 0x00, 0x00, 0x04, /* operands: restrict type XMMSV_TYPE_COLL */
		0x00, 0x00, 0x00, 0x00, /* operands: count */
	}
	var buffer bytes.Buffer

	match := XmmsColl{
		Type:       CollectionTypeMatch,
		Attributes: XmmsDict{"seed": XmmsInt(31337)},
		Operands: []XmmsColl{
			XmmsColl{Type: CollectionTypeUniverse},
		},
	}

	err := serializeXmmsValue(&buffer, match)
	if err != nil {
		t.Fatal(err)
	}

	checkBuffer(t, expected, buffer.Bytes())
}
