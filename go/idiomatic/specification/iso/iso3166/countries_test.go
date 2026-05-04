package iso3166_test

import (
	"testing"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/specification/iso/iso3166"
	"github.com/stretchr/testify/assert"
)

func Test_Countries_FindOk(t *testing.T) {
	for _, country := range iso3166.Countries.All() {
		found, ok := iso3166.Countries.FindOk(country.Name)
		assert.Equal(t, ok, true, "expected to find country with name %s, but it was not found", country.Name)
		assert.Equal(t, 1, len(found), "expected to find exactly one country with name %s", country.Name)
		assert.Equal(t, country.Name, found[0].Name, "expected to find country with name %s, but found %s", country.Name, found[0].Name)
	}

	for _, country := range iso3166.Countries.All() {
		found, ok := iso3166.Countries.FindOk(country.Alpha2)
		assert.Equal(t, ok, true, "expected to find country with alpha2 %s, but it was not found", country.Alpha2)
		assert.Equal(t, 1, len(found), "expected to find exactly one country with alpha2 %s", country.Alpha2)
		assert.Equal(t, country.Name, found[0].Name, "expected to find country with alpha2 %s, but found %s", country.Alpha2, found[0].Name)
	}

	for _, country := range iso3166.Countries.All() {
		found, ok := iso3166.Countries.FindOk(country.Alpha3)
		assert.Equal(t, ok, true, "expected to find country with alpha3 %s, but it was not found", country.Alpha3)
		assert.Equal(t, 1, len(found), "expected to find exactly one country with alpha3 %s", country.Alpha3)
		assert.Equal(t, country.Name, found[0].Name, "expected to find country with alpha3 %s, but found %s", country.Alpha3, found[0].Name)
	}

	for _, country := range iso3166.Countries.All() {
		found, ok := iso3166.Countries.FindOk(country.Numeric)
		assert.Equal(t, ok, true, "expected to find country with numeric %s, but it was not found", country.Numeric)
		assert.Equal(t, 1, len(found), "expected to find exactly one country with numeric %s", country.Numeric)
		assert.Equal(t, country.Name, found[0].Name, "expected to find country with numeric %s, but found %s", country.Numeric, found[0].Name)
	}
}
