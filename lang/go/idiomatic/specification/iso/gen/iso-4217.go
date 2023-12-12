package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/dave/jennifer/jen"
	"github.com/gertd/go-pluralize"
)

// https://en.wikipedia.org/wiki/ISO_4217
// https://www.six-group.com/en/products-services/financial-information/data-standards.html
// https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml

type iso4217Input struct {
	Code      string
	Number    string
	Decimal   string
	Name      string
	Countries []string
}

func processIso4217() error {
	rootDir, err := getRootDir()

	if err != nil {
		return err
	}

	txtPath := pather.Dirs.Join(rootDir, "idiomatic/specification/iso/gen/iso-4217.txt")
	txtContent, err := os.ReadFile(txtPath)

	if err != nil {
		return err
	}

	titles := []string{"Code", "Number", "Decimal", "Name", "Countries"}
	pattern := regexp.MustCompile(`\((\w\w)\)`)

	inputs := []iso4217Input{}

	for _, line := range stringer.Split[string](string(txtContent), "\n") {
		fields := stringer.Split(line, "\t")

		if len(fields) < len(titles) {
			fmt.Printf("Skipping: [%v]", line)
			continue
		}

		input := iso4217Input{
			Code:      stringer.TrimSpace(fields[0]),
			Number:    stringer.TrimSpace(fields[1]),
			Decimal:   stringer.TrimSpace(fields[2]),
			Name:      stringer.TrimSpace(fields[3]),
			Countries: []string{},
		}

		rawCountries := fields[4]
		input.Countries = append(input.Countries, input.Code[0:2])
		countryCodes := slicer.Map(
			stringer.RemoveSymbols[string],
			pattern.FindAllString(rawCountries, -1)...,
		)
		input.Countries = append(input.Countries, countryCodes...)

		inputs = append(inputs, input)
	}

	err = writeYaml("idiomatic/specification/iso/gen/iso-4217.yaml", inputs)

	if err != nil {
		return err
	}

	// var iso iso4217Record

	// if err := parseIso4217(&iso); err != nil {
	// 	return err
	// }

	// iso.Table.Entries = iso.Table.Entries[:3]

	// var content string

	// if err := generateIso4217(iso, &content); err != nil {
	// 	return err
	// }

	// if err := writeFile("idiomatic/currency/currency.gen.go", []byte(content)); err != nil {
	// 	return err
	// }

	return nil
}

// /////////////////////////////////////////////////////////////////
//  Code gen
// /////////////////////////////////////////////////////////////////

func enumIso4217() error {
	return nil
}

func generateIso4217(iso iso4217Record, content *string) error {
	pc := pluralize.NewClient()
	name := "currency"
	infoId := fmt.Sprintf("%vInfo", stringer.UppercaseFirst(name))
	structId := pc.Plural(name)
	varId := stringer.UppercaseFirst(structId)

	file := jen.NewFile("currency")

	file.Type().Id(infoId).Struct(
		jen.Id("Code").String(),
		jen.Id("Country").String(),
		jen.Id("Name").String(),
		jen.Id("Number").Int(),
		jen.Id("DecimalPlaces").Int(),
	).Line()

	file.Type().Id(structId).StructFunc(func(g *jen.Group) {
		g.Id("code2Currency").Map(jen.String()).Id(infoId)
		g.Id("name2Currency").Map(jen.String()).Id(infoId)
		g.Id("number2Currency").Map(jen.Int()).Id(infoId)
		g.Id("country2Currency").Map(jen.String()).Id(infoId)

		for _, entry := range iso.Table.Entries {
			g.Id(entry.Code).Id(infoId)
		}
	})

	file.Var().Id(varId).Op("=").Id(structId).Values()

	file.Func().Id("init").Params().BlockFunc(func(init *jen.Group) {
		for _, record := range iso.Table.Entries {
			init.Id(varId).Dot(record.Code).Op("=").Id(infoId).Values(jen.Dict{
				jen.Id("Name"):    jen.Lit(record.Name),
				jen.Id("Code"):    jen.Lit(record.Code),
				jen.Id("Country"): jen.Lit(record.Country),
				jen.Id("Number"):  jen.Lit(record.Number),
			}).Line()
		}

		mapper := func(mapName string, index *jen.Statement, fn func(e iso4217Entry) any) {
			init.Id(varId).Dot(mapName).Op("=").Map(index).Id(infoId).Values(jen.DictFunc(func(d jen.Dict) {
				for _, entry := range iso.Table.Entries {
					d[jen.Lit(fn(entry))] = jen.Id(varId).Dot(entry.Code)
				}
			})).Line()
		}

		mapper("code2Currency", jen.String(), func(e iso4217Entry) any { return e.Code })
		mapper("name2Currency", jen.String(), func(e iso4217Entry) any { return e.Name })
		mapper("number2Currency", jen.Int(), func(e iso4217Entry) any { return e.Number })
		mapper("country2Currency", jen.String(), func(e iso4217Entry) any { return e.Country })
	})

	*content = fmt.Sprintf("%#v", file)

	return nil
}

// /////////////////////////////////////////////////////////////////
//  ISO 4217 parsing
// /////////////////////////////////////////////////////////////////

func parseIso4217(iso *iso4217Record) error {
	rootDir, err := getRootDir()

	if err != nil {
		return err
	}

	xmlFilename := pather.Dirs.Join(rootDir, "idiomatic/specification/iso/gen/list-one.xml")

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

type iso4217Record struct {
	XMLName   xml.Name     `xml:"ISO_4217"`
	Published string       `xml:"Pblshd,attr"`
	Table     iso4217Table `xml:"CcyTbl"`
}

type iso4217Table struct {
	XMLName xml.Name       `xml:"CcyTbl"`
	Entries []iso4217Entry `xml:"CcyNtry"`
}

type iso4217Entry struct {
	XMLName       xml.Name          `xml:"CcyNtry"`
	Country       string            `xml:"CtryNm"`
	Name          string            `xml:"CcyNm"`
	Code          string            `xml:"Ccy"`
	Number        int               `xml:"CcyNbr"`
	DecimalPlaces iso4217CcyMnrUnts `xml:"CcyMnrUnts"`
}

type iso4217CcyMnrUnts int64

func (t iso4217CcyMnrUnts) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	var s string

	switch {
	case t < 0:
		s = "N.A."
	default:
		s = strconv.FormatInt(int64(t), 10)
	}

	return e.EncodeElement(s, start)
}

func (t *iso4217CcyMnrUnts) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
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
			*t = iso4217CcyMnrUnts(i)
		}
	}

	return nil
}
