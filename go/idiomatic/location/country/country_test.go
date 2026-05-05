package country_test

import (
	"testing"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/location/country"
	"github.com/stretchr/testify/assert"
)

func Test_Country_ByName(t *testing.T) {
	var found []country.Country
	var ok bool
	for _, entry := range country.Countries.All() {
		found, ok = country.Countries.FindOk(entry.Name)
		assert.Equal(t, ok, true, "expected name %s, but it was not found", entry.Name)
		assert.Equal(t, 1, len(found), "expected one with name %s", entry.Name)
		assert.Equal(t, entry.Name, found[0].Name,
			"expected name %s, but found %s", entry.Name, found[0].Name)

		for _, alias := range entry.Aliases {
			found, ok = country.Countries.FindOk(alias)
			assert.Equal(t, ok, true, "expected name%s with alias[%s], but it was not found", entry.Name, alias)
			assert.Equal(t, 1, len(found), "expected one with name %s with alias[%s]", entry.Name, alias)
			assert.Equal(t, entry.Name, found[0].Name,
				"expected name %s with alias[%s], but found %s", entry.Name, alias, found[0].Name)
		}

		found, ok = country.Countries.FindOk(entry.Iso.Alpha2)
		assert.Equal(t, ok, true,
			"expected name %s, ISO alpha 2 %s, but it was not found", entry.Name, entry.Iso.Alpha2)
		assert.Equal(t, 1, len(found),
			"expected one with name %s, ISO alpha 2 %s", entry.Name, entry.Iso.Alpha2)
		assert.Equal(t, entry.Name, found[0].Name,
			"expected `name %s, ISO alpha 2 %s, but found %s", entry.Name, entry.Iso.Alpha2, found[0].Name)

		found, ok = country.Countries.FindOk(entry.Iso.Alpha3)
		assert.Equal(t, ok, true,
			"expected name %s, ISO alpha 3 %s, but it was not found", entry.Name, entry.Iso.Alpha3)
		assert.Equal(t, 1, len(found),
			"expected one with name %s, ISO alpha 3 %s", entry.Name, entry.Iso.Alpha3)
		assert.Equal(t, entry.Name, found[0].Name,
			"expected `name %s, ISO alpha 3 %s, but found %s", entry.Name, entry.Iso.Alpha3, found[0].Name)

		found, ok = country.Countries.FindOk(entry.Iso.Numeric)
		assert.Equal(t, ok, true,
			"expected name %s, ISO numeric %s, but it was not found", entry.Name, entry.Iso.Numeric)
		assert.Equal(t, 1, len(found),
			"expected one with name %s, ISO numeric %s", entry.Name, entry.Iso.Numeric)
		assert.Equal(t, entry.Name, found[0].Name,
			"expected `name %s, ISO numeric %s, but found %s", entry.Name, entry.Iso.Numeric, found[0].Name)

		found, ok := country.Countries.FindOk(entry.Iso.Name)
		assert.Equal(t, ok, true,
			"expected name %s, ISO name %s, but it was not found", entry.Name, entry.Iso.Name)
		assert.Equal(t, 1, len(found),
			"expected one with name %s, ISO name %s", entry.Name, entry.Iso.Name)
		assert.Equal(t, entry.Name, found[0].Name,
			"expected `name %s, but found %s", entry.Name, found[0].Name)

	}
}
