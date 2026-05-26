package model

import (
	o "github.com/boundedinfinity/go-commoner/functional/optioner"
)

var _ Model = &FloatModel{}

type FloatModel struct {
	Name     string            `json:"name"`
	Min      o.Option[float64] `json:"min"`
	Max      o.Option[float64] `json:"max"`
	Positive o.Option[bool]    `json:"positive"`
	Negative o.Option[bool]    `json:"negative"`
}

func (_ FloatModel) GetKind() Kind {
	return Kinds.Float
}
