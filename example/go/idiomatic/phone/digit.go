package phone

// https://en.wikipedia.org/wiki/North_American_Numbering_Plan#Alphabetic_mnemonic_system
// https://en.wikipedia.org/wiki/E.161
// https://www.itu.int/rec/T-REC-E.161

import (
	"fmt"

	"errors"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type Digit struct {
	Key       string   `json:"key,omitempty,omitzero"`
	Mnemonics []string `json:"mnemonics,omitempty,omitzero"`
}

type SelectedDigit struct {
	Digit
	Mnemonic string `json:"mnemonic,omitempty,omitzero"`
}

var (
	ErrDigitNotValid  = errors.New(`digit must be one of: a number between 0 and 9, a mnemonic letter between A-Z or a-z, or the characters +, #, ','`)
	errDigitNotValidf = func(digit string) error { return fmt.Errorf("%s: %w", digit, ErrDigitNotValid) }
)

func (t Digit) String() string {
	return t.Key
}

var (
	digitMap = map[string]Digit{}
)

func (this digits) Parse(s string) (Digit, error) {
	var digit *Digit
	var err error

	for _, d := range this.All() {
		nlc := slicer.Map(func(_ int, s string) string { return stringer.Lowercase(s) }, d.Mnemonics...)

		if d.Key == s || stringer.ContainsAny(s, d.Mnemonics...) || stringer.ContainsAny(s, nlc...) {
			digit = &d
		}
	}

	if digit == nil {
		err = errDigitNotValidf(s)
	}

	return *digit, err
}
