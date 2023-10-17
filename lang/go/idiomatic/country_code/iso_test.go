package country_code_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/country_code"
	"github.com/stretchr/testify/assert"
)

func Test_CountryCodeLookup_Alpha2(t *testing.T) {
	code, ok := country_code.CountryCodeLookup("US")
	assert.Equal(t, true, ok)
	assert.Equal(t, "US", code.Alpha2)
	assert.Equal(t, "USA", code.Alpha3)
	assert.Equal(t, 840, code.Numeric)
}

func Test_CountryCodeLookup_Numeric(t *testing.T) {
	code, ok := country_code.CountryCodeLookup("840")
	assert.Equal(t, true, ok)
	assert.Equal(t, "US", code.Alpha2)
	assert.Equal(t, "USA", code.Alpha3)
	assert.Equal(t, 840, code.Numeric)
}
