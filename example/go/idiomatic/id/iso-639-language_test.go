package id_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/stretchr/testify/assert"
)

func Test_Language_Lookup(t *testing.T) {
	actual, ok := id.Iso639LanguageLookup("en")

	assert.True(t, ok)
	assert.Equal(t, "English", actual.Name)
	assert.Equal(t, "en", actual.V1)
	assert.Equal(t, "eng", actual.V2b)
	assert.Equal(t, "eng", actual.V2t)
	assert.Equal(t, "eng", actual.V3)
}
