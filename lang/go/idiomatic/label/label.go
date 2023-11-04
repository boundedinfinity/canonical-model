package label

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type ValueLabel interface {
	Type() string
}

type NameOnlyLabel struct {
	Id   id.Id     `json:"id,omitempty"`
	Name LabelName `json:"name,omitempty"`
}

func (t NameOnlyLabel) Type() string {
	return "name-only"
}

var _ ValueLabel = &NameOnlyLabel{}

// valueLabel

type valueLabel[T any] struct {
	Id    id.Id     `json:"id,omitempty"`
	Name  LabelName `json:"name,omitempty"`
	Value T         `json:"value,omitempty"`
	typ   string
}

func (t valueLabel[T]) Type() string {
	return t.typ
}

// StringLabel

type StringLabel = valueLabel[string]

func NewString() *StringLabel {
	return &StringLabel{typ: "string"}
}

var _ ValueLabel = NewString()

// IntegerLabel

type IntegerLabel = valueLabel[int]

func NewInteger() *IntegerLabel {
	return &IntegerLabel{typ: "int"}
}

var _ ValueLabel = NewInteger()

// FloatLabel

type FloatLabel = valueLabel[float64]

func NewFloat() *FloatLabel {
	return &FloatLabel{typ: "float"}
}

var _ ValueLabel = NewFloat()

// DateLabel

type DateLabel = valueLabel[rfc3339date.Rfc3339Date]

func NewDate() *DateLabel {
	return &DateLabel{typ: "date"}
}

var _ ValueLabel = NewDate()

// DateTimeLabel

type DateTimeLabel = valueLabel[rfc3339date.Rfc3339Date]

func NewDateTime() *DateTimeLabel {
	return &DateTimeLabel{typ: "date-time"}
}

var _ ValueLabel = NewDateTime()
