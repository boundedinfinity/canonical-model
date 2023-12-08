package main

// https://www.six-group.com/en/products-services/financial-information/data-standards.html
// https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml

import (
	"encoding/xml"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

type iso4217 struct {
	XMLName   xml.Name `xml:"ISO_4217"`
	Published string   `xml:"Pblshd,attr"`
	Table     table    `xml:"CcyTbl"`
}

type table struct {
	XMLName xml.Name `xml:"CcyTbl"`
	Entries []entry  `xml:"CcyNtry"`
}

type entry struct {
	XMLName            xml.Name   `xml:"CcyNtry"`
	CountryName        string     `xml:"CtryNm"`
	CurrencyName       string     `xml:"CcyNm"`
	CurrencyCode       string     `xml:"Ccy"`
	CurrencyNumber     int        `xml:"CcyNbr"`
	CurrencyMinorUnits CcyMnrUnts `xml:"CcyMnrUnts"`
}

type CcyMnrUnts int64

func (t CcyMnrUnts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var s string

	switch {
	case t < 0:
		s = "N.A."
	default:
		s = strconv.FormatInt(int64(t), 10)
	}

	return e.EncodeElement(s, start)
}

func (t *CcyMnrUnts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string

	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	switch s {
	case "N.A.":
		*t = -1
	default:
		if i, err := strconv.ParseInt(s, 10, 64); err != nil {
			return err
		} else {
			*t = CcyMnrUnts(i)
		}
	}

	return nil
}

func main() {
	fmt.Println("ISO Currency Code Generator")

	// https://stackoverflow.com/a/70050843
	_, filename, _, _ := runtime.Caller(0)
	dir := pather.Files.Dir(filename)
	xmlFilename := pather.Join(dir, "list-one.xml")

	fmt.Printf("Input XML file: %v\n", xmlFilename)
	xmlData, err := os.ReadFile(xmlFilename)

	if err != nil {
		panic(err)
	}

	var iso iso4217

	if err := xml.Unmarshal(xmlData, &iso); err != nil {
		panic(err)
	}

	for _, x := range iso.Table.Entries {
		fmt.Printf("%v\n", x)
	}
}
