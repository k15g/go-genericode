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
	// Unmarshal data
	var codelist CodeList
	if err := xml.Unmarshal(data, &codelist); err != nil {
		return nil, err
	}

	// Injecting reference to codelist in column definitions and index definitions
	codelist.columnIndex = map[string]*Column{}
	for _, column := range codelist.Columns {
		column.codeList = &codelist
		codelist.columnIndex[*column.Id] = column
	}

	// Injecting reference to codelist in key definitions and index definitions
	codelist.keyIndex = map[string]*Key{}
	for _, key := range codelist.Keys {
		key.codeList = &codelist
		codelist.keyIndex[*key.Id] = key
	}

	return &codelist, nil
}
