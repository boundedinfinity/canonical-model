package people

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

// ///////////////////////////////////////////////////
// Name Formatter
// ///////////////////////////////////////////////////

func NewNameFormatter(prefix PrefixFormat, name NameFormat, suffix SuffixFormat, override bool) NameFormatter {
	return NameFormatter{
		name:     name,
		override: override,
		suffix:   NewSuffixFormatterWithIndex(suffix, 0, override),
		prefix:   NewPrefixFormatterWithIndex(prefix, 0, override),
	}
}

type NameFormatter struct {
	prefix   PrefixFormatter
	name     NameFormat
	suffix   SuffixFormatter
	override bool
}

func (t *NameFormatter) Prefix(prefix PrefixFormatter) *NameFormatter {
	t.prefix = prefix
	return t
}

func (t *NameFormatter) Suffix(suffix SuffixFormatter) *NameFormatter {
	t.suffix = suffix
	return t
}

func (t *NameFormatter) Override() *NameFormatter {
	t.override = true
	return t
}

func (t NameFormatter) Format(name Name) string {
	var items []string

	for _, prefix := range name.Prefixes {
		items = append(items, t.prefix.Format(prefix))
	}

	nameFormat, _ := slicer.FirstNotZero(name.Format, t.name)

	if t.override {
		nameFormat = t.name
	}

	switch nameFormat {
	case NameFormats.GivenNameFamilyName:
		items = append(items, name.FamilyName(), name.GivenName())
	default:
		items = append(items, name.GivenName(), name.FamilyName())
	}

	for _, suffix := range name.Suffixes {
		items = append(items, t.suffix.Format(suffix))
	}

	s := strings.Join(items, " ")
	s = strings.Join(strings.Fields(s), " ")

	return s
}

// ///////////////////////////////////////////////////
// Prefix Formatter
// ///////////////////////////////////////////////////

func NewPrefixFormatter(formatter PrefixFormat) PrefixFormatter {
	return PrefixFormatter{formatter: formatter}
}

func NewPrefixFormatterWithIndex(formatter PrefixFormat, index int, override bool) PrefixFormatter {
	return PrefixFormatter{formatter: formatter, index: index, override: override}
}

type PrefixFormatter struct {
	formatter PrefixFormat
	index     int
	override  bool
}

func (t *PrefixFormatter) Override() *PrefixFormatter {
	t.override = true
	return t
}

func (t *PrefixFormatter) Index(index int) *PrefixFormatter {
	t.index = index
	return t
}

func (t PrefixFormatter) Format(prefix Prefix) string {
	var s string

	formatter, _ := slicer.FirstNotZero(prefix.Format, t.formatter)

	if t.override {
		formatter = t.formatter
	}

	switch formatter {
	case PrefixFormats.Full:
		s = prefix.Text
	default:
		if len(prefix.Abbreviation)-1 >= t.index {
			s = prefix.Abbreviation[0]
		} else if len(prefix.Abbreviation) > 0 {
			s = prefix.Abbreviation[0]
		} else {
			s = prefix.Text
		}
	}

	return s
}

// ///////////////////////////////////////////////////
// Suffix Formatter
// ///////////////////////////////////////////////////

func NewSuffixFormatter(formatter SuffixFormat) SuffixFormatter {
	return SuffixFormatter{formatter: formatter}
}

func NewSuffixFormatterWithIndex(formatter SuffixFormat, index int, override bool) SuffixFormatter {
	return SuffixFormatter{formatter: formatter, index: index, override: override}
}

type SuffixFormatter struct {
	formatter SuffixFormat
	index     int
	override  bool
}

func (t *SuffixFormatter) Override() *SuffixFormatter {
	t.override = true
	return t
}

func (t *SuffixFormatter) Index(index int) *SuffixFormatter {
	t.index = index
	return t
}

func (t SuffixFormatter) Format(suffix Suffix) string {
	var s string

	formatter, _ := slicer.FirstNotZero(suffix.Format, t.formatter)

	if t.override {
		formatter = t.formatter
	}

	switch formatter {
	case SuffixFormats.Full:
		s = suffix.Text
	default:
		if len(suffix.Abbreviation)-1 >= t.index {
			s = suffix.Abbreviation[0]
		} else if len(suffix.Abbreviation) > 0 {
			s = suffix.Abbreviation[0]
		} else {
			s = suffix.Text
		}
	}

	return s
}
