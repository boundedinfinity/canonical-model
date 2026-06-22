package term

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// Format represents the format of the term's string representation. It uses two placeholders:
//
// - "{name}" for the term's name
//
// - "{abbr}" for the term's primary (first) abbreviation.
//
// Any other characters in the format string will be included as-is in the output.
//
// Note: If the format string contains "{abbr}" but the term has no abbreviations, the "{abbr}"
// placeholder will be replaced with an empty string. The formatter will also attempt to remove any
// surrounding parentheses or brackets left from the empty abbreviation.
type Format string

var Formats = formats{
	NameOnly: "{name}",
	NameAbbr: "{name} ({abbr})",
	AbbrOnly: "{abbr}",
	AbbrName: "{abbr} ({name})",
}

type formats struct {
	NameOnly Format
	NameAbbr Format
	AbbrOnly Format
	AbbrName Format
}

// TODO: Add optimizations
// TODO: Allow selecting a custom abbr

var ErrFormat = errorer.New("invalid format")
var errFormatErrorFn = errorer.Func(ErrTerm, ErrFormat)

func formatTerm(format Format, term Term) (string, error) {
	containsName := stringer.ContainsAnyIgnoreCase(format, "{name}")
	containsAbbr := stringer.ContainsAnyIgnoreCase(format, "{abbr}")

	if !(containsName || containsAbbr) {
		errFormatErrorFn("must contain at least one of {name} or {abbr} in the format: %s", format)
	}

	s := string(format)

	if containsName {
		s = stringer.Replace(s, "{name}", term.Name)
	}

	if containsAbbr {
		if len(term.Abbreviations) > 0 {
			s = stringer.Replace(s, "{abbr}", term.Abbreviations[0])
		} else {
			s = stringer.Replace(s, "{abbr}", "")
			s = stringer.Replace(s, "()", "")
			s = stringer.Replace(s, "[]", " ")
		}
	}

	return s, nil
}
