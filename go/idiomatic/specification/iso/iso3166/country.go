package iso3166

// https://en.wikipedia.org/wiki/ISO_3166-1#Codes
// ISO 3166-1 Name, Alpha-2 code, Alpha-3 code, Numeric code, Link to ISO 3166-2, Independent

type Country struct {
	Source       string        `json:"source,omitempty,omitzero"`
	Name         string        `json:"name,omitempty,omitzero"`
	Alpha2       string        `json:"alpha-2,omitempty,omitzero"`
	Alpha3       string        `json:"alpha-3,omitempty,omitzero"`
	Numeric      string        `json:"numeric,omitempty,omitzero"`
	Independent  string        `json:"independent,omitempty,omitzero"`
	Reference    string        `json:"reference,omitempty,omitzero"`
	Subdivisions []Subdivision `json:"subdivisions,omitempty,omitzero"`
}
