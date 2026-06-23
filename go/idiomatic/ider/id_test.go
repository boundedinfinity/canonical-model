package ider_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/stretchr/testify/assert"
)

func Test_id(t *testing.T) {
	tcs := []struct {
		name     string
		input    ider.Id
		expected ider.Id
	}{
		{
			name:     "is zero",
			input:    ider.Zero,
			expected: ider.Zero,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			assert.Equal(t, tc.expected, tc.input)
		})
	}
}

func Test_id_json(t *testing.T) {
	id := ider.New()
	data, err := json.Marshal(id)
	assert.NoError(t, err)

	var parsedId ider.Id
	err = json.Unmarshal(data, &parsedId)
	assert.NoError(t, err)
	assert.Equal(t, id, parsedId)

	type test struct {
		Id   ider.Id `json:"id"`
		Name string  `json:"name"`
	}

	s1 := test{Id: id, Name: "n"}
	data, err = json.Marshal(s1)
	assert.NoError(t, err)

	var ps1 test
	err = json.Unmarshal(data, &ps1)
	assert.NoError(t, err)
	assert.Equal(t, s1, ps1)

	a1 := []test{{Id: id, Name: "n"}}
	data, err = json.Marshal(a1)
	assert.NoError(t, err)

	var pa1 []test
	err = json.Unmarshal(data, &pa1)
	assert.NoError(t, err)
	assert.Equal(t, a1, pa1)
}
