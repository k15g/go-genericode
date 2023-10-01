package genericode

import "encoding/xml"

type CodeList struct {
	XMLName        xml.Name `xml:"http://docs.oasis-open.org/codelist/ns/genericode/1.0/ CodeList"`
	Identification *Identification
	Columns        []Column `xml:"ColumnSet>Column"`
	Keys           []Key    `xml:"ColumnSet>Key"`
	Rows           []Row    `xml:"SimpleCodeList>Row"`
}

type Identification struct {
	ShortName           []ShortName `xml:"ShortName"`
	Version             *string     `xml:"Version"`
	CanonicalUri        *string     `xml:"CanonicalUri"`
	CanonicalVersionUri *string     `xml:"CanonicalVersionUri"`
}

type Column struct {
	Id        *string `xml:"Id,attr"`
	Use       *string `xml:"Use,attr"`
	ShortName *string `xml:"ShortName"`
	Data      *Data   `xml:"Data"`
}

type Key struct {
	Id        *string     `xml:"Id,attr"`
	ShortName []ShortName `xml:"ShortName"`
	ColumnRef ColumnRef   `xml:"ColumnRef"`
}

type Row struct {
	Values []Value `xml:"Value"`
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
