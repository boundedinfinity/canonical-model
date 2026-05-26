package model

var _ Model = &BooleanModel{}

type BooleanModel struct {
	Name string `json:"name"`
	Kind Kind   `json:"kind"`
}

func (_ BooleanModel) GetKind() Kind {
	return Kinds.Boolean
}
