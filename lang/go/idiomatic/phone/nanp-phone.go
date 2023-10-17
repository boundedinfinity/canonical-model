package phone

import "github.com/google/uuid"

// https://en.wikipedia.org/wiki/North_American_Numbering_Plan
// https://en.wikipedia.org/wiki/National_conventions_for_writing_telephone_numbers

type NanpPhone struct {
	Id          uuid.UUID `json:"id,omitempty"`
	Title       string    `json:"title,omitempty"`
	CountryCode []Digit   `json:"country-code,omitempty"`
	Npa         []Digit   `json:"npa,omitempty"`
	CoCode      []Digit   `json:"co,omitempty"`
	LineNumber  []Digit   `json:"line-number,omitempty"`
	Extention   Extention `json:"extention,omitempty"`
}
