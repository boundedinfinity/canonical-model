package sql

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

type foreignKey struct {
	Table    *sqlTable
	Name     string
	Domestic *sqlColumn
	Foreign  *sqlColumn
	Unique   bool
}

func (this *foreignKey) Init(table *sqlTable) error {
	this.Table = table
	return nil
}

func (this *foreignKey) Generate() (string, error) {
	// https://www.sqlite.org/lang_createtable.html
	var words []string

	words = append(words, "FOREIGN", "KEY", this.Domestic.Name, "REFERENCES", this.Foreign.Table.Name, "(", this.Foreign.Name, ")")

	if this.Unique {
		words = append(words, "UNIQUE")
	}

	return stringer.Join(" ", words...), nil
}
