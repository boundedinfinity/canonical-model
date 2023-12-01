package messenger_test

import (
	"log"
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/messenger"
	"github.com/boundedinfinity/schema/idiomatic/people"
	"github.com/stretchr/testify/assert"
)

func TestMarshal(t *testing.T) {
	expected := people.Person{
		Name: people.Name{
			Prefixes:    []people.Prefix{people.Prefixes[0]},
			GivenNames:  []string{"James"},
			FamilyNames: []string{"Bond"},
			Suffixes:    []people.Suffix{people.Suffixes[0]},
		},
	}

	marshaler := messenger.NewMarshaller()

	bs, err := marshaler.Marshal(expected)
	assert.Nil(t, err)

	log.Println(string(bs))

	iface, err := marshaler.Unmarshal(bs)
	assert.Nil(t, err)

	actual, ok := iface.(people.Person)

	assert.True(t, ok)
	assert.Equal(t, expected, actual)
}
