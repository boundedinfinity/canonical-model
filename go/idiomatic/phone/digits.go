package phone

type digits struct {
	dm map[string]Digit
	n2 Digit
	n3 Digit
	n4 Digit
	n5 Digit
	n6 Digit
	n7 Digit
	n8 Digit
	n9 Digit
}

var Digits = digits{
	n2: Digit{Key: "2", Mnemonics: []string{"A", "B", "C"}},
	n3: Digit{Key: "3", Mnemonics: []string{"D", "E", "F"}},
	n4: Digit{Key: "4", Mnemonics: []string{"G", "H", "I"}},
	n5: Digit{Key: "5", Mnemonics: []string{"J", "K", "L"}},
	n6: Digit{Key: "6", Mnemonics: []string{"M", "Q", "R", "S"}},
	n7: Digit{Key: "7", Mnemonics: []string{"P", "N", "O"}},
	n8: Digit{Key: "8", Mnemonics: []string{"T", "U", "V"}},
	n9: Digit{Key: "9", Mnemonics: []string{"W", "X", "Y", "Z"}},
}

func (this digits) All() []Digit {
	return []Digit{
		this.n2,
		this.n3,
		this.n4,
		this.n5,
		this.n6,
		this.n7,
		this.n8,
		this.n9,
	}
}

func init() {
	Digits.dm = map[string]Digit{
		"2": Digits.n2,
		"A": Digits.n2,
		"a": Digits.n2,
		"B": Digits.n2,
		"b": Digits.n2,
		"C": Digits.n2,
		"c": Digits.n2,
		"3": Digits.n3,
		"D": Digits.n3,
		"d": Digits.n3,
		"E": Digits.n3,
		"e": Digits.n3,
		"F": Digits.n3,
		"f": Digits.n3,
		"4": Digits.n4,
		"G": Digits.n4,
		"g": Digits.n4,
		"H": Digits.n4,
		"h": Digits.n4,
		"I": Digits.n4,
		"i": Digits.n4,
		"5": Digits.n5,
		"J": Digits.n5,
		"j": Digits.n5,
		"K": Digits.n5,
		"k": Digits.n5,
		"L": Digits.n5,
		"l": Digits.n5,
		"6": Digits.n6,
		"M": Digits.n6,
		"m": Digits.n6,
		"Q": Digits.n6,
		"q": Digits.n6,
		"R": Digits.n6,
		"r": Digits.n6,
		"S": Digits.n6,
		"s": Digits.n6,
		"7": Digits.n7,
		"P": Digits.n7,
		"p": Digits.n7,
		"N": Digits.n7,
		"n": Digits.n7,
		"O": Digits.n7,
		"o": Digits.n7,
		"8": Digits.n8,
		"T": Digits.n8,
		"t": Digits.n8,
		"U": Digits.n8,
		"u": Digits.n8,
		"V": Digits.n8,
		"v": Digits.n8,
		"9": Digits.n9,
		"W": Digits.n9,
		"w": Digits.n9,
		"X": Digits.n9,
		"x": Digits.n9,
		"Y": Digits.n9,
		"y": Digits.n9,
		"Z": Digits.n9,
		"z": Digits.n9,
	}
}
