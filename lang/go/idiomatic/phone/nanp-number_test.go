package phone_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/phone"
	"github.com/stretchr/testify/assert"
)

func Test_NanpNumber_Format_Default(t *testing.T) {
	actual := phone.NanpNumber{
		Npa: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		Nxx: []phone.Digit{
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
		Extention: phone.Extention{
			Items: []phone.ExtentionItem{
				phone.AtCommands.Pause2Seconds,
				phone.AtCommands.Pause2Seconds,
				phone.NewDigitMust(1),
			},
		},
	}

	assert.Equal(t, `(555) 555-5555 ext ,,1`, actual.String())
}

func Test_NanpNumber_Format_Dashes(t *testing.T) {
	actual := phone.NanpNumber{
		SeparatorFormat: phone.NapaPhoneSeparatorFormats.DashSeparated,
		Npa: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		Nxx: []phone.Digit{
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

func Test_NanpNumber_Format_Dots(t *testing.T) {
	actual := phone.NanpNumber{
		SeparatorFormat: phone.NapaPhoneSeparatorFormats.DotSeparated,
		Npa: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		Nxx: []phone.Digit{
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

func Test_NanpNumber_Format_None(t *testing.T) {
	actual := phone.NanpNumber{
		SeparatorFormat: phone.NapaPhoneSeparatorFormats.None,
		Npa: []phone.Digit{
			{Number: 5},
			{Number: 5},
			{Number: 5},
		},
		Nxx: []phone.Digit{
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
