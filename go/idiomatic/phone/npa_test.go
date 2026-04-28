package phone_test

import (
	"testing"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/phone"
	"github.com/stretchr/testify/assert"
)

func Test_NpaCodeLookup(t *testing.T) {
	code, ok := phone.NpaCodeLookup("720")
	assert.Equal(t, true, ok)
	assert.Equal(t, "CO", code.Alpha2)
	assert.Equal(t, 720, code.Numeric)
}
