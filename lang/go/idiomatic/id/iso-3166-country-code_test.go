package id_test

import (
	"fmt"
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/stretchr/testify/assert"
)

func Test_Iso3166_CountryCodes_ByName(t *testing.T) {
	for _, expected := range id.Iso3166CountryCodes.Values {
		actual, ok := id.Iso3166CountryCodes.ByName(expected.Name)
		assert.Equal(t, true, ok)
		assert.Equal(t, expected, actual, expected.Name, actual.Name)
	}
}

func Test_Iso3166__CountryCodes_ByAlpha2(t *testing.T) {
	for _, expected := range id.Iso3166CountryCodes.Values {
		if expected.Alpha2 != "" {
			actual, ok := id.Iso3166CountryCodes.ByAlpha2(expected.Alpha2)
			assert.Equal(t, true, ok)
			assert.Equal(t, expected, actual)
		}
	}
}

func Test_Iso3166__CountryCodes_ByAlpha3(t *testing.T) {
	for _, expected := range id.Iso3166CountryCodes.Values {
		if expected.Alpha3 != "" {
			actual, ok := id.Iso3166CountryCodes.ByAlpha3(expected.Alpha3)
			assert.Equal(t, true, ok)
			assert.Equal(t, expected, actual)
		}
	}
}

func Test_Iso3166__CountryCodes_ByNumeric(t *testing.T) {
	for _, expected := range id.Iso3166CountryCodes.Values {
		if expected.Numeric != 0 {
			actual, ok := id.Iso3166CountryCodes.ByNumeric(expected.Numeric)
			assert.Equal(t, true, ok)
			assert.Equal(t, expected, actual)
		}
	}
}

func Test_Iso3166__CountryCodes_ByNumericString(t *testing.T) {
	for _, expected := range id.Iso3166CountryCodes.Values {
		if expected.Numeric != 0 {
			actual, ok := id.Iso3166CountryCodes.ByNumericString(fmt.Sprintf("%v", expected.Numeric))
			assert.Equal(t, true, ok)
			assert.Equal(t, expected, actual)
		}
	}
}

func Test_Iso3166__CountryCodes_Lookup(t *testing.T) {
	for _, expected := range id.Iso3166CountryCodes.Values {
		actual, ok := id.Iso3166CountryCodes.Lookup(expected.Name)
		assert.Equal(t, true, ok)
		assert.Equal(t, expected, actual)

		if expected.Alpha2 != "" {
			actual, ok = id.Iso3166CountryCodes.Lookup(expected.Alpha2)
			assert.Equal(t, true, ok)
			assert.Equal(t, expected, actual)
		}

		if expected.Alpha3 != "" {
			actual, ok = id.Iso3166CountryCodes.Lookup(expected.Alpha3)
			assert.Equal(t, true, ok)
			assert.Equal(t, expected, actual)
		}

		if expected.Numeric != 0 {
			actual, ok := id.Iso3166CountryCodes.Lookup(fmt.Sprintf("%v", expected.Numeric))
			assert.Equal(t, true, ok)
			assert.Equal(t, expected, actual)
		}
	}
}
