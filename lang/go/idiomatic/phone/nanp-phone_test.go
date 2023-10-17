package phone_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/phone"
	"github.com/stretchr/testify/assert"
)

func Test_NanpPhone_Format_Default(t *testing.T) {
	actual := phone.NanpPhone{
		Npa: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		CoCode: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		LineNumber: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
	}

	assert.Equal(t, `(555) 555-5555`, actual.String())
}

func Test_NanpPhone_Format_Dashes(t *testing.T) {
	actual := phone.NanpPhone{
		SeparatorFormat: phone.NapaPhoneSeparatorFormats.DashSeparated,
		Npa: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		CoCode: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		LineNumber: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
	}

	assert.Equal(t, `555-555-5555`, actual.String())
}

func Test_NanpPhone_Format_Dots(t *testing.T) {
	actual := phone.NanpPhone{
		SeparatorFormat: phone.NapaPhoneSeparatorFormats.DotSeparated,
		Npa: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		CoCode: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		LineNumber: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
	}

	assert.Equal(t, `555.555.5555`, actual.String())
}

func Test_NanpPhone_Format_None(t *testing.T) {
	actual := phone.NanpPhone{
		SeparatorFormat: phone.NapaPhoneSeparatorFormats.None,
		Npa: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		CoCode: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		LineNumber: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
	}

	assert.Equal(t, `5555555555`, actual.String())
}
