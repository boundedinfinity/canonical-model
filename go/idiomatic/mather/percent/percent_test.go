package percent_test

import (
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/mather/percent"
	"github.com/stretchr/testify/assert"
)

func Test_Percent(t *testing.T) {
	p := percent.FromFloat[int](6.124)
	assert.Equal(t, 6.124, p.Percentage())
	assert.Equal(t, .06124, p.Decimal())
}
