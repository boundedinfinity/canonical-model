package check_algo

import (
	"errors"
	"fmt"
	"strconv"
)

type LuhnCheckOptions struct {
	NoCheckDigit bool
}

var (
	ErrLuhnValidationFailure  = errors.New("Luhn validation failure")
	errLuhnValidationFailuref = func(format string, a ...any) error {
		message := fmt.Sprintf(format, a...)
		return fmt.Errorf("%w : %s", ErrLuhnValidationFailure, message)
	}
)

type LuhnCheck string

func (this LuhnCheck) atoi(position int, char string) (int, error) {
	digit, err := strconv.Atoi(string(char))

	if err != nil {
		return digit, errLuhnValidationFailuref(
			"the character %s (located at position %d) is not a valid number: %s",
			char, position, err.Error(),
		)
	}

	return digit, nil
}

func (this LuhnCheck) CalculateCheckDigit() (int, error) {
	var result int
	payloadLength := len(this)

	if payloadLength <= 0 {
		return result, errLuhnValidationFailuref("number length is zero")
	}

	payload := this[:payloadLength-1]
	payloadLength -= 1

	if payloadLength <= 0 {
		return result, errLuhnValidationFailuref("payload length is zero")
	}

	digits := []int{}

	for i, char := range payload {
		digit, err := this.atoi(i, string(char))

		if err != nil {
			return digit, err
		}

		digits = append(digits, digit)
	}

	for i := payloadLength - 1; i > 0; i -= 2 {
		digits[i] <<= 1 // double the value

		if digits[i] > 9 {
			digits[i] -= 9
		}
	}

	var sum int

	for i := 0; i < payloadLength; i += 1 {
		sum += digits[i]
	}

	result = (10 - sum%10) % 10
	return result, nil
}

func (this LuhnCheck) GivenCheckDigit() (int, error) {
	position := len(this) - 1

	if position <= 0 {
		return 0, errLuhnValidationFailuref("number length is zero")
	}

	return this.atoi(position, string(this[position]))
}

func (this LuhnCheck) Validate() error {
	calculatedCheckDigit, err := this.CalculateCheckDigit()
	if err != nil {
		return err
	}

	givenCheckDigit, err := this.GivenCheckDigit()
	if err != nil {
		return err
	}

	if calculatedCheckDigit != givenCheckDigit {
		return errLuhnValidationFailuref(
			"the given check digit %d doesn't match the calculated check digit %d",
			givenCheckDigit, calculatedCheckDigit,
		)
	}

	return nil
}
