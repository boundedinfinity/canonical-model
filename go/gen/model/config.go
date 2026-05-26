package model

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	o "github.com/boundedinfinity/go-commoner/functional/optioner"
)

var (
	ErrKind     = errorer.New("kind error")
	errKindFunc = errorer.Func(ErrKind)
)

type Config struct {
	Package    string   `json:"package"`
	OutputRoot string   `json:"outputRoot"`
	Kinds      []Model  `json:"kinds"`
	Language   Language `json:"language"`
}

type Database struct {
	Name o.Option[string] `json:"name"`
	Type o.Option[string] `json:"type"`
}
