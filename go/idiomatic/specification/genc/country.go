package genc

type Country struct {
	Name string `json:"name,omitempty,omitzero"`
	Code string `json:"code,omitempty,omitzero"`
}
