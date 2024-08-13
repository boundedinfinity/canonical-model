package id_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/stretchr/testify/assert"
)

func Test_Id_Json(t *testing.T) {
	tcs := []struct {
		name     string
		input    id.Id
		expected string
		err      error
	}{
		{
			name:     "case 1",
			input:    id.Id{},
			expected: "null",
			err:      nil,
		},
		{
			name:     "case 3",
			input:    id.Ids.MustParse("E29FE45A-B1B6-4D9D-8084-2012E39F37AE"),
			expected: `"E29FE45A-B1B6-4D9D-8084-2012E39F37AE"`,
			err:      nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := json.Marshal(tc.input)

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, tc.expected, string(actual))
		})
	}

}
