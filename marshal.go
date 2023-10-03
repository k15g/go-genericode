package genericode

import (
	"encoding/xml"
	"regexp"
	"strings"
)

var marshalPrefix = []byte("<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<gc:CodeList xmlns:gc=\"http://docs.oasis-open.org/codelist/ns/genericode/1.0/\">")
var marshalPostfix = []byte("</gc:CodeList>")
var marshalSelfClosing = regexp.MustCompile(`></.+>`)

func Marshal(codeList *CodeList) ([]byte, error) {
	content, err := xml.MarshalIndent(codeList, "", "\t")
	if err != nil {
		return nil, err
	}

	var result []byte
	result = append(result, marshalPrefix...)
	result = append(result, content[73:len(content)-11]...)
	result = append(result, marshalPostfix...)

	return []byte(marshalSelfClosing.ReplaceAllString(
			strings.ReplaceAll(string(result), "&#xA;", "\n"),
			" />")),
		nil
}
