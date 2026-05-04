package validatorer_test

import (
	"testing"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/validatorer"
	"github.com/stretchr/testify/assert"
)

func Test_error(t *testing.T) {
	testCases := []struct {
		name       string
		validator  validatorer.Validator
		inputName  string
		input      string
		targetErrs []error
		errMsg     string
	}{
		{
			name:       "min 5",
			validator:  validatorer.String().Min(5),
			inputName:  "s",
			input:      "12345",
			targetErrs: []error{},
			errMsg:     "",
		},
		{
			name:       "min 5",
			validator:  validatorer.String().Min(5),
			inputName:  "s",
			input:      "1234",
			targetErrs: []error{validatorer.ErrValidate, validatorer.ErrStringValidate},
			errMsg:     "less than 5",
		},
		{
			name:       "min 5",
			validator:  validatorer.String().Max(5),
			inputName:  "s",
			input:      "12345",
			targetErrs: []error{},
			errMsg:     "",
		},
		{
			name:       "min 5",
			validator:  validatorer.String().Max(5),
			inputName:  "s",
			input:      "123456",
			targetErrs: []error{validatorer.ErrValidate, validatorer.ErrStringValidate},
			errMsg:     "greater than 5",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(tt *testing.T) {
			err := tc.validator.Validate(tc.inputName, tc.input)
			handleErrors(t, err, tc.targetErrs, tc.errMsg)
		})
	}
}

func handleErrors(t *testing.T, err error, targetErrs []error, errMsg string) {
	if len(targetErrs) > 0 && err == nil {
		t.Errorf("err is null but should contain an error with %d targets", len(targetErrs))
	} else {
		for _, targetErr := range targetErrs {
			assert.ErrorIs(t, err, targetErr)
		}
	}

	if errMsg != "" {
		if err == nil {
			t.Errorf("err is null but should contain '%s'", errMsg)
		} else {
			assert.Contains(t, err.Error(), errMsg)
		}
	}
}
