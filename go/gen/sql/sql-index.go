package sql

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

type sqlIndex struct {
	Name        string
	Columns     []*sqlColumn
	IfNotExists bool
	Unique      bool
}

func (this *sqlIndex) Generate() (string, error) {
	// https://sqlite.org/lang_createindex.html
	var words []string

	words = append(words, "CREATE")

	if this.Unique {
		words = append(words, "UNIQUE")
	}

	words = append(words, "INDEX")

	if this.IfNotExists {
		words = append(words, "IF", "NOT", "EXISTS")
	}

	words = append(words, this.Name, "ON", this.Columns[0].Table.Name, "(")

	for _, column := range this.Columns {
		words = append(words, column.Name+",")
	}

	words = append(words, ");")

	return stringer.Join(" ", words...), nil
}
