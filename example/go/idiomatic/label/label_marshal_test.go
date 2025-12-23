package label_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/label"
	"github.com/stretchr/testify/assert"
)

func Test_Label_Marshal(t *testing.T) {
	nameOnlyLabel := label.Label{
		Name: label.LabelNames.MustFind("dob"),
	}

	bs, err := label.LabelMarshalJSON(&nameOnlyLabel)

	assert.Nil(t, err)
	assert.Equal(t, `Mr. James Bond Jr.`, string(bs))
}
