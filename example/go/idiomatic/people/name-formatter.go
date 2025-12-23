package people

import (
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// ///////////////////////////////////////////////////
// Name Formatter
// ///////////////////////////////////////////////////

//
// {f}, {g:i}. {m:i}.
// {p} {g} {m} {f} {s:f}
//

type affixMode string
type nameMode string

const (
	AffixAbbreviation affixMode = "a"
	AffixFull         affixMode = "f"

	NameFull    nameMode = "a"
	NameInitial nameMode = "i"
)

type nameFormatter struct {
	format           string
	prefixCategories []string
	prefix           affixMode
	given            nameMode
	middleInitial    nameMode
	lastNameAll      nameMode
	suffix           affixMode
	suffixCategories []string
}

func (this nameFormatter) Format(name Name) string {
	result := this.format

	if len(this.prefixCategories) > 0 && len(name.Prefixes) > 0 {
	check_pefix:
		for _, prefixName := range this.prefixCategories {
			for _, prefix := range name.Prefixes {
				if prefix.Category.Name == prefixName {
					switch this.prefix {
					case AffixAbbreviation:
						if len(prefix.Abbrs) > 0 {
							result = strings.ReplaceAll(result, "{p}", prefix.Abbrs[0])
						}
					default:
						result = strings.ReplaceAll(result, "{p}", prefix.Text)
					}
					break check_pefix
				}
			}
		}
	}

	return result
}

func (this nameFormatter) getAffix(name Name) string {
	var results []string

	for _, prefixName := range this.prefixCategories {
		for _, prefix := range name.Prefixes {
			if prefix.Category.Name == prefixName {
				switch this.prefix {
				case AffixAbbreviation:
					if len(prefix.Abbrs) > 0 {
						results = append(results, prefix.Abbrs[0])
					}
				default:
					results = append(results, prefix.Text)
				}
			}
		}
	}

	return stringer.Join(" ", results...)
}
