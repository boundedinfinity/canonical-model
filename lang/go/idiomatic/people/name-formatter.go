package people

// ///////////////////////////////////////////////////
// Name Formatter
// ///////////////////////////////////////////////////

type nameFormatter struct {
	prefixSocial       bool
	prefixFormal       bool
	prefixAcademic     bool
	prefixProfessional bool
	prefixReligious    bool
	firstNameAll       bool
	middleInitial      bool
	lastNameAll        bool
	suffixSocial       bool
	suffixProfessional bool
}

func NameFormatter() *nameFormatter {
	return &nameFormatter{
		suffixSocial: true,
	}
}

func (t *nameFormatter) PrefixAll() *nameFormatter {
	t.prefixSocial = true
	return t
}

func (t *nameFormatter) PrefixSocial() *nameFormatter {
	t.prefixSocial = true
	return t
}

func (t *nameFormatter) PrefixFormal() *nameFormatter {
	t.prefixSocial = false
	t.prefixFormal = true
	return t
}

// func (t NameFormatter) Format(name Name) string {
// 	var items []string

// 	nameFormat, _ := slicer.FirstNotZero(name.Format, t.name)

// 	if t.override {
// 		nameFormat = t.name
// 	}

// 	switch nameFormat {
// 	case NameFormats.First, NameFormats.Given:
// 		if item, ok := slicer.Head(name.Firsts...); ok {
// 			items = append(items, item)
// 		}
// 	case NameFormats.FirstAll, NameFormats.GivenAll:
// 		items = append(items, name.Firsts...)
// 	case NameFormats.Last, NameFormats.Surname:
// 		if item, ok := slicer.Last(name.Lasts...); ok {
// 			items = append(items, item)
// 		}
// 	case NameFormats.LastAll, NameFormats.SurnameAll:
// 		items = append(items, name.Lasts...)
// 	default:
// 		items = append(items, name.GivenName(), name.FamilyName())
// 	}

// 	for _, suffix := range name.Suffixes {
// 		items = append(items, t.suffix.Format(suffix))
// 	}

// 	s := strings.Join(items, " ")
// 	s = strings.Join(strings.Fields(s), " ")

// 	return s
// }

// // ///////////////////////////////////////////////////
// // Prefix Formatter
// // ///////////////////////////////////////////////////

// func NewPrefixFormatter(formatter PrefixFormat) PrefixFormatter {
// 	return PrefixFormatter{formatter: formatter}
// }

// func NewPrefixFormatterWithIndex(formatter PrefixFormat, index int, override bool) PrefixFormatter {
// 	return PrefixFormatter{formatter: formatter, index: index, override: override}
// }

// type PrefixFormatter struct {
// 	formatter PrefixFormat
// 	index     int
// 	override  bool
// }

// func (t *PrefixFormatter) Override() *PrefixFormatter {
// 	t.override = true
// 	return t
// }

// func (t *PrefixFormatter) Index(index int) *PrefixFormatter {
// 	t.index = index
// 	return t
// }

// func (t PrefixFormatter) Format(prefix Prefix) string {
// 	var s string

// 	formatter, _ := slicer.FirstNotZero(prefix.Format, t.formatter)

// 	if t.override {
// 		formatter = t.formatter
// 	}

// 	switch formatter {
// 	case PrefixFormats.Full:
// 		s = prefix.Text
// 	default:
// 		if len(prefix.Abbreviation)-1 >= t.index {
// 			s = prefix.Abbreviation[0]
// 		} else if len(prefix.Abbreviation) > 0 {
// 			s = prefix.Abbreviation[0]
// 		} else {
// 			s = prefix.Text
// 		}
// 	}

// 	return s
// }

// // ///////////////////////////////////////////////////
// // Suffix Formatter
// // ///////////////////////////////////////////////////

// func NewSuffixFormatter(formatter SuffixFormat) SuffixFormatter {
// 	return SuffixFormatter{formatter: formatter}
// }

// func NewSuffixFormatterWithIndex(formatter SuffixFormat, index int, override bool) SuffixFormatter {
// 	return SuffixFormatter{formatter: formatter, index: index, override: override}
// }

// type SuffixFormatter struct {
// 	formatter SuffixFormat
// 	index     int
// 	override  bool
// }

// func (t *SuffixFormatter) Override() *SuffixFormatter {
// 	t.override = true
// 	return t
// }

// func (t *SuffixFormatter) Index(index int) *SuffixFormatter {
// 	t.index = index
// 	return t
// }

// func (t SuffixFormatter) Format(suffix Suffix) string {
// 	var s string

// 	formatter, _ := slicer.FirstNotZero(suffix.Format, t.formatter)

// 	if t.override {
// 		formatter = t.formatter
// 	}

// 	switch formatter {
// 	case SuffixFormats.Full:
// 		s = suffix.Text
// 	default:
// 		if len(suffix.Abbreviation)-1 >= t.index {
// 			s = suffix.Abbreviation[0]
// 		} else if len(suffix.Abbreviation) > 0 {
// 			s = suffix.Abbreviation[0]
// 		} else {
// 			s = suffix.Text
// 		}
// 	}

// 	return s
// }
