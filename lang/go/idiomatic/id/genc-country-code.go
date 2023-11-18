package id

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

// Geopolitical Entities, Names, and CODES
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
	country, ok := t.nameToCountry[stringer.ToLower(s)]
	return country, ok
}

func (t genCodes) ByCode(s string) (GenCode, bool) {
	country, ok := t.codeToCountry[stringer.ToLower(s)]
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
		Name: stringer.ToLower(country.Name),
		Code: stringer.ToLower(country.Code),
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
