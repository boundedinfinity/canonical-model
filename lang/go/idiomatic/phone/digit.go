package phone

type Digit struct {
	Number int  `json:"number,omitempty"`
	Letter rune `json:"letter,omitempty"`
}
