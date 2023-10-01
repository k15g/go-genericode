package genericode

import (
	"encoding/xml"
	"os"
)

func UnmarshalFile(path string) (*CodeList, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return Unmarshal(data)
}

func Unmarshal(data []byte) (*CodeList, error) {
	var codelist CodeList
	if err := xml.Unmarshal(data, &codelist); err != nil {
		return nil, err
	}

	return &codelist, nil
}
