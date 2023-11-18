package id

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

type StanagCode struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}

type stanagCodes struct {
	Values        []StanagCode
	nameToCountry map[string][]StanagCode
	codeToCountry map[string][]StanagCode
}

func (t stanagCodes) ByName(s string) ([]StanagCode, bool) {
	countries, ok := t.nameToCountry[stringer.ToLower(s)]
	return countries, ok && len(countries) > 0
}

func (t stanagCodes) ByCode(s string) ([]StanagCode, bool) {
	countries, ok := t.codeToCountry[stringer.ToLower(s)]
	return countries, ok && len(countries) > 0
}

func (t stanagCodes) Lookup(s string) ([]StanagCode, bool) {
	countries, ok := t.ByName(s)

	if !ok && len(countries) == 0 {
		countries, ok = t.ByCode(s)
	}

	return countries, ok && len(countries) > 0
}

var (
	StanagCodes = stanagCodes{
		nameToCountry: map[string][]StanagCode{},
		codeToCountry: map[string][]StanagCode{},
	}
)

func initStanag(fields []string) error {
	country := StanagCode{
		Name: fields[0],
		Code: fields[5],
	}

	countryLower := StanagCode{
		Name: stringer.ToLower(country.Name),
		Code: stringer.ToLower(country.Code),
	}

	var include bool

	if _, ok := GenCodes.nameToCountry[countryLower.Name]; !ok {
		StanagCodes.nameToCountry[countryLower.Name] = []StanagCode{}

		StanagCodes.nameToCountry[countryLower.Name] = append(StanagCodes.nameToCountry[countryLower.Name], country)
		include = true
	}

	if country.Code != "" {
		if _, ok := GenCodes.codeToCountry[countryLower.Code]; !ok {
			StanagCodes.codeToCountry[countryLower.Code] = []StanagCode{}
			StanagCodes.codeToCountry[countryLower.Name] = append(StanagCodes.codeToCountry[countryLower.Name], country)
			include = true
		}
	}

	if include {
		StanagCodes.Values = append(StanagCodes.Values, country)
	}

	return nil
}
