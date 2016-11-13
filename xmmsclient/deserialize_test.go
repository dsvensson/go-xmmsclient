package xmmsclient

import (
	"bytes"
	"testing"
)

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

	value, err := DeserializeXmmsValue(buffer)
	if err != nil {
		t.Fatal("deserialization failed")
	}

	coll := value.(XmmsColl)

	if coll.Type != CollectionTypeMatch {
		t.Fatal("wrong collection type")
	}

	if len(coll.Attributes) != 3 {
		t.Fatal("wrong attributes count")
	}

	if coll.Attributes["seed"] != XmmsInt(31337) {
		t.Fatal("wrong attribute, seed != 31337")
	}

	if coll.Attributes["field"] != XmmsString("artist") {
		t.Fatal("wrong attribute, field != artist")
	}

	if coll.Attributes["value"] != XmmsString("*sentenced*") {
		t.Fatal("wrong attribute, field != artist")
	}

	if len(coll.Operands.Entries) != 1 {
		t.Fatal("wrong attributes count")
	}

	operand := coll.Operands.Entries[0].(XmmsColl)

	if operand.Type != CollectionTypeUniverse {
		t.Fatal("wrong collection type")
	}

	if len(operand.Attributes) != 0 {
		t.Fatal("wrong attributes count")
	}

	if len(operand.IdList.Entries) != 0 {
		t.Fatal("wrong idlist size")
	}

	if len(operand.Operands.Entries) != 0 {
		t.Fatal("wrong operands count")
	}

}
