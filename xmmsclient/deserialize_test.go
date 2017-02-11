package xmmsclient

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeserializeString(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x00, 0x03, // XMMSV_TYPE_STRING
		0x00, 0x00, 0x00, 0x04, // 4 (length of following bytes)
		0x66, 0x6f, 0x6f, 0x00, // "foo\0"
	})

	value, err := deserializeXmmsValue(buffer)
	require.NoError(t, err)
	require.Equal(t, XmmsString("foo"), value)
}

func TestDeserializeError(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x00, 0x01, // XMMSV_TYPE_ERROR
		0x00, 0x00, 0x00, 0x04, // 4 (length of following bytes)
		0x66, 0x6f, 0x6f, 0x00, // "foo\0"
	})

	value, err := deserializeXmmsValue(buffer)
	require.NoError(t, err)
	require.Equal(t, XmmsError("foo"), value)
}

/*
func TestDeserializeBindata(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x00, 0x05, // XMMSV_TYPE_BIN
		0x00, 0x00, 0x00, 0x08, // 8 (length of following bytes)
		0x01, 0x02, 0x03, 0x04,
		0x00, 0x01, 0x02, 0x03,
	})

	value, err := deserializeXmmsValue(buffer)
	require.NoError(t, err)

	bindata := value.(XmmsBindata)
	require.Equal(t, XmmsInt(42), bindata)
}
*/

func TestDeserializeInt(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x00, 0x02, // XMMSV_TYPE_INT64
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x2a, // 42
	})

	value, err := deserializeXmmsValue(buffer)
	require.NoError(t, err)
	require.Equal(t, XmmsInt(42), value)
}

func TestDeserializeFloat(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x00, 0x09, // XMMSV_TYPE_FLOAT
		0x60, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	})

	value, err := deserializeXmmsValue(buffer)
	require.NoError(t, err)

	float := value.(XmmsFloat)

	rounded := float64(int(float*10000)) / 10000.0
	require.Equal(t, float64(0.75), rounded)
}

func TestDeserializeList(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x00, 0x06, // XMMSV_TYPE_LIST
		0x00, 0x00, 0x00, 0x00, // restrict to XMMSV_TYPE_NONE
		0x00, 0x00, 0x00, 0x03, // 3 (number of list items)

		0x00, 0x00, 0x00, 0x02, // list[0]: XMMSV_TYPE_INT64
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x2a, // list[0]: 42

		0x00, 0x00, 0x00, 0x09, // list[0]: XMMSV_TYPE_FLOAT
		0xc0, 0x00, 0x00, 0x00, // list[0]: -1.0
		0x00, 0x00, 0x00, 0x01,

		0x00, 0x00, 0x00, 0x03, // list[1]: XMMSV_TYPE_STRING
		0x00, 0x00, 0x00, 0x04, // list[1]: 4 (length of following bytes)
		0x66, 0x6f, 0x6f, 0x00, // list[1]: "foo\0"
	})

	value, err := deserializeXmmsValue(buffer)
	require.NoError(t, err)

	list := value.(XmmsList)
	require.Len(t, list, 3)
	require.Equal(t, XmmsInt(42), list[0])
	require.Equal(t, XmmsFloat(-1.), list[1])
	require.Equal(t, XmmsString("foo"), list[2])
}

func TestDeserializeRestrictedList(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x00, 0x06, // XMMSV_TYPE_LIST
		0x00, 0x00, 0x00, 0x02, // restrict to XMMSV_TYPE_INT64
		0x00, 0x00, 0x00, 0x02, // 2 (number of list items)

		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x2a, // list[0]: 42

		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x17, // list[1]: 23
	})

	value, err := deserializeXmmsValue(buffer)
	require.NoError(t, err)

	list := value.(XmmsList)
	require.Len(t, list, 2)
	require.Equal(t, XmmsInt(42), list[0])
	require.Equal(t, XmmsInt(23), list[1])
}

