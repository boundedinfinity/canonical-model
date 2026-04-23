package sql

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

// https://www.sqlitetutorial.net/sqlite-index/

type IndexSchema struct {
	Column *ColumnSchema
	Preifx optioner.Option[string]
}

func (this IndexSchema) name() string {
	return quote(
		fmt.Sprintf("%s_%s_%s",
			this.Preifx.OrElse("idx"),
			this.Column.Table.Name,
			this.Column.Name),
	)
}

func (this IndexSchema) Generate() (string, error) {
	if this.Column == nil {
		return "", ErrIndexNoColumns
	}

	var sb stringBuiler

	sb.Writef("CREATE")

	if this.Column.Unique {
		sb.Writef(" UNIQUE")
	}

	sb.Writef(" INDEX %s ON %s (%s);",
		this.name(),
		this.Column.Table.qualifiedName(),
		this.Column.unqualifiedName(),
	)

	return sb.String(), nil
}

var ErrIndexNoColumns = errors.New("index statement has no column")

func Index() *indexBuilder {
	return &indexBuilder{
		schema: new(IndexSchema),
	}
}

type indexBuilder struct {
	schema *IndexSchema
}

func (this *indexBuilder) Column(column *ColumnSchema) *indexBuilder {
	return setAndReturn(this, &this.schema.Column, column)
}

func (this *indexBuilder) Prefix(prefix string) *indexBuilder {
	return setOptAndReturn(this, &this.schema.Preifx, prefix)
}

func (this *indexBuilder) Build() *IndexSchema {
	return this.schema
}
