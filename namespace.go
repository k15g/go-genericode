package genericode

import (
	"encoding/xml"
)

type CodeList struct {
	XMLName        xml.Name `xml:"http://docs.oasis-open.org/codelist/ns/genericode/1.0/ CodeList"`
	Identification *Identification
	Columns        []*Column `xml:"ColumnSet>Column"`
	Keys           []*Key    `xml:"ColumnSet>Key"`
	Rows           []*Row    `xml:"SimpleCodeList>Row"`

	columnIndex map[string]*Column
	keyIndex    map[string]*Key
}

func (cl *CodeList) Column(id string) *Column {
	if cl.columnIndex != nil {
		return cl.columnIndex[id]
	}

	// Fallback if index is not provided
	for _, column := range cl.Columns {
		if *column.Id == id {
			return column
		}
	}

	return nil
}

type Identification struct {
	ShortName           []ShortName `xml:"ShortName"`
	Version             *string     `xml:"Version"`
	CanonicalUri        *string     `xml:"CanonicalUri"`
	CanonicalVersionUri *string     `xml:"CanonicalVersionUri"`
}

type Column struct {
	Id        *string      `xml:"Id,attr"`
	Use       *string      `xml:"Use,attr"`
	ShortName []*ShortName `xml:"ShortName"`
	Data      *Data        `xml:"Data"`

	codeList *CodeList
}

type Key struct {
	Id        *string      `xml:"Id,attr"`
	ShortName []*ShortName `xml:"ShortName"`
	ColumnRef ColumnRef    `xml:"ColumnRef"`

	codeList *CodeList
}

type Row struct {
	Values []Value `xml:"Value"`
}

func (row *Row) Get(key string) *string {
	for _, value := range row.Values {
		if *value.ColumnRef == key {
			return value.Value
		}
	}

	return nil
}

type Value struct {
	ColumnRef *string `xml:"ColumnRef,attr"`
	Value     *string `xml:"SimpleValue"`
}

type ShortName struct {
	Value    string  `xml:",chardata"`
	Language *string `xml:"lang,attr"`
}

type Data struct {
	Type *string `xml:"Type,attr"`
}

type ColumnRef struct {
	Ref *string `xml:"Ref,attr"`
}