func TestDeserializeDict(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x00, 0x07, // XMMSV_TYPE_DICT
		0x00, 0x00, 0x00, 0x02, // 2 (number of dict items)

		0x00, 0x00, 0x00, 0x04, // key[1]: 4 (length of following bytes)
		0x66, 0x6f, 0x6f, 0x00, // key[1]: "foo\0"

		0x00, 0x00, 0x00, 0x02, // value[1]: XMMSV_TYPE_INT64
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x2a, // value[1]: 42

		0x00, 0x00, 0x00, 0x04, // key[0]: 4 (length of following bytes)
		0x62, 0x61, 0x72, 0x00, // key[0]: "bar\0"

		0x00, 0x00, 0x00, 0x02, // value[0]: XMMSV_TYPE_INT64
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x25, 0xc3, // value[0]: 9667
	})

	value, err := deserializeXmmsValue(buffer)
	require.NoError(t, err)

	dict := value.(XmmsDict)
	require.Len(t, dict, 2)
	require.Equal(t, XmmsInt(42), dict["foo"])
	require.Equal(t, XmmsInt(9667), dict["bar"])
}

func TestDeserializeColl(t *testing.T) {
	buffer := bytes.NewBuffer([]byte{
		0x00, 0x00, 0x00, 0x04, /* XMMSV_TYPE_COLL */
		0x00, 0x00, 0x00, 0x06, /* XMMS_COLLECTION_TYPE_MATCH */
		0x00, 0x00, 0x00, 0x03, /* number of attributes*/

		0x00, 0x00, 0x00, 0x05, /* attr[0] key length */
		0x73, 0x65, 0x65, 0x64, /* attr[0] key "seed" */
		0x00, /*             "\0"   */

		0x00, 0x00, 0x00, 0x02, /* attr[0] value type  */
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x7a, 0x69, /* attr[0] value 31337 */

		0x00, 0x00, 0x00, 0x06, /* attr[1] key length */
		0x66, 0x69, 0x65, 0x6c, /* attr[1] key "fiel" */
		0x64, 0x00, /*              "d\0" */

		0x00, 0x00, 0x00, 0x03, /* attr[1] value type   */
		0x00, 0x00, 0x00, 0x07, /* attr[1] value length */
		0x61, 0x72, 0x74, 0x69, /* attr[1] value "arti" */
		0x73, 0x74, 0x00, /*               "st\0" */

		0x00, 0x00, 0x00, 0x06, /* attr[2] key length */
		0x76, 0x61, 0x6c, 0x75, /* attr[2] key "valu" */
		0x65, 0x00, /*             "e\0" */

		0x00, 0x00, 0x00, 0x03, /* attr[2] value type   */
		0x00, 0x00, 0x00, 0x0c, /* attr[2] value length */
		0x2a, 0x73, 0x65, 0x6e, /* attr[2] value "*sen"*/
		0x74, 0x65, 0x6e, 0x63, /*               "tenc" */
		0x65, 0x64, 0x2a, 0x00, /*               "ed*\0" */

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
	})

	value, err := deserializeXmmsValue(buffer)
	require.NoError(t, err)

	coll := value.(XmmsColl)
	require.Equal(t, CollectionTypeMatch, int(coll.Type))

	require.Len(t, coll.Attributes, 3)
	require.Equal(t, XmmsInt(31337), coll.Attributes["seed"])
	require.Equal(t, XmmsString("artist"), coll.Attributes["field"])
	require.Equal(t, XmmsString("*sentenced*"), coll.Attributes["value"])

	require.Len(t, coll.Operands, 1)
	require.Equal(t, CollectionTypeUniverse, int(coll.Operands[0].Type))

	require.Len(t, coll.Operands[0].Attributes, 0)
	require.Len(t, coll.Operands[0].IDList, 0)
	require.Len(t, coll.Operands[0].Operands, 0)
}
