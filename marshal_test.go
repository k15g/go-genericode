package genericode

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalSimple(t *testing.T) {
	cl, err := UnmarshalFile("testdata/peppol-icd.gc")
	assert.Nil(t, err)

	content, err := Marshal(cl)
	assert.Nil(t, err)

	expected, _ := os.ReadFile("testdata/peppol-icd-exported.gc")

	assert.Equal(t, expected, content)

	if writeDuringTests {
		os.WriteFile("testdata/peppol-icd-exported.gc", content, os.ModePerm)
	}
}
