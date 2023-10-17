package country_code_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/country_code"
	"github.com/stretchr/testify/assert"
)

func Test_NpaCodeLookup(t *testing.T) {
	code, ok := country_code.NpaCodeLookup("720")
	assert.Equal(t, true, ok)
	assert.Equal(t, "CO", code.Alpha2)
	assert.Equal(t, 720, code.Numeric)
}
