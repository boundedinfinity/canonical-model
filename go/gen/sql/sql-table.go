package sql

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

type sqlTable struct {
	Database    *sqlDatabase
	Name        string
	Columns     []*sqlColumn
	ForeignKeys []*foreignKey
	IfNotExists bool
	Temporary   bool
}

func (this *sqlTable) Init(database *sqlDatabase) error {
	this.Database = database
	for _, column := range this.Columns {
		if err := column.Init(this); err != nil {
			return err
		}
	}
	return nil
}

func (this *sqlTable) AddColumn(column *sqlColumn) error {
	// TODO check for duplicate and valid column name
	// TODO check for valid column type
	column.Table = this
	this.Columns = append(this.Columns, column)
	return nil
}

func (this *sqlTable) GetColumn(name string) (*sqlColumn, bool) {
	for _, column := range this.Columns {
		if column.Name == name {
			return column, true
		}
	}

	return nil, false
}

func (this *sqlTable) Generate() (string, error) {
	// https://www.sqlite.org/lang_createtable.html
	var words []string

	words = append(words, "CREATE")

	if this.Temporary {
		words = append(words, "TEMPORARY")
	}

	words = append(words, "TABLE")

	if this.IfNotExists {
		words = append(words, "IF", "NOT", "EXISTS")
	}

	words = append(words, this.Name, "(")

	params := make([]string, len(this.Columns)+len(this.ForeignKeys))

	for i := range this.Columns {
		if str, err := this.Columns[i].Generate(); err != nil {
			return "", err
		} else {
			params[i] = str
		}
	}

	for i := range this.ForeignKeys {
		if str, err := this.ForeignKeys[i].Generate(); err != nil {
			return "", err
		} else {
			params[len(this.Columns)+i] = str
		}
	}

	words = append(words, stringer.Join(", ", params...))
	words = append(words, ");")
	return stringer.Join(" ", words...), nil
}
