package location_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/location"
	"github.com/stretchr/testify/assert"
)

func Test_NpaCodeLookup(t *testing.T) {
	code, ok := location.NpaCodeLookup("720")
	assert.Equal(t, true, ok)
	assert.Equal(t, "CO", code.Alpha2)
	assert.Equal(t, 720, code.Numeric)
}
