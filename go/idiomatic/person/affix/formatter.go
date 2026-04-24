package affix

import (
	"regexp"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/util/slicer"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/util/stringer"
)

var (
	// https://regex101.com/
	affixRegex = regexp.MustCompile(`{((?:prefix|suffix)(?:\.abbr)?)(?::(.*))?}`)
)

func PrefixFormatter() Formatter {
	return Formatter{
		Pattern: "{prefix.abbr:common}",
		Sep:     " ",
	}
}

func SuffixFormatter() Formatter {
	return Formatter{
		Pattern: "{suffix.abbr:common}",
		Sep:     " ",
	}
}

type Formatter struct {
	Pattern string
	Sep     string
}

func (this Formatter) Format(prefixes, suffixes []Affix) string {
	var kind string
	var abbr bool
	var categoryNames []string
	result := affixRegex.FindAllString(this.Pattern, -1)

	switch len(result) {
	case 2:
		kind, abbr = processName(result[1])
	case 3:
		kind, abbr = processName(result[1])
		categoryNames = processCategory(result[2])
	}

	var working []Affix

	switch kind {
	case "prefix":
		if len(categoryNames) > 0 {
			working = slicer.FilterFunc(prefixes, categoryMatches(categoryNames))
		}

		working = prefixes
	case "suffix":
		if len(categoryNames) > 0 {
			working = slicer.FilterFunc(suffixes, categoryMatches(categoryNames))
		}

		working = suffixes
	}

	name := stringer.JoinFunc(working, this.Sep, getNameOrAbbr(abbr))

	return name
}

func processName(s string) (string, bool) {
	var name string
	var abbr bool

	parts := stringer.Split(s, ".")
	switch len(parts) {
	case 1:
		name = parts[0]
	case 2:
		name = parts[0]
		abbr = true
	}

	return name, abbr
}

func processCategory(s string) []string {
	parts := stringer.Split(s, ",")
	parts = stringer.SliceToLower(parts)
	return parts
}

func categoryMatches(names []string) func(Affix) bool {
	return func(affix Affix) bool {
		return slicer.Contains(names, stringer.ToLower(affix.Category.Name))
	}
}

func getNameOrAbbr(abbr bool) func(Affix) string {
	return func(affix Affix) string {
		if abbr && len(affix.Abbreviations) > 0 {
			return affix.Abbreviations[0]
		}
		return affix.Name
	}
}
