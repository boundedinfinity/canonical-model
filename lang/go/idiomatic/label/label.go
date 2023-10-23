package label

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
	"golang.org/x/exp/constraints"
)

type Label struct {
	Id   id.Id     `json:"id,omitempty"`
	Name LabelName `json:"name,omitempty"`
}

type StringLabel[T ~string] valueLabel[T]
type IntegerLabel[T constraints.Integer] valueLabel[T]
type FloatLabel[T constraints.Float] valueLabel[T]
type DateLabel valueLabel[rfc3339date.Rfc3339Date]
type DateTimeLabel valueLabel[rfc3339date.Rfc3339DateTime]

type valueLabel[V constraints.Ordered | rfc3339date.Rfc3339Date | rfc3339date.Rfc3339DateTime] struct {
	Id    id.Id     `json:"id,omitempty"`
	Name  LabelName `json:"name,omitempty"`
	Value V         `json:"value,omitempty"`
}
