package iso3166

type Subdivision struct {
	Source        string `json:"source,omitempty,omitzero"`
	Name          string `json:"name,omitempty,omitzero"`
	Code          string `json:"code,omitempty,omitzero"`
	Category      string `json:"category,omitempty,omitzero"`
	Parent        string `json:"parent,omitempty,omitzero"`
	Reference     string `json:"reference,omitempty,omitzero"`
	AlternateCode string `json:"alternate-code,omitempty,omitzero"`
	PreviousCode  string `json:"previous-code,omitempty,omitzero"`
}
