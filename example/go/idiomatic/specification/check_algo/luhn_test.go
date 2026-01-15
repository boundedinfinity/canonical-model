package check_algo_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/specification/check_algo"
	"github.com/stretchr/testify/assert"
)

func Test_Luhn_valid(t *testing.T) {
	luhn := check_algo.LuhnCheck("17893729974")

	given, errGiven := luhn.GivenCheckDigit()
	assert.Equal(t, given, 4)
	assert.NoError(t, errGiven)

	calculated, errCalculated := luhn.CalculateCheckDigit()
	assert.Equal(t, calculated, 4)
	assert.NoError(t, errCalculated)

	errValidate := luhn.Validate()
	assert.NoError(t, errValidate)
}

func Test_Luhn_empty(t *testing.T) {
	luhn := check_algo.LuhnCheck("")

	given, errGiven := luhn.GivenCheckDigit()
	assert.Equal(t, given, 0)
	assert.ErrorIs(t, errGiven, check_algo.ErrLuhnValidationFailure)

	calculated, errCalculated := luhn.CalculateCheckDigit()
	assert.Equal(t, calculated, 0)
	assert.ErrorIs(t, errCalculated, check_algo.ErrLuhnValidationFailure)

	errValidate := luhn.Validate()
	assert.ErrorIs(t, errValidate, check_algo.ErrLuhnValidationFailure)

}

func Test_Luhn_fail_check(t *testing.T) {
	luhn := check_algo.LuhnCheck("27893729974")

	given, errGiven := luhn.GivenCheckDigit()
	assert.Equal(t, given, 4)
	assert.NoError(t, errGiven)

	calculated, errCalculated := luhn.CalculateCheckDigit()
	assert.Equal(t, calculated, 3)
	assert.NoError(t, errCalculated)

	errValidate := luhn.Validate()
	assert.ErrorIs(t, errValidate, check_algo.ErrLuhnValidationFailure)
}

func Test_Luhn_not_all_numbers(t *testing.T) {
	luhn := check_algo.LuhnCheck("A7893729974")

	given, errGiven := luhn.GivenCheckDigit()
	assert.Equal(t, given, 4)
	assert.NoError(t, errGiven)

	calculated, errCalculated := luhn.CalculateCheckDigit()
	assert.Equal(t, calculated, 3)
	assert.NoError(t, errCalculated)

	errValidate := luhn.Validate()
	assert.ErrorIs(t, errValidate, check_algo.ErrLuhnValidationFailure)
}
