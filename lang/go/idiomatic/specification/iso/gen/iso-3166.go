package main

import (
	"fmt"

	"github.com/boundedinfinity/enumer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/dave/jennifer/jen"
	"github.com/gertd/go-pluralize"
)

// https://en.wikipedia.org/wiki/ISO_3166-1#Codes
// ISO 3166-1 Name, Alpha-2 code, Alpha-3 code, Numeric code, Link to ISO 3166-2, Independent

type iso3166Record struct {
	Name        string `yaml:"name"`
	Alpha2      string `yaml:"alpha-2"`
	Alpha3      string `yaml:"alpha-3"`
	Number      int    `yaml:"number"`
	Independent string `yaml:"independent"`
}

func processIso3166() error {
	pc := pluralize.NewClient()
	var records []iso3166Record

	if err := loadYaml("idiomatic/specification/iso/gen/iso-3166.yaml", &records); err != nil {
		return err
	}

	records = slicer.DedupFn(func(r iso3166Record) int {
		return r.Number
	}, records...)

	records = slicer.DedupFn(func(r iso3166Record) string {
		return r.Alpha2
	}, records...)

	records = slicer.DedupFn(func(r iso3166Record) string {
		return r.Alpha3
	}, records...)

	// records = records[:3]

	name := "country"
	infoId := fmt.Sprintf("%vInfo", stringer.UppercaseFirst(name))
	alpha2Id := fmt.Sprintf("%vAlpha2", stringer.UppercaseFirst(name))
	alpha3Id := fmt.Sprintf("%vAlpha3", stringer.UppercaseFirst(name))
	structId := pc.Plural(name)
	varId := stringer.UppercaseFirst(structId)

	file := jen.NewFile("location")

	file.Type().Id(infoId).Struct(
		jen.Id("Name").String(),
		jen.Id("Alpha2").Id(alpha2Id),
		jen.Id("Alpha3").Id(alpha3Id),
		jen.Id("Number").Int(),
	).Line()

	file.Type().Id(structId).StructFunc(func(g *jen.Group) {
		g.Id("alpha2ToInfo").Map(jen.Id(alpha2Id)).Id(infoId)
		g.Id("alpha3ToInfo").Map(jen.Id(alpha3Id)).Id(infoId)
		g.Id("nameToInfo").Map(jen.String()).Id(infoId)
		g.Id("numberToInfo").Map(jen.Int()).Id(infoId)
		g.Id("All").Index().Id(infoId)

		for _, record := range records {
			g.Id(record.Alpha2).Id(infoId)
		}
	})

	file.Var().Id(varId).Op("=").Id(structId).Values()

	file.Func().Id("init").Params().BlockFunc(func(init *jen.Group) {
		for _, record := range records {
			init.Id(varId).Dot(record.Alpha2).Op("=").Id(infoId).Values(jen.Dict{
				jen.Id("Name"):   jen.Lit(record.Name),
				jen.Id("Alpha2"): jen.Lit(record.Alpha2),
				jen.Id("Alpha3"): jen.Lit(record.Alpha3),
				jen.Id("Number"): jen.Lit(record.Number),
			}).Line()
		}

		init.Id(varId).Dot("All").Op("=").Index().Id(infoId).ValuesFunc(func(g *jen.Group) {
			for _, record := range records {
				g.Line().Id(varId).Dot(record.Alpha2)
			}
		}).Line()

		init.Id(varId).Dot("alpha2ToInfo").Op("=").Map(jen.Id(alpha2Id)).Id(infoId).Values(jen.DictFunc(func(d jen.Dict) {
			for _, record := range records {
				d[jen.Id(pc.Plural(alpha2Id)).Dot(record.Alpha2)] = jen.Id(varId).Dot(record.Alpha2)
			}
		})).Line()

		init.Id(varId).Dot("alpha3ToInfo").Op("=").Map(jen.Id(alpha3Id)).Id(infoId).Values(jen.DictFunc(func(d jen.Dict) {
			for _, record := range records {
				if record.Alpha3 != "" {
					d[jen.Id(pc.Plural(alpha3Id)).Dot(record.Alpha3)] = jen.Id(varId).Dot(record.Alpha2)
				}
			}
		})).Line()

		init.Id(varId).Dot("nameToInfo").Op("=").Map(jen.String()).Id(infoId).Values(jen.DictFunc(func(d jen.Dict) {
			for _, record := range records {
				d[jen.Lit(stringer.ToLower(record.Name))] = jen.Id(varId).Dot(record.Alpha2)
			}
		})).Line()

		init.Id(varId).Dot("numberToInfo").Op("=").Map(jen.Int()).Id(infoId).Values(jen.DictFunc(func(d jen.Dict) {
			for _, record := range records {
				if record.Number > 0 {
					d[jen.Lit(record.Number)] = jen.Id(varId).Dot(record.Alpha2)
				}
			}
		})).Line()
	})

	content := fmt.Sprintf("%#v", file)

	if err := writeFile("idiomatic/location/country.gen.go", []byte(content)); err != nil {
		return err
	}

	alpha2Enum := enumer.EnumData{
		Values: []enumer.EnumValue{},
	}

	alpha3Enum := enumer.EnumData{
		Values: []enumer.EnumValue{},
	}

	for _, record := range records {
		alpha2Enum.Values = append(alpha2Enum.Values, enumer.EnumValue{
			Name:       record.Alpha2,
			Serialized: record.Alpha2,
		})

		if record.Alpha3 != "" {
			alpha3Enum.Values = append(alpha3Enum.Values, enumer.EnumValue{
				Name:       record.Alpha3,
				Serialized: record.Alpha3,
			})
		}
	}

	if err := writeYaml("idiomatic/location/country-alpha-2.enum.yaml", alpha2Enum); err != nil {
		return err
	}

	if err := writeYaml("idiomatic/location/country-alpha-3.enum.yaml", alpha3Enum); err != nil {
		return err
	}

	return nil
}
