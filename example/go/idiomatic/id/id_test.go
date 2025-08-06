package id_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/test_utils"
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

func Test_Id_Db(t *testing.T) {
	db, dbcleanup := test_utils.GetDb("Test_Id_Db")
	defer dbcleanup()

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
			name:     "case 2",
			input:    id.Ids.MustParse("E29FE45A-B1B6-4D9D-8084-2012E39F37AE"),
			expected: `"E29FE45A-B1B6-4D9D-8084-2012E39F37AE"`,
			err:      nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			err := test_utils.TableCreate(db, tc.name, map[string]string{"col": "TEXT"})
			assert.ErrorIs(tt, err, nil)

			err = test_utils.TableInsert(db, tc.name, map[string]any{"col": tc.input})
			assert.ErrorIs(tt, err, nil)
		})
	}
}
