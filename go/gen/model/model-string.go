package model

import o "github.com/boundedinfinity/go-commoner/functional/optioner"

var _ Model = &StringModel{}

type StringModel struct {
	Name string        `json:"name"`
	Min  o.Option[int] `json:"min"`
	Max  o.Option[int] `json:"max"`
}

func (_ StringModel) GetKind() Kind {
	return Kinds.String
}
