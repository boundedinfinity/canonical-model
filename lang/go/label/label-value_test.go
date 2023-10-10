package label_test

import (
	"testing"

	"github.com/boundedinfinity/optioner"
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/label"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	dobLN := label.LabelName{
		Id:          optioner.Some(uuid.MustParse("B12E7628-35B0-4AE5-B2D9-0F34F4828229")),
		Name:        "dob",
		Description: optioner.Some("Date of Birth"),
	}

	label := label.Label[rfc3339date.Rfc3339Date]{
		Id:   optioner.Some(uuid.MustParse("0A431E6D-CA13-4E89-8114-02EF92808E72")),
		Name: dobLN,
	}

	assert.Equal(t, "dob", label.Name.Name)
}
