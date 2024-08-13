package id

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// https://evs.nci.nih.gov/ftp1/GENC/
// https://www.dni.gov/index.php/who-we-are/organizations/ic-cio/ic-cio-related-menus/ic-cio-related-links/ic-technical-specifications/geopolitical-entities-names-and-codes

// type GenCodeUpdateMessage messenger.Message[GenCode]

// GenCode Geopolitical Entities, Names, and CODES
type GenCode struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type genCodes struct {
	Values        []GenCode
	nameToCountry map[string]GenCode
	codeToCountry map[string]GenCode
}

func (t genCodes) ByName(s string) (GenCode, bool) {
	country, ok := t.nameToCountry[stringer.Lowercase(s)]
	return country, ok
}

func (t genCodes) ByCode(s string) (GenCode, bool) {
	country, ok := t.codeToCountry[stringer.Lowercase(s)]
	return country, ok
}

func (t genCodes) Lookup(s string) (GenCode, bool) {
	country, ok := t.ByName(s)

	if !ok {
		country, ok = t.ByCode(s)
	}

	return country, ok
}

func initGenCode(fields []string) error {
	country := GenCode{
		Name: fields[0],
		Code: fields[1],
	}

	countryLower := GenCode{
		Name: stringer.Lowercase(country.Name),
		Code: stringer.Lowercase(country.Code),
	}

	var include bool

	if _, ok := GenCodes.nameToCountry[countryLower.Name]; !ok {
		GenCodes.nameToCountry[countryLower.Name] = country
		include = true
	}

	if country.Code != "" {
		if _, ok := GenCodes.codeToCountry[countryLower.Code]; !ok {
			GenCodes.codeToCountry[countryLower.Code] = country
			include = true
		}
	}

	if include {
		GenCodes.Values = append(GenCodes.Values, country)
	}

	return nil
}

var (
	GenCodes = genCodes{
		nameToCountry: map[string]GenCode{},
		codeToCountry: map[string]GenCode{},
	}
)
