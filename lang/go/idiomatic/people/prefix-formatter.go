package people

import (
	"fmt"

	"github.com/boundedinfinity/schema/idiomatic/id"
)

type PrefixFormatter struct {
	Id     id.Id
	Name   string
	Desc   string
	Format func(Prefix) (string, error)
}

func (t PrefixFormatter) String() string {
	return fmt.Sprintf("%s Prefix Formatter", t.Name)
}

// /////////////////////////////////////////////////////////////////
//  Initialization
// /////////////////////////////////////////////////////////////////

var PrefixFormatters = prefixFormatters{}

func init() {
	PrefixFormatters.Abbr = PrefixFormatter{
		Id:   id.Ids.MustParse("25483819-F00D-4E36-AB52-8DE49F74A319"),
		Name: "Abbreviation Formatter",
		Desc: `If the Prefix contains an abbreviation that will be returned.
        If there isn't an abbreviation, then the full text is used.
        `,
		Format: func(p Prefix) (string, error) {
			if p.Abbr != "" {
				return p.Abbr, nil
			}
			return p.Text, nil
		},
	}

	PrefixFormatters.Text = PrefixFormatter{
		Id:   id.Ids.MustParse("A746758C-3E4E-4E3C-A651-84F710391667"),
		Name: "Text Formatter",
		Desc: `Used the text field of the Prefix.
        `,
		Format: func(p Prefix) (string, error) {
			return p.Text, nil
		},
	}

	PrefixFormatters.list = []PrefixFormatter{
		PrefixFormatters.Abbr,
		PrefixFormatters.Text,
	}
}

// /////////////////////////////////////////////////////////////////
//  Companion
// /////////////////////////////////////////////////////////////////

type prefixFormatters struct {
	list []PrefixFormatter
	Abbr PrefixFormatter
	Text PrefixFormatter
}

func (t *prefixFormatters) List() []PrefixFormatter {
	return t.list
}

func (t *prefixFormatters) Get(term string) (PrefixFormatter, bool) {
	var found PrefixFormatter

	for _, formatter := range t.List() {
		if term == formatter.Id.String() {
			found = formatter
			break
		}

		if term == formatter.Name {
			found = formatter
			break
		}
	}

	return found, false
}
