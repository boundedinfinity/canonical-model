package location_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/location"
	"github.com/stretchr/testify/assert"
)

func Test_Country_Flag(t *testing.T) {
	content, err := location.Countries.Flag("US")

	assert.Nil(t, err)
	assert.Contains(t, string(content), "#us-a")
}

func Test_Country_ByName(t *testing.T) {
	actual := location.Countries.ByName("Cook")
	expected := []location.CountryInfo{location.Countries.CK}

	assert.Equal(t, expected, actual)
}
