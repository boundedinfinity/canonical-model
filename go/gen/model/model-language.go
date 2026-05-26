package model

import o "github.com/boundedinfinity/go-commoner/functional/optioner"

var _ Model = &LanguageModel{}

type LanguageModel struct {
	Name     o.Option[string]   `json:"name"`
	Package  string             `json:"package"`
	Type     string             `json:"type"`
	Alias    o.Option[string]   `json:"alias"`
	Database o.Option[Database] `json:"database"`
}

func (_ LanguageModel) GetKind() Kind {
	return Kinds.Language
}
