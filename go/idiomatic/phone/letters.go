package phone

import "github.com/boundedinfinity/go-commoner/idiomatic/stringer"

// https://en.wikipedia.org/wiki/Telephone_keypad

var (
	letter2number = map[string]int{
		"A": 2,
		"B": 2,
		"C": 2,
		"D": 3,
		"E": 3,
		"F": 3,
		"G": 4,
		"H": 4,
		"I": 4,
		"J": 5,
		"K": 5,
		"L": 5,
		"M": 6,
		"N": 6,
		"O": 6,
		"P": 7,
		"Q": 7,
		"R": 7,
		"S": 7,
		"T": 8,
		"U": 8,
		"V": 8,
		"W": 9,
		"X": 9,
		"Y": 9,
		"Z": 9,
	}

	number2Letter = map[int][]string{}
)

func init() {
	for k, v := range letter2number {
		letter2number[stringer.ToLower(k)] = v

		if _, ok := number2Letter[v]; !ok {
			number2Letter[v] = []string{k}
		} else {
			number2Letter[v] = append(number2Letter[v], k)
		}
	}
}
