package model

var _ Model = &ArrayModel{}

type ArrayModel struct {
	Name string `json:"name"`
	Kind Kind   `json:"kind"`
}

func (_ ArrayModel) GetKind() Kind {
	return Kinds.Array
}
