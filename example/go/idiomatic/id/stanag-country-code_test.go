package id_test

// import (
// 	"testing"

// 	"github.com/boundedinfinity/schema/idiomatic/id"
// 	"github.com/stretchr/testify/assert"
// )

// func Test_Stanag_CountryCodes_ByName(t *testing.T) {
// 	for _, expected := range id.StanagCodes.Values {
// 		actual, ok := id.StanagCodes.ByName(expected.Name)
// 		assert.Equal(t, true, ok)
// 		assert.Equal(t, expected, actual)
// 	}
// }

// func Test_Stanag_CountryCodes_ByCode(t *testing.T) {
// 	for _, expected := range id.StanagCodes.Values {
// 		if expected.Code != "" {
// 			actual, ok := id.StanagCodes.ByCode(expected.Code)
// 			assert.Equal(t, true, ok)
// 			assert.Equal(t, expected, actual)
// 		}
// 	}
// }

// func Test_Stanag_CountryCodes_Lookup(t *testing.T) {
// 	for _, expected := range id.StanagCodes.Values {
// 		actual, ok := id.StanagCodes.Lookup(expected.Name)
// 		assert.Equal(t, true, ok)
// 		assert.Equal(t, expected, actual)

// 		if expected.Code != "" {
// 			actual, ok := id.StanagCodes.Lookup(expected.Code)
// 			assert.Equal(t, true, ok)
// 			assert.Equal(t, expected, actual)
// 		}
// 	}
// }
