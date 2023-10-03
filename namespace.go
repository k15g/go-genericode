package genericode

import (
	"encoding/xml"
)

type Genericode struct {
}

type CodeList struct {
	Genericode

	XMLName        xml.Name        `xml:"http://docs.oasis-open.org/codelist/ns/genericode/1.0/ CodeList"`
	Identification *Identification `xml:"Identification"`
	Columns        []*Column       `xml:"ColumnSet>Column"`
	Keys           []*Key          `xml:"ColumnSet>Key"`
	Rows           []*Row          `xml:"SimpleCodeList>Row"`

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
	ShortName                  *TranslatableName   `xml:"ShortName"`
	LongName                   []*TranslatableName `xml:"LongName"`
	Version                    *string             `xml:"Version"`
	CanonicalUri               *string             `xml:"CanonicalUri"`
	CanonicalVersionUri        *string             `xml:"CanonicalVersionUri"`
	LocationUri                *string             `xml:"LocationUri"`
	AlternateFormatLocationUri *MimeTypedUri       `xml:"AlternateFormatLocationUri"`
	Agency                     *Agency             `xml:"Agency"`
}

type Agency struct {
	ShortName  *TranslatableName   `xml:"ShortName"`
	LongName   *[]TranslatableName `xml:"LongName"`
	Identifier *[]string           `xml:"Identifier"`
}

type Column struct {
	Id                  *string             `xml:"Id,attr"`  // Required
	Use                 *string             `xml:"Use,attr"` // Required
	ShortName           *TranslatableName   `xml:"ShortName"`
	LongName            []*TranslatableName `xml:"LongName"`
	CanonicalUri        *string             `xml:"CanonicalUri"`
	CanonicalVersionUri *string             `xml:"CanonicalVersionUri"`
	Data                *Data               `xml:"Data"` // Required

	codeList *CodeList
}

type Key struct {
	Id                  *string             `xml:"Id,attr"` // Required
	ShortName           *TranslatableName   `xml:"ShortName"`
	LongName            []*TranslatableName `xml:"LongName"`
	CanonicalUri        *string             `xml:"CanonicalUri"`
	CanonicalVersionUri *string             `xml:"CanonicalVersionUri"`
	ColumnRef           *[]ColumnRef        `xml:"ColumnRef"` // Required

	codeList *CodeList
}

type Row struct {
	Values []*Value `xml:"Value"`
}

func (row *Row) Get(column string) *string {
	for _, value := range row.Values {
		if *value.ColumnRef == column {
			return value.Value
		}
	}

	return nil
}

type Value struct {
	ColumnRef *string `xml:"ColumnRef,attr"`
	Value     *string `xml:"SimpleValue"`
}

type TranslatableName struct {
	Value    string  `xml:",chardata"`
	Language *string `xml:"lang,attr"`
}

type Data struct {
	Type            *string          `xml:"Type,attr"`
	DatatypeLibrary *string          `xml:"DatatypeLibrary,attr"`
	Language        *string          `xml:"Lang,attr"`
	Parameter       *[]DatatypeFacet `xml:"Parameter"`
}

type DatatypeFacet struct {
	Facet     string  `xml:",chardata"`      // Required
	ShortName *string `xml:"ShortName,attr"` // Required
	LongName  *string `xml:"LongName,attr"`
}

type ColumnRef struct {
	Ref *string `xml:"Ref,attr"`
}

type MimeTypedUri struct {
	URI      string  `xml:",chardata"` // Required
	MimeType *string `xml:"MimeType,attr"`
}
