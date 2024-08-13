package phone

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

// https://en.wikipedia.org/wiki/North_American_Numbering_Plan
// https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers

type NanpNumber struct {
	Id              id.Id                    `json:"id,omitempty"`
	Title           string                   `json:"title,omitempty"`
	CountryCode     []Digit                  `json:"country-code,omitempty"`
	Npa             []Digit                  `json:"npa,omitempty"`
	Nxx             []Digit                  `json:"nxx,omitempty"`
	LineNumber      []Digit                  `json:"line-number,omitempty"`
	Extention       Extention                `json:"extention,omitempty"`
	SeparatorFormat NanpPhoneSeparatorFormat `json:"separator-format,omitempty"`
	Format          NanpPhoneFormat          `json:"format,omitempty"`
}

var _ id.TypeNamer = &NanpNumber{}

func (t NanpNumber) TypeName() string {
	return id.TypeNamers.Dotted(NanpNumber{})
}

func (t NanpNumber) String() string {
	var sb stringer.Builder[string]
	format, _ := slicer.NotZero(t.Format, NanpPhoneFormats.Common)
	separatorFormat, _ := slicer.NotZero(t.SeparatorFormat, NanpPhoneSeparatorFormats.ParenthesesAndDashes)

	switch format {
	default:
	case NanpPhoneFormats.Full:
		if len(t.CountryCode) > 0 {
			sb.WriteRune('+')

			for _, code := range t.CountryCode {
				sb.WriteString(fmt.Sprint(code.Number))
			}

			switch separatorFormat {
			case NanpPhoneSeparatorFormats.DashSeparated:
				sb.WriteByte('-')
			case NanpPhoneSeparatorFormats.DotSeparated:
				sb.WriteByte('.')
			case NanpPhoneSeparatorFormats.None:
			default:
				sb.WriteByte(' ')
			}
		}
	}

	if separatorFormat == NanpPhoneSeparatorFormats.ParenthesesAndDashes {
		sb.WriteRune('(')
	}

	for _, code := range t.Npa {
		sb.WriteString(fmt.Sprint(code.Number))
	}

	switch separatorFormat {
	case NanpPhoneSeparatorFormats.DashSeparated:
		sb.WriteByte('-')
	case NanpPhoneSeparatorFormats.DotSeparated:
		sb.WriteByte('.')
	case NanpPhoneSeparatorFormats.None:
	default:
		sb.WriteString(") ")
	}

	for _, code := range t.Nxx {
		sb.WriteString(fmt.Sprint(code.Number))
	}

	switch separatorFormat {
	case NanpPhoneSeparatorFormats.DotSeparated:
		sb.WriteByte('.')
	case NanpPhoneSeparatorFormats.None:
	default:
		sb.WriteRune('-')
	}

	for _, code := range t.LineNumber {
		sb.WriteString(fmt.Sprint(code.Number))
	}

	if t.Extention.Has() {
		sb.WriteString(" " + t.Extention.String())
	}

	return sb.String()
}

func FromString(s string) (NanpNumber, error) {
	phone := NanpNumber{}

	return phone, nil
}
