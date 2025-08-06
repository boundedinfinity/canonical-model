package phone

import (
	"strconv"

	"errors"
)

type Digit struct {
	Number int  `json:"number,omitempty"`
	Letter rune `json:"letter,omitempty"`
}

var (
	ErrDigitNotBetween0And9 = errors.New("digit must be between 0 and 9")
)

func (t Digit) String() string {
	return strconv.FormatInt(int64(t.Number), 10)
}

func NewDigit(digit int) (Digit, error) {
	if digit < 0 || digit > 9 {
		return Digit{}, ErrDigitNotBetween0And9
	}

	return Digit{
		Number: digit,
	}, nil
}

func NewDigitMust(digit int) Digit {
	d, err := NewDigit(digit)

	if err != nil {
		panic(err)
	}

	return d
}
