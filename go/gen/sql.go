package gen

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type sqlDatabase struct {
	Tables  []sqlTable
	Indexes []sqlIndex
}

func (this *sqlDatabase) Init() error {
	for _, table := range this.Tables {
		if err := table.Init(this); err != nil {
			return err
		}
	}

	return nil
}

type sqlTable struct {
	Database    *sqlDatabase
	Name        string
	Columns     []sqlColumn
	ForeignKeys []foreignKey
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

type foreignKey struct {
	Table    *sqlTable
	Name     string
	Domestic sqlColumn
	Foreign  sqlColumn
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

type sqlIndex struct {
	Name        string
	Columns     []sqlColumn
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
