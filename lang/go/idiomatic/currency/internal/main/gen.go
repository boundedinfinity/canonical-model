package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"runtime"
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/dave/jennifer/jen"
)

func main() {
	fmt.Println("ISO Currency Code Generator")

	var iso iso4217

	if err := parse(&iso); err != nil {
		panic(err)
	}

	var content string

	if err := generate(iso, &content); err != nil {
		panic(err)
	}

	fmt.Println(content)
}

// /////////////////////////////////////////////////////////////////
//  Code gen
// /////////////////////////////////////////////////////////////////

func generate(iso iso4217, content *string) error {
	typeName := "IsoCurrency"
	stop := 2

	file := jen.NewFile("main")
	// file.Type().Id(typeName).String()

	file.Type().Id(typeName).Struct(
		jen.Id("Code").String(),
		jen.Id("Country").String(),
		jen.Id("Name").String(),
		jen.Id("Number").Int(),
		jen.Id("DecimalPlaces").Int(),
	)

	mapper := func(mapName string, fn func(e entry) any) {
		file.Var().Id(mapName).Op("=").Map(jen.String()).Id(typeName).Values(jen.DictFunc(func(d jen.Dict) {
			for i, entry := range iso.Table.Entries {
				if i > stop {
					break
				}

				d[jen.Lit(fn(entry))] = jen.Id(entry.Code)
			}
		}))
	}

	mapper("code2Currency", func(e entry) any { return e.Code })
	mapper("name2Currency", func(e entry) any { return e.Name })
	mapper("number2Currency", func(e entry) any { return e.Number })
	mapper("country2Currency", func(e entry) any { return e.Country })

	for i, entry := range iso.Table.Entries {
		if entry.Code == "" {
			continue
		}

		if i > stop {
			break
		}

		file.Var().Id(entry.Code).Op("=").Id(typeName).Values(jen.Dict{
			jen.Id("Code"):    jen.Lit(entry.Code),
			jen.Id("Number"):  jen.Lit(entry.Number),
			jen.Id("Country"): jen.Lit(entry.Country),
		})
	}

	*content = fmt.Sprintf("%#v", file)

	return nil
}

// /////////////////////////////////////////////////////////////////
//  ISO 4217 parsing
// /////////////////////////////////////////////////////////////////

func parse(iso *iso4217) error {
	// https://stackoverflow.com/a/70050843
	_, filename, _, ok := runtime.Caller(0)

	if !ok {
		return errors.New("can't determine current directory")
	}

	dir := pather.Files.Dir(filename)
	xmlFilename := pather.Join(dir, "list-one.xml")

	fmt.Printf("Input XML file: %v\n", xmlFilename)
	xmlData, err := os.ReadFile(xmlFilename)

	if err != nil {
		return err
	}

	if err := xml.Unmarshal(xmlData, iso); err != nil {
		return err
	}

	return nil
}

// https://www.six-group.com/en/products-services/financial-information/data-standards.html
// https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml

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
	XMLName       xml.Name   `xml:"CcyNtry"`
	Country       string     `xml:"CtryNm"`
	Name          string     `xml:"CcyNm"`
	Code          string     `xml:"Ccy"`
	Number        int        `xml:"CcyNbr"`
	DecimalPlaces CcyMnrUnts `xml:"CcyMnrUnts"`
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
