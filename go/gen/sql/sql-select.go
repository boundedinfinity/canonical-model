package sql

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

type sqlSelect struct {
	Distinct bool
	All      bool
}

func (this *sqlSelect) Generate() (string, error) {
	var words []string

	words = append(words, "SELECT")

	if this.Distinct {
		words = append(words, "DISTINCT")
	} else if this.All {
		words = append(words, "ALL")
	} else if this.Distinct && this.All {
		// TODO error here
	}

	return stringer.Join(" ", words...), nil
}
