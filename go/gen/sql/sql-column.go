package sql

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

type sqlColumn struct {
	Table      *sqlTable
	Name       string
	Type       string
	NotNull    bool
	PrimaryKey bool
}

func (this *sqlColumn) Init(table *sqlTable) error {
	this.Table = table
	return nil
}

func (this *sqlColumn) Generate() (string, error) {
	// // https://www.sqlite.org/lang_createtable.html
	var words []string

	words = append(words, this.Name, this.Type)

	if this.PrimaryKey {
		words = append(words, "PRIMARY", "KEY")
	}

	if this.NotNull && !this.PrimaryKey {
		words = append(words, "NOT", "NULL")
	}

	return stringer.Join(" ", words...), nil
}
