package sql

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// ====================================================================================
// Types
// ====================================================================================

type TableSchema struct {
	Schema       optioner.Option[string]
	Name         string
	IfNotExists  bool
	WithoutRowId bool
	Columns      []*ColumnSchema
	ForeignKeys  []*ForeignKeySchema
	Formatted    bool
	columnMap    map[string]*ColumnSchema
}

func (this TableSchema) qualifiedName() string {
	return tableName(&this)
}

func (this *TableSchema) Generate() (string, error) {
	this.ensureId()

	var errs []error
	var fragments []string

	columns, err := slicer.MapErr(func(column *ColumnSchema) (string, error) {
		return column.Generate()
	}, this.Columns...)

	fragments = append(fragments, columns...)
	errs = append(errs, err)

	primaryKeys, err := PrimaryKey(this.primaryKeyColumns()...).Generate()
	errs = append(errs, err)
	fragments = append(fragments, primaryKeys)

	foreignKeys, err := slicer.MapErr(func(key *ForeignKeySchema) (string, error) {
		return key.Generate()
	})
	errs = append(errs, err)
	fragments = append(fragments, foreignKeys...)

	sep := ";"

	if this.Formatted {
		sep += "\n"
	}

	var sb stringBuiler

	sb.Writef("CREATE TABLE")

	if this.IfNotExists {
		sb.Writef(" IF NOT EXISTS")
	}

	if this.Schema.Defined() {
		sb.Writef(" %s.%s", quote(this.Schema.Get()), quote(this.Name))
	} else {
		sb.Writef(" %s (", quote(this.Name))
	}

	if this.Formatted {
		sb.Writef("\n")
	}

	sb.Writef(stringer.Join(sep, fragments...))
	sb.Writef(" ")

	sb.Writef(")")

	if this.WithoutRowId {
		sb.Writef(" WITHOUT ROWID")
	}

	sb.Writef(";")

	return sb.String(), errors.Join(errs...)
}

func (this TableSchema) ensureId() {
	for _, column := range this.Columns {
		if column.PrimaryKey {
			return
		}
	}

	this.AddColumn(&ColumnSchema{Name: "id", Type: ColumnTypes.TEXT, PrimaryKey: true})
}

func (this TableSchema) primaryKeyColumns() []*ColumnSchema {
	return slicer.Filter(
		func(_ int, column *ColumnSchema) bool { return column.PrimaryKey },
		this.Columns...)
}

func (this TableSchema) indexedColumns() []*ColumnSchema {
	return slicer.Filter(
		func(_ int, column *ColumnSchema) bool { return column.Indexed },
		this.Columns...)
}

func (this *TableSchema) Column(name string) *ColumnSchema {
	var found *ColumnSchema

	for _, column := range this.Columns {
		if column.Name == name {
			found = column
			break
		}
	}

	if found == nil {
		panic(&ErrTableColumnNotFoundDetails{TableName: this.Name, ColumnName: name})
	}

	return found
}

func (this *TableSchema) AddColumn(column *ColumnSchema) *TableSchema {
	column.Table = this
	this.Columns = append(this.Columns, column)

	if this.columnMap == nil {
		this.columnMap = map[string]*ColumnSchema{}
	}

	this.columnMap[column.Name] = column
	return this
}

// ====================================================================================
// Builder
// ====================================================================================

func Table() *tableBuilder {
	return &tableBuilder{
		schema:  new(TableSchema),
		columns: []*ColumnBuilder{},
	}
}

type tableBuilder struct {
	schema  *TableSchema
	columns []*ColumnBuilder
}

func (this *tableBuilder) Build() *TableSchema {
	for _, column := range this.columns {
		this.schema.AddColumn(column.Build())
	}

	return this.schema
}

func (this *tableBuilder) Column(columns ...*ColumnBuilder) *tableBuilder {
	return appendAndReturn(this, &this.columns, columns...)
}

func (this *tableBuilder) Name(name string) *tableBuilder {
	return setAndReturn(this, &this.schema.Name, name)
}

func (this *tableBuilder) IfNotExists(enable bool) *tableBuilder {
	return setAndReturn(this, &this.schema.IfNotExists, enable)
}

func (this *tableBuilder) WithoutRowId(enable bool) *tableBuilder {
	return setAndReturn(this, &this.schema.WithoutRowId, enable)
}

func (this *tableBuilder) Formatted(enable bool) *tableBuilder {
	return setAndReturn(this, &this.schema.Formatted, enable)
}

// ====================================================================================
// Errors
// ====================================================================================

var ErrTableColumnNotFound = errors.New("column not found")

type ErrTableColumnNotFoundDetails struct {
	TableName  string
	ColumnName string
}

func (e ErrTableColumnNotFoundDetails) Error() string {
	return fmt.Sprintf("%s : %s.%s", ErrTableColumnNotFound.Error(), e.TableName, e.ColumnName)
}

func (e ErrTableColumnNotFoundDetails) Unwrap() error {
	return ErrTableColumnNotFound
}
