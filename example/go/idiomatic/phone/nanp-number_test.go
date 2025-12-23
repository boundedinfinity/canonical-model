package phone_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/phone"
	"github.com/stretchr/testify/assert"
)

func Test_NanpNumber_Format_Default(t *testing.T) {
	actual := phone.NanpNumber{
		Npa: []phone.Digit{
			{Key: 5},
			{Key: 5},
			{Key: 5},
		},
		Nxx: []phone.Digit{
			{Key: 5},
			{Key: 5},
			{Key: 5},
		},
		LineNumber: []phone.Digit{
			{Key: 5},
			{Key: 5},
			{Key: 5},
			{Key: 5},
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
		SeparatorFormat: phone.NanpPhoneSeparatorFormats.DashSeparated,
		Npa: []phone.Digit{
			{Key: 5},
			{Key: 5},
			{Key: 5},
		},
		Nxx: []phone.Digit{
			{Key: 5},
			{Key: 5},
			{Key: 5},
		},
		LineNumber: []phone.Digit{
			{Key: 5},
			{Key: 5},
			{Key: 5},
			{Key: 5},
		},
	}

	assert.Equal(t, `555-555-5555`, actual.String())
}

func Test_NanpNumber_Format_Dots(t *testing.T) {
	actual := phone.NanpNumber{
		SeparatorFormat: phone.NanpPhoneSeparatorFormats.DotSeparated,
		Npa: []phone.Digit{
			phone.NewDigitMust(5),
			phone.NewDigitMust(5),
			phone.NewDigitMust(5),
		},
		Nxx: []phone.Digit{
			phone.NewDigitMust(5),
			phone.NewDigitMust(5),
			phone.NewDigitMust(5),
		},
		LineNumber: []phone.Digit{
			phone.NewDigitMust(5),
			phone.NewDigitMust(5),
			phone.NewDigitMust(5),
			phone.NewDigitMust(5),
		},
	}

	assert.Equal(t, `555.555.5555`, actual.String())
}

func Test_NanpNumber_Format_None(t *testing.T) {
	actual := phone.NanpNumber{
		SeparatorFormat: phone.NanpPhoneSeparatorFormats.None,
		Npa: []phone.Digit{
			{Key: 5},
			{Key: 5},
			{Key: 5},
		},
		Nxx: []phone.Digit{
			{Key: 5},
			{Key: 5},
			{Key: 5},
		},
		LineNumber: []phone.Digit{
			{Key: 5},
			{Key: 5},
			{Key: 5},
			{Key: 5},
		},
	}

	assert.Equal(t, `5555555555`, actual.String())
}
