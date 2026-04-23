package mather

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func calcLeftOfDecimal(s string) int {
	s = stringer.TrimLeft(s, " 0")
	size := len(s)
	return size
}

func calcRightOfDecimal(s string) int {
	s = stringer.TrimRight(s, " 0")
	size := len(s)
	return size
}

// PartSizes returns the number of digits before and after the decimal in a floating point number.
//
// If there are no non-0 digits to the left of the decimal point, this function returns 0 for the integer part.
//
// If there are no digits to the right of the decimal point (the input is an integer value), this
// function returns 0.  This function also ignores all 0 digits after the right most non-0 digit.
//
// For example given the following number:
//
//	003.141500
//
// The result of this function is 1, 4.
func PartSizes[F idiomatic.Float](f F) (int, int) {
	var integer int
	var fraction int
	s := fmt.Sprintf("%f", f)
	c := strings.Split(s, ".")

	switch len(c) {
	case 1:
		integer = calcLeftOfDecimal(c[0])
	case 2:
		integer = calcLeftOfDecimal(c[0])
		fraction = calcRightOfDecimal(c[1])
	default:
		panic("can't process: " + s)
	}

	return integer, fraction
}

// IntegerPartSize returns the number of digits before the decimal in a floating point number.
//
// If there are no digits to the left of the decimal point or the numbers to the left of the
// decimal is 0, this function returns 0.
//
// For example given the following number:
//
//	003.141500
//
// The result of this function is 1.
func IntegerPartSize[F idiomatic.Float](f F) int {
	var size int
	s := fmt.Sprintf("%f", f)
	c := strings.Split(s, ".")

	switch len(c) {
	case 1:
		size = calcLeftOfDecimal(c[0])
	default:
		panic("can't process: " + s)
	}

	return size
}

// FractionalPartSize returns the number of digits after the decimal in a floating point number.
//
// If there are no digits to the right of the decimal point (the input is an integer value), this
// function returns 0.
//
// For example given the following number:
//
//	003.141500
//
// The result of this function is 4.
//
// This will ignore non-significat trailing zeros
func FractionalPartSize[F idiomatic.Float](f F) int {
	var size int
	s := fmt.Sprintf("%f", f)
	c := strings.Split(s, ".")

	switch len(c) {
	case 1:
		// no decimal
	case 2:
		size = calcRightOfDecimal(c[1])
	default:
		panic("can't process: " + s)
	}

	return size
}

func FractionalPartPlace[F idiomatic.Float](f F) int {
	size := FractionalPartSize(f)
	place, err := strconv.Atoi("1" + strings.Repeat("0", size))

	if err != nil {
		panic(err)
	}

	return place
}

// Signbit reports whether n is negative or negative zero.
func Signbit[N idiomatic.Number](n N) bool {
	return math.Signbit(float64(n))
}

// IsNegative reports whether n is negative or negative zero.
func IsNegative[I idiomatic.Number](n I) bool {
	return Signbit(n)
}

func Itoa[I idiomatic.Integer](i I) string {
	return strconv.FormatInt(int64(i), 10)
}

func OrdinalIndicator[I idiomatic.Integer](n I) string {
	// https://en.wikipedia.org/wiki/Ordinal_indicator

	var indicator string
	s := Itoa(n)
	c := s[len(s)-1:]

	switch c {
	case "0":
		switch len(s) {
		case 2:
			indicator = "ieth"
		default:
			indicator = "th"
		}
	case "1":
		indicator = "st"
	case "2":
		indicator = "nd"
	case "3":
		indicator = "rd"
	default:
		indicator = "th"
	}

	return indicator
}

func Ordinalize[I idiomatic.Integer](n I) string {
	// https://en.wikipedia.org/wiki/Ordinal_numeral

	return fmt.Sprintf("%d%s", n, OrdinalIndicator(n))
}
