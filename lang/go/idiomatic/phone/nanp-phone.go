package phone

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/google/uuid"
)

// https://en.wikipedia.org/wiki/North_American_Numbering_Plan
// https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers

type NanpPhone struct {
	Id              uuid.UUID                `json:"id,omitempty"`
	Title           string                   `json:"title,omitempty"`
	CountryCode     []Digit                  `json:"country-code,omitempty"`
	Npa             []Digit                  `json:"npa,omitempty"`
	CoCode          []Digit                  `json:"co,omitempty"`
	LineNumber      []Digit                  `json:"line-number,omitempty"`
	Extention       Extention                `json:"extention,omitempty"`
	SeparatorFormat NapaPhoneSeparatorFormat `json:"separator-format,omitempty"`
	Format          NapaPhoneFormat          `json:"format,omitempty"`
}

func (t NanpPhone) String() string {
	var sb stringer.Builder[string]
	format, _ := slicer.FirstNotZero(t.Format, NapaPhoneFormats.Common)
	separatorFormat, _ := slicer.FirstNotZero(t.SeparatorFormat, NapaPhoneSeparatorFormats.ParenthesesAndDashes)

	switch format {
	default:
	case NapaPhoneFormats.Full:
		if len(t.CountryCode) > 0 {
			sb.WriteRune('+')

			for _, code := range t.CountryCode {
				sb.WriteString(fmt.Sprint(code.Number))
			}

			switch separatorFormat {
			case NapaPhoneSeparatorFormats.DashSeparated:
				sb.WriteByte('-')
			case NapaPhoneSeparatorFormats.DotSeparated:
				sb.WriteByte('.')
			case NapaPhoneSeparatorFormats.None:
			default:
				sb.WriteByte(' ')
			}
		}
	}

	if separatorFormat == NapaPhoneSeparatorFormats.ParenthesesAndDashes {
		sb.WriteRune('(')
	}

	for _, code := range t.Npa {
		sb.WriteString(fmt.Sprint(code.Number))
	}

	switch separatorFormat {
	case NapaPhoneSeparatorFormats.DashSeparated:
		sb.WriteByte('-')
	case NapaPhoneSeparatorFormats.DotSeparated:
		sb.WriteByte('.')
	case NapaPhoneSeparatorFormats.None:
	default:
		sb.WriteString(") ")
	}

	for _, code := range t.CoCode {
		sb.WriteString(fmt.Sprint(code.Number))
	}

	switch separatorFormat {
	case NapaPhoneSeparatorFormats.DotSeparated:
		sb.WriteByte('.')
	case NapaPhoneSeparatorFormats.None:
	default:
		sb.WriteRune('-')
	}

	for _, code := range t.LineNumber {
		sb.WriteString(fmt.Sprint(code.Number))
	}

	return sb.String()
}

func FromString(s string) (NanpPhone, error) {
	phone := NanpPhone{}

	return phone, nil
}
