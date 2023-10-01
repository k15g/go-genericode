package genericode

import "encoding/xml"

func Marshal(codeList any) ([]byte, error) {
	// TODO Fix "gc" namespace
	// TODO Add XML header
	// TODO Fix self-closing tags

	return xml.Marshal(codeList)
}
