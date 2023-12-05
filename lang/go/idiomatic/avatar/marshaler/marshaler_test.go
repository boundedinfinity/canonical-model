package marshaler_test

import (
	"fmt"
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/avatar/marshaler"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model"
	"github.com/boundedinfinity/schema/idiomatic/avatar/model/people"
	cpeople "github.com/boundedinfinity/schema/idiomatic/people"
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
				Header: model.EventHeader{
					Type: people.PersonEvent{}.Value.TypeName(),
				},
				Value: cpeople.Person{
					Name: cpeople.Name{
						GivenNames:  []string{"James"},
						FamilyNames: []string{"Bond"},
					},
				},
			},
			expected: `{"header":{"type":"people.Person","received-at":null},"value":{"id":null,"name":{"id":null,"given-names":["James"],"family-names":["Bond"]},"measurements":{},"dates":{"birth-date":null,"death-date":null}}}`,
		},
	}

	em := marshaler.NewMarshaller()
	// em.Formatted(true)

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

func Test_Marshaler_Unmarshal(t *testing.T) {
	// tcs := []struct {
	// 	name     string
	// 	expected any
	// 	input    string
	// 	err      error
	// }{
	// 	{
	// 		name: "person.PersonEvent",
	// 		expected: people.PersonEvent{
	// 			Header: model.EventHeader{
	// 				Type: people.PersonEvent{}.Value.TypeName(),
	// 			},
	// 			Value: cpeople.Person{
	// 				Name: cpeople.Name{
	// 					GivenNames:  []string{"James"},
	// 					FamilyNames: []string{"Bond"},
	// 				},
	// 			},
	// 		},
	// 		input: `{"header":{"type":"people.Person","received-at":null},"value":{"id":null,"name":{"id":null,"given-names":["James"],"family-names":["Bond"]},"measurements":{},"dates":{"birth-date":null,"death-date":null}}}`,
	// 	},
	// }

	// for _, tc := range tcs {
	// 	t.Run(tc.name, func(tt *testing.T) {
	// 		em := marshaler.NewMarshaller()

	// 		var actual people.PersonEvent

	// 		em.RegisterHandlerFn(func(discrim string, val any) {
	// 			switch concreate := val.(type) {
	// 			case people.PersonEvent:
	// 				actual = concreate
	// 			}
	// 		})

	// 		err := em.Unmarshal([]byte(tc.input))
	// 		fmt.Println(string(tc.input))

	// 		if tc.err != nil {
	// 			assert.Equal(t, tc.err, err)
	// 		} else {
	// 			assert.Nil(t, err)
	// 		}

	// 		assert.Equal(t, tc.expected, actual)
	// 	})
	// }
}

// km.RegisterHandlerFn(func(key string, raw any) {
// 	switch val := raw.(type) {
// 	case people.PersonEvent:

// 	default:
// 	}
// })
