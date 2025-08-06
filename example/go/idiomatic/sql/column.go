package sql

import (
	"errors"
	"fmt"
)

// ====================================================================================
// Types
// ====================================================================================

type ColumnType string

type columnTypes struct {
	NULL    ColumnType
	INTEGER ColumnType
	REAL    ColumnType
	TEXT    ColumnType
	BLOB    ColumnType
}

var ColumnTypes = columnTypes{
	NULL:    "NULL",
	INTEGER: "INTEGER",
	REAL:    "REAL",
	TEXT:    "TEXT",
	BLOB:    "BLOB",
}

type ColumnSchema struct {
	Table         *TableSchema
	Name          string
	Type          ColumnType
	NotNull       bool
	PrimaryKey    bool
	Unique        bool
	Check         bool
	AutoIncrement bool
	Indexed       bool
}

func (this ColumnSchema) unqualifiedName() string {
	return unqualifiedColumnName(&this)
}

func (this ColumnSchema) qualifiedName() string {
	return qualifiedColumnName(&this)
}

func (this ColumnSchema) Generate() (string, error) {
	if this.Type == "" {
		return "", &ErrColumnSchema{
			Column: &this,
			Cause:  ErrColumnSchemaNoType,
		}
	}

	var sb stringBuiler

	sb.Writef("%s %s", quote(this.Name), this.Type)

	if this.PrimaryKey {
		sb.Writef(" PRIMARY KEY")
	}

	if this.NotNull {
		sb.Writef(" NOT NULL")
	}

	if this.Unique {
		sb.Writef(" UNIQUE")
	}

	if this.Check {
		sb.Writef(" CHECK")
	}

	if this.AutoIncrement {
		sb.Writef(" AUTOINCREMENT")
	}

	return sb.String(), nil
}

// ====================================================================================
// Err
// ====================================================================================

var (
	ErrColumnSchemaNoType = errors.New("column has not type")
)

type ErrColumnSchema struct {
	Column *ColumnSchema
	Cause  error
}

func (e ErrColumnSchema) Error() string {
	return fmt.Sprintf("%s : %s", e.Cause.Error(), e.Column.qualifiedName())
}

func (e ErrColumnSchema) Unwrap() error {
	return e.Cause
}

// ====================================================================================
// Builder
// ====================================================================================

func Column() *ColumnBuilder {
	return &ColumnBuilder{
		schema: new(ColumnSchema),
	}
}

type ColumnBuilder struct {
	schema *ColumnSchema
}

func (this *ColumnBuilder) Table(table *TableSchema) *ColumnBuilder {
	return setAndReturn(this, &this.schema.Table, table)
}

func (this *ColumnBuilder) Name(name string) *ColumnBuilder {
	return setAndReturn(this, &this.schema.Name, name)
}

func (this *ColumnBuilder) Type(typ ColumnType) *ColumnBuilder {
	return setAndReturn(this, &this.schema.Type, typ)
}

func (this *ColumnBuilder) PrimaryKey(enable bool) *ColumnBuilder {
	return setAndReturn(this, &this.schema.PrimaryKey, enable)
}

func (this *ColumnBuilder) Unique(enable bool) *ColumnBuilder {
	return setAndReturn(this, &this.schema.Unique, enable)
}

func (this *ColumnBuilder) AutoIncrement(enable bool) *ColumnBuilder {
	return setAndReturn(this, &this.schema.AutoIncrement, enable)
}

func (this *ColumnBuilder) Indexed(enable bool) *ColumnBuilder {
	return setAndReturn(this, &this.schema.Indexed, enable)
}

func (this *ColumnBuilder) Build() *ColumnSchema {
	return this.schema
}
