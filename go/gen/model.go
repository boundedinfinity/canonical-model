package gen

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	o "github.com/boundedinfinity/go-commoner/functional/optioner"
)

var (
	ErrKind     = errorer.New("kind error")
	errKindFunc = errorer.Func(ErrKind)
)

type Config struct {
	Package    string       `json:"package"`
	OutputRoot string       `json:"outputRoot"`
	Kinds      []Kind       `json:"kinds"`
	Language   LanguageKind `json:"language"`
}

type Kind interface {
	GetKind() KindKind
}

var _ Kind = &LanguageKind{}

type LanguageKind struct {
	Name    o.Option[string] `json:"name"`
	Package string           `json:"package"`
	Type    string           `json:"type"`
	Alias   o.Option[string] `json:"alias"`
}

func (_ LanguageKind) GetKind() KindKind {
	return Kinds.Language
}

var _ Kind = &StringKind{}

type StringKind struct {
	Name string        `json:"name"`
	Min  o.Option[int] `json:"min"`
	Max  o.Option[int] `json:"max"`
}

func (_ StringKind) GetKind() KindKind {
	return Kinds.String
}

var _ Kind = &ReferenceKind{}

type ReferenceKind struct {
	Name      o.Option[string] `json:"name"`
	Reference string           `json:"reference"`
}

func (_ ReferenceKind) GetKind() KindKind {
	return Kinds.Reference
}
