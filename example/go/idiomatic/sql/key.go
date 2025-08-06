package sql

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// ====================================================================================
// Primary Key
// ====================================================================================

func PrimaryKey(columns ...*ColumnSchema) *primaryKeyFragment {
	return &primaryKeyFragment{columns: columns}
}

type primaryKeyFragment struct {
	columns []*ColumnSchema
}

func (this primaryKeyFragment) Generate() (string, error) {
	column2str := func(column *ColumnSchema) (string, error) {
		return column.Generate()
	}

	str2quote := func(s string) (string, error) {
		return fmt.Sprintf(`"%s"`, s), nil
	}

	columns, err := Chain(column2str).Then(str2quote).Run(this.columns...)

	if err != nil {
		return "", err
	}

	var sb stringBuiler

	sb.Writef(`PRIMARY KEY (%s)`, stringer.Join(", ", columns...))

	return sb.String(), nil
}

// ====================================================================================
// Foreign Key
// ====================================================================================

func ForeignKey() *ForeignKeySchema {
	return &ForeignKeySchema{}
}

type ForeignKeySchema struct {
	Foreign  *ColumnSchema
	Domestic *ColumnSchema
	OnUpdate optioner.Option[ForeignKeyAction]
	OnDelete optioner.Option[ForeignKeyAction]
}

func (this ForeignKeySchema) Generate() (string, error) {
	var sb stringBuiler

	sb.Writef(
		"FOREIGN KEY (%s) REFERENCES %s(%s)",
		quote(this.Domestic.Name),
		quote(this.Foreign.Table.Name),
		quote(this.Foreign.Name),
	)

	if this.OnDelete.Defined() {
		sb.Writef("ON DELETE %s", this.OnDelete.Get())
	}

	if this.OnUpdate.Defined() {
		sb.Writef("ON UPDATE %s", this.OnUpdate.Get())
	}

	return sb.String(), nil
}

type ForeignKeyAction string

type foreignKeyActions struct {
	SET_NULL    ForeignKeyAction
	SET_DEFAULT ForeignKeyAction
	RESTRICT    ForeignKeyAction
	NO_ACTION   ForeignKeyAction
	CASCADE     ForeignKeyAction
}

var ForeignKeyActions = foreignKeyActions{
	SET_NULL:    "SET NULL",
	SET_DEFAULT: "SET DEFAULT",
	RESTRICT:    "RESTRICT",
	NO_ACTION:   "NO ACTION",
	CASCADE:     "CASCADE",
}
