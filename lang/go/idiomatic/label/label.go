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

type StringLabel struct {
	Id    id.Id     `json:"id,omitempty"`
	Name  LabelName `json:"name,omitempty"`
	Value string    `json:"value,omitempty"`
}

func (t StringLabel) Type() string {
	return "string"
}

var _ ValueLabel = &StringLabel{}

type IntegerLabel struct {
	Id    id.Id     `json:"id,omitempty"`
	Name  LabelName `json:"name,omitempty"`
	Value int       `json:"value,omitempty"`
}

func (t IntegerLabel) Type() string {
	return "int"
}

var _ ValueLabel = &IntegerLabel{}

type FloatLabel struct {
	Id    id.Id     `json:"id,omitempty"`
	Name  LabelName `json:"name,omitempty"`
	Value float64   `json:"value,omitempty"`
}

func (t FloatLabel) Type() string {
	return "float"
}

var _ ValueLabel = &FloatLabel{}

type DateLabel struct {
	Id    id.Id                   `json:"id,omitempty"`
	Name  LabelName               `json:"name,omitempty"`
	Value rfc3339date.Rfc3339Date `json:"value,omitempty"`
}

func (t DateLabel) Type() string {
	return "date"
}

var _ ValueLabel = &DateLabel{}

type DateTimeLabel struct {
	Id    id.Id                       `json:"id,omitempty"`
	Name  LabelName                   `json:"name,omitempty"`
	Value rfc3339date.Rfc3339DateTime `json:"value,omitempty"`
}

func (t DateTimeLabel) Type() string {
	return "datetime"
}

var _ ValueLabel = &DateTimeLabel{}
