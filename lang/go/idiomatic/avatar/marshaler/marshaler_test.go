package marshaler_test

import (
	"fmt"
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/avatar/marshaler"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model/people"
	"github.com/stretchr/testify/assert"
)

func Test_Marshaler_Marshal(t *testing.T) {
	tcs := []struct {
		name     string
		input    any
		expected string
		err      error
	}{
		{
			name: "person.PersonEvent",
			input: people.PersonEvent{
				Header: model.EventHeader{},
			},
			expected: "{}",
		},
	}

	em := marshaler.NewMarshaller()
	em.Formatted(true)

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := em.Marshal(tc.input)
			actual := string(bs)
			fmt.Println(actual)

			if tc.err != nil {
				assert.Equal(t, tc.err, err)
			} else {
				assert.Nil(t, err)
			}

			assert.JSONEq(t, tc.expected, actual)
		})
	}

}

// km.RegisterHandlerFn(func(key string, raw any) {
// 	switch val := raw.(type) {
// 	case people.PersonEvent:

// 	default:
// 	}
// })
