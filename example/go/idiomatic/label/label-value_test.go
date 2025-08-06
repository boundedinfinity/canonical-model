package label_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/label"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	assert.Equal(t, "dob", label.Lookup["dob"].Abbreviation)
}
