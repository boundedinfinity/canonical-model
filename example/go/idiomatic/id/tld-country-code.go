package id

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

// Top Level Domain
type CountryTld struct {
	Name string `json:"name,omitempty"`
	Tld  string `json:"tld,omitempty"`
}

type countryTlds struct {
	Values        []CountryTld
	nameToCountry map[string]CountryTld
	tldToCountry  map[string]CountryTld
}

func (t countryTlds) ByName(s string) (CountryTld, bool) {
	country, ok := t.nameToCountry[stringer.Lowercase(s)]
	return country, ok
}

func (t countryTlds) ByTld(s string) (CountryTld, bool) {
	country, ok := t.tldToCountry[stringer.Lowercase(s)]
	return country, ok
}

var (
	CountryTlds = countryTlds{
		nameToCountry: map[string]CountryTld{},
		tldToCountry:  map[string]CountryTld{},
	}
)

func initCountryTld(fields []string) error {
	country := CountryTld{
		Name: fields[0],
		Tld:  fields[6],
	}

	countryLower := CountryTld{
		Name: stringer.Lowercase(country.Name),
		Tld:  stringer.Lowercase(country.Tld),
	}

	if _, ok := CountryTlds.nameToCountry[countryLower.Name]; !ok {
		CountryTlds.nameToCountry[countryLower.Name] = country
	}

	CountryTlds.Values = append(CountryTlds.Values, country)

	return nil
}
