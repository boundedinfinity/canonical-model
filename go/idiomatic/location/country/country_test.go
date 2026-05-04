package country_test

import (
	"testing"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/location/country"
	"github.com/stretchr/testify/assert"
)

func Test_Country_ByName(t *testing.T) {
	for _, entry := range country.Countries.All() {
		found, ok := country.Countries.FindOk(entry.Name)
		assert.Equal(t, ok, true, "expected to find country with name %s, but it was not found", entry.Name)
		assert.Equal(t, 1, len(found), "expected to find exactly one country with name %s", entry.Name)
		assert.Equal(t, entry.Name, found[0].Name, "expected to find country with name %s, but found %s", entry.Name, found[0].Name)
	}
}
