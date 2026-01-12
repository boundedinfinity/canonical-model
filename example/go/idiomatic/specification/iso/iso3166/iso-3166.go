package iso3166

// https://en.wikipedia.org/wiki/ISO_3166-1#Codes
// ISO 3166-1 Name, Alpha-2 code, Alpha-3 code, Numeric code, Link to ISO 3166-2, Independent

type Iso3166_1_Alpha2 string
type Iso3166_1_Alpha3 string
type Iso3166_1_Numeric string

type Iso3166_1 struct {
	Source       string      `json:"source,omitempty,omitzero"`
	Name         string      `json:"name,omitempty,omitzero"`
	Alpha2       string      `json:"alpha-2,omitempty,omitzero"`
	Alpha3       string      `json:"alpha-3,omitempty,omitzero"`
	Numeric      string      `json:"numeric,omitempty,omitzero"`
	Independent  string      `json:"independent,omitempty,omitzero"`
	Reference    string      `json:"reference,omitempty,omitzero"`
	Subdivisions []Iso3166_2 `json:"subdivisions,omitempty,omitzero"`
}

type Iso3166_2_Code string

type Iso3166_2 struct {
	Source    string `json:"source,omitempty,omitzero"`
	Name      string `json:"name,omitempty,omitzero"`
	Code      string `json:"code,omitempty,omitzero"`
	Category  string `json:"cagetory,omitempty,omitzero"`
	Reference string `json:"reference,omitempty,omitzero"`
}
