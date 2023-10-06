package genericode

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const writeDuringTests = true

func TestUnmarshalPeppolICD(t *testing.T) {
	cl, err := UnmarshalFile("testdata/peppol-icd.gc")
	assert.Nil(t, err)
	assert.NotNil(t, cl)

	assert.Equal(t, *cl.Identification.Version, "8.6")
	assert.Equal(t, *cl.Columns[0].Id, "schemeid")
	assert.Equal(t, *cl.Keys[0].Id, "schemeidKey")
	assert.Equal(t, *cl.Rows[0].Values[0].SimpleValue, "FR:SIRENE")

	assert.NotNil(t, cl.Column("schemeid"))
	assert.Nil(t, cl.Column("schemeid-unknown"))

	assert.Equal(t, *cl.Rows[0].Get("schemeid"), "FR:SIRENE")
	assert.Nil(t, cl.Rows[0].Get("schemeid-unknown"))

	if writeDuringTests {
		content, _ := json.Marshal(cl)
		os.WriteFile("test.json", content, os.ModePerm)

		content, _ = Marshal(cl)
		os.WriteFile("test.xml", content, os.ModePerm)
	}
}

func TestUBLSignatureEntites(t *testing.T) {
	cl, err := UnmarshalFile("testdata/UBL-Signature-Entities-2.3.gc")
	assert.Nil(t, err)
	assert.NotNil(t, cl)

	assert.Equal(t, "en", *cl.Identification.Agency.LongName[0].Language)
	assert.Equal(t, "UBL", *cl.Identification.Agency.Identifier[0])

	if writeDuringTests {
		content, _ := Marshal(cl)
		os.WriteFile("UBL-Signature-Entities-2.3-exported.gc", content, os.ModePerm)
	}
}
