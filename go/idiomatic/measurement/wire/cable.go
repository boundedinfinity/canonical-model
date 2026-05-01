package wire

// https://www.systoncable.com/nm-cable-vs-bx-cable/
// https://aerosusa.com/wires-vs-cables/

type NonMetallicCable struct {
	Name           string `json:"name"`
	Wire           Wire   `json:"wire"`
	ConductorCount int    `json:"conductor-count"`
}

type ArmoredCable struct {
	Name           string `json:"name"`
	Wire           Wire   `json:"wire"`
	ConductorCount int    `json:"conductor-count"`
}

type FlexibleMetalConduitCable struct {
	Name string `json:"name"`
	Wire Wire   `json:"wire"`
}
