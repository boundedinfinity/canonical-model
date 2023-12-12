package location

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func (t countries) ByAlpha2(code string) (CountryInfo, bool) {
	found, ok := t.alpha2ToInfo[CountryAlpha2(code)]
	return found, ok
}

func (t countries) ByAlpha3(code string) (CountryInfo, bool) {
	found, ok := t.alpha3ToInfo[CountryAlpha3(code)]
	return found, ok
}

func (t countries) ByNumber(number int) (CountryInfo, bool) {
	found, ok := t.numberToInfo[number]
	return found, ok
}

func (t countries) ByName(term string) []CountryInfo {
	found := slicer.Filter(func(country CountryInfo) bool {
		return stringer.ContainsIgnoreCase(country.Name, term)
	}, t.All...)

	return found
}

func (t countries) Lookup(term string) (CountryInfo, bool) {
	var found CountryInfo
	var ok bool

	found, ok = t.ByAlpha2(term)

	if !ok {
		found, ok = t.ByAlpha3(term)
	}

	if !ok {
		i64, err := strconv.ParseInt(term, 10, 64)
		if err == nil {
			found, ok = t.ByNumber(int(i64))
		}
	}

	return found, ok
}

var (
	ErrCountryFlagNotFound  = errors.New("flag not found")
	ErrCountryFlagNotFoundv = errorer.Wrapv(ErrCountryFlagNotFound)
)

func (t countries) Flag(term string) ([]byte, error) {
	info, ok := t.Lookup(term)

	if !ok {
		return nil, ErrCountryFlagNotFoundv(info)
	}

	path := fmt.Sprintf("flags/%v.svg", info.Alpha2.String())
	data, err := flagsFS.ReadFile(path)
	return data, err
}
