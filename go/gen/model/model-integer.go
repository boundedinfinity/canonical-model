package model

import o "github.com/boundedinfinity/go-commoner/functional/optioner"

var _ Model = &IntegerModel{}

type IntegerModel struct {
	Name       string          `json:"name"`
	Min        o.Option[int]   `json:"min"`
	Max        o.Option[int]   `json:"max"`
	MultipleOf o.Option[int]   `json:"multiple-of"`
	Positive   o.Option[bool]  `json:"positive"`
	Negative   o.Option[bool]  `json:"negative"`
	OneOf      o.Option[[]int] `json:"one-of"`
}

func (_ IntegerModel) GetKind() Kind {
	return Kinds.Integer
}
