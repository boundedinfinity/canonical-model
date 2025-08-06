package id_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/stretchr/testify/assert"
)

func Test_Genc_CountryCodes_ByName(t *testing.T) {
	for _, expected := range id.GenCodes.Values {
		actual, ok := id.GenCodes.ByName(expected.Name)
		assert.Equal(t, true, ok)
		assert.Equal(t, expected, actual, expected.Name, actual.Name)
	}
}

func Test_Genc_CountryCodes_ByCode(t *testing.T) {
	for _, expected := range id.GenCodes.Values {
		if expected.Code != "" {
			actual, ok := id.GenCodes.ByCode(expected.Code)
			assert.Equal(t, true, ok)
			assert.Equal(t, expected, actual)
		}
	}
}

func Test_Genc_CountryCodes_Lookup(t *testing.T) {
	for _, expected := range id.GenCodes.Values {
		actual, ok := id.GenCodes.Lookup(expected.Name)
		assert.Equal(t, true, ok)
		assert.Equal(t, expected, actual)

		if expected.Code != "" {
			actual, ok := id.GenCodes.Lookup(expected.Code)
			assert.Equal(t, true, ok)
			assert.Equal(t, expected, actual)
		}
	}
}
