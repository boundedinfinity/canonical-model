package main

import (
	"fmt"

	"github.com/boundedinfinity/enumer"
	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/dave/jennifer/jen"
	"github.com/gertd/go-pluralize"
)

// https://en.wikipedia.org/wiki/ISO_4217
// https://www.six-group.com/en/products-services/financial-information/data-standards.html
// https://www.six-group.com/dam/download/financial-information/data-center/iso-currrency/lists/list-one.xml

type iso4217Record struct {
	Code      string
	Number    int
	Decimal   int
	Name      string
	Countries []string
}

func processIso4217() error {
	var iso4217Inputs []iso4217Record

	if err := loadYaml("idiomatic/specification/iso/gen/iso-4217.yaml", &iso4217Inputs); err != nil {
		return err
	}

	var iso3166Inputs []iso3166Record

	if err := loadYaml("idiomatic/specification/iso/gen/iso-3166.yaml", &iso3166Inputs); err != nil {
		return err
	}

	iso4217Inputs = checkIso4217Country(iso4217Inputs, iso3166Inputs)

	if err := generateIso4217Enum(iso4217Inputs); err != nil {
		return err
	}

	// inputs = inputs[:3]

	if err := generateIso4217(iso4217Inputs); err != nil {
		return err
	}

	return nil
}

// /////////////////////////////////////////////////////////////////
//  Code gen
// /////////////////////////////////////////////////////////////////

func checkIso4217Country(iso4217Inputs []iso4217Record, iso3166Inputs []iso3166Record) []iso4217Record {
	var outputs []iso4217Record
	validCountries := slicer.Map(func(record iso3166Record) string {
		return record.Alpha2
	}, iso3166Inputs...)

	validCountries = slicer.Dedup(validCountries...)

	for _, input := range iso4217Inputs {
		var codes []string

		for _, code := range input.Countries {
			if slicer.Contains(code, validCountries...) {
				codes = append(codes, code)
			}
		}

		input.Countries = codes
		outputs = append(outputs, input)
	}

	outputs = slicer.DedupFn(func(r iso4217Record) int {
		return r.Number
	}, outputs...)

	outputs = slicer.DedupFn(func(r iso4217Record) string {
		return r.Code
	}, outputs...)

	return outputs
}

func generateIso4217Enum(inputs []iso4217Record) error {
	var enum enumer.EnumData

	for _, input := range inputs {
		enum.Values = append(enum.Values, enumer.EnumValue{
			Name:       input.Code,
			Serialized: input.Code,
		})
	}

	if err := writeYaml("idiomatic/currency/currency-code.enum.yaml", enum); err != nil {
		return err
	}

	return nil
}

func generateIso4217(inputs []iso4217Record) error {
	pc := pluralize.NewClient()
	name := "currency"
	infoId := fmt.Sprintf("%vInfo", stringer.UppercaseFirst(name))
	structId := pc.Plural(name)
	varId := stringer.UppercaseFirst(structId)

	file := jen.NewFile("currency")

	file.Type().Id(infoId).Struct(
		jen.Id("Code").Id("CurrencyCode"),
		jen.Id("CountryCodes").Index().Qual("github.com/boundedinfinity/schema/idiomatic/location", "CountryAlpha2"),
		jen.Id("Name").String(),
		jen.Id("Number").Int(),
		jen.Id("DecimalPlaces").Int(),
	).Line()

	file.Type().Id(structId).StructFunc(func(g *jen.Group) {
		g.Id("codeToCurrency").Map(jen.Id("CurrencyCode")).Id(infoId)
		g.Id("nameToCurrency").Map(jen.String()).Id(infoId)
		g.Id("numberToCurrency").Map(jen.Int()).Id(infoId)
		g.Id("countryCodeToCurrency").Map(jen.Qual("github.com/boundedinfinity/schema/idiomatic/location", "CountryAlpha2")).Id(infoId)

		for _, input := range inputs {
			g.Id(input.Code).Id(infoId)
		}
	})

	file.Var().Id(varId).Op("=").Id(structId).Values()

	file.Func().Id("init").Params().BlockFunc(func(init *jen.Group) {
		for _, input := range inputs {
			init.Id(varId).Dot(input.Code).Op("=").Id(infoId).Values(jen.Dict{
				jen.Id("Name"): jen.Lit(input.Name),
				jen.Id("Code"): jen.Id("CurrencyCodes").Dot(input.Code),
				jen.Id("CountryCodes"): jen.Index().Qual("github.com/boundedinfinity/schema/idiomatic/location", "CountryAlpha2").ValuesFunc(func(g *jen.Group) {
					for _, code := range input.Countries {
						g.Line().Id("location").Dot("CountryAlpha2s").Dot(code)
					}
					g.Line()
				}),
				jen.Id("Number"): jen.Lit(input.Number),
			}).Line()
		}

		mapper := func(mapName string, index *jen.Statement, fn func(e iso4217Record) any) {
			init.Id(varId).Dot(mapName).Op("=").Map(index).Id(infoId).Values(jen.DictFunc(func(d jen.Dict) {
				for _, input := range inputs {
					val := fn(input)

					if !reflecter.Instances.IsZero(val) {
						d[jen.Lit(val)] = jen.Id(varId).Dot(input.Code)
					}

				}
			})).Line()
		}

		mapper("nameToCurrency", jen.String(), func(e iso4217Record) any { return e.Name })
		mapper("numberToCurrency", jen.Int(), func(e iso4217Record) any { return e.Number })

		init.Id(varId).Dot("codeToCurrency").Op("=").Map(jen.Id("CurrencyCode")).Id(infoId).Values(jen.DictFunc(func(d jen.Dict) {
			for _, input := range inputs {
				d[jen.Id("CurrencyCodes").Dot(input.Code)] = jen.Id(varId).Dot(input.Code)
			}
		})).Line()

		init.Id(varId).Dot("countryCodeToCurrency").Op("=").Map(jen.Qual("github.com/boundedinfinity/schema/idiomatic/location", "CountryAlpha2")).Id(infoId).Values(jen.DictFunc(func(d jen.Dict) {
			for _, input := range inputs {
				for _, code := range input.Countries {
					d[jen.Id("location").Dot("CountryAlpha2s").Dot(code)] = jen.Id(varId).Dot(input.Code)
				}
			}
		}))
	})

	content := fmt.Sprintf("%#v", file)

	if err := writeFile("idiomatic/currency/currency.gen.go", []byte(content)); err != nil {
		return err
	}

	return nil
}
