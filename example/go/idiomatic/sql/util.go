package sql

import (
	"fmt"
	"strings"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

func quote(s string) string {
	return fmt.Sprintf(`"%s"`, s)
}

func quotes(s []string) []string {
	return slicer.Map(func(s string) string { return quote(s) }, s...)
}

func setAndReturn[B any, V any](builder B, field *V, value V) B {
	*field = value
	return builder
}

func setOptAndReturn[B any, V any](builder B, field *optioner.Option[V], value V) B {
	*field = optioner.Some(value)
	return builder
}

func appendAndReturn[B any, V any](builder B, field *[]V, values ...V) B {
	*field = append(*field, values...)
	return builder
}

func tableName(table *TableSchema) string {
	return quote(table.Name)
}

func unqualifiedColumnName(column *ColumnSchema) string {
	return quote(column.Name)
}

func qualifiedColumnName(column *ColumnSchema) string {
	return quote(column.Table.Name) + "." + quote(column.Name)
}

// func getTableNames(tables []*TableSchema) []string {
// 	return slicer.Map(func(table *TableSchema) string { return table.Name }, tables...)
// }

// func columnUnqualifiedNames(columns []*ColumnSchema) []string {
// 	return slicer.Map(func(column *ColumnSchema) string { return column.Name }, columns...)
// }

// func columnQualifiedNames(columns []*ColumnSchema) []string {
// 	return slicer.Map(
// 		func(column *ColumnSchema) string { return column.qualifiedName() },
// 		columns...)
// }

// ====================================================================================
// stringBuilder
// ====================================================================================

type stringBuiler struct {
	sb strings.Builder
}

func (this *stringBuiler) Writef(format string, a ...any) {
	this.sb.WriteString(fmt.Sprintf(format, a...))
}

func (this stringBuiler) String() string {
	return this.sb.String()
}

// ====================================================================================
// Chainer
// ====================================================================================

func Chain[T, U any](fn func(T) (U, error)) *chainer[T, U] {
	return &chainer[T, U]{
		initial: fn,
	}
}

type chainer[T, U any] struct {
	initial func(T) (U, error)
	nexts   []func(U) (U, error)
}

func (this *chainer[T, U]) Then(fn func(U) (U, error)) *chainer[T, U] {
	this.nexts = append(this.nexts, fn)
	return this
}

func (this *chainer[T, U]) Run(items ...T) ([]U, error) {
	var results []U
	var err error

	results, err = slicer.MapErr(this.initial, items...)

	if err != nil {
		return results, err
	}

	for _, next := range this.nexts {
		results, err = slicer.MapErr(next, results...)

		if err != nil {
			break
		}
	}

	return results, err
}
