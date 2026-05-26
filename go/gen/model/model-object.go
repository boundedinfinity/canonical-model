package model

var _ Model = &ObjectModel{}

type ObjectModel struct {
	Name       string     `json:"name"`
	Properties []Property `json:"properties"`
}

func (_ ObjectModel) GetKind() Kind {
	return Kinds.Object
}

type Property struct {
	Kind     Model
	Optional bool `json:"optional"`
}
