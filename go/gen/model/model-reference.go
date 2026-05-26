package model

import o "github.com/boundedinfinity/go-commoner/functional/optioner"

var _ Model = &ReferenceModel{}

type ReferenceModel struct {
	Name      o.Option[string] `json:"name"`
	Reference string           `json:"reference"`
}

func (_ ReferenceModel) GetKind() Kind {
	return Kinds.Reference
}
