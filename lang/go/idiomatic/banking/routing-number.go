package banking

type FederalReserveRoutingSymbol struct {
	Number int `json:"number,omitempty"`
}

type AbaInstitutionIdentifier struct {
	Number int `json:"number,omitempty"`
}

type ProcessingCenterLocation struct {
	Number int `json:"number,omitempty"`
}

type CheckDigit struct {
	Number int `json:"number,omitempty"`
}

type FractionFormNumber struct {
	Frrs     FederalReserveRoutingSymbol `json:"frrs,omitempty"`
	AbaId    AbaInstitutionIdentifier    `json:"aba-id,omitempty"`
	Location ProcessingCenterLocation    `json:"location,omitempty"`
}

type MagneticInkCharacterRecognitionNumber struct {
	Frrs       FederalReserveRoutingSymbol `json:"frrs,omitempty"`
	AbaId      AbaInstitutionIdentifier    `json:"aba-id,omitempty"`
	CheckDigit CheckDigit                  `json:"check-digit,omitempty"`
}

// 102001017
// XXXXYYYYC
