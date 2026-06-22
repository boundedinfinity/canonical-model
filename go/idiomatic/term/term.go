package term

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/go-commoner/errorer"
)

// Term represents a term item, including its name, abbreviations, and description.
//
// If the abbreviations field is not empty and the length is greater than zero, the first
// abbreviation (index 0) will be considered the primary abbreviation and will be used during
// formmatting the string epresentation of the term.
//
// Otherwise, the name of the term will be used as the stringrepresentation.
type Term struct {
	Id            ider.Id  `json:"id"`
	Name          string   `json:"name"`
	Abbreviations []string `json:"abbreviations"`
	Description   string   `json:"description"`
}

func (this Term) String() string {
	s, err := this.Format(Formats.NameAbbr)

	if err != nil {
		panic(err)
	}

	return s
}

// Format returns the string representation of the term based on the provided format.
func (this Term) Format(format Format) (string, error) {
	return formatTerm(format, this)
}

var ErrTerm = errorer.New("term error")
