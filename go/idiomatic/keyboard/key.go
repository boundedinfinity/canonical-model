package keyboard

type Key struct {
	Name string
	Code string
}

func (this Key) Match(target string) bool {
	return this.Code == target
}

var Keys = keys{
	Super:      &Key{Name: "Super", Code: "super"},
	Meta:       &Key{Name: "Meta", Code: "meta"},
	Alt:        &Key{Name: "Alt", Code: "alt"},
	Ctrl:       &Key{Name: "Ctrl", Code: "ctrl"},
	UpperA:     &Key{Name: "UpperA", Code: "A"},
	UpperB:     &Key{Name: "UpperB", Code: "B"},
	UpperC:     &Key{Name: "UpperC", Code: "C"},
	UpperD:     &Key{Name: "UpperD", Code: "D"},
	UpperE:     &Key{Name: "UpperE", Code: "E"},
	UpperF:     &Key{Name: "UpperF", Code: "F"},
	UpperG:     &Key{Name: "UpperG", Code: "G"},
	UpperH:     &Key{Name: "UpperH", Code: "H"},
	UpperI:     &Key{Name: "UpperI", Code: "I"},
	UpperJ:     &Key{Name: "UpperJ", Code: "J"},
	UpperK:     &Key{Name: "UpperK", Code: "K"},
	UpperL:     &Key{Name: "UpperL", Code: "L"},
	UpperM:     &Key{Name: "UpperM", Code: "M"},
	UpperN:     &Key{Name: "UpperN", Code: "N"},
	UpperO:     &Key{Name: "UpperO", Code: "O"},
	UpperP:     &Key{Name: "UpperP", Code: "P"},
	UpperQ:     &Key{Name: "UpperQ", Code: "Q"},
	UpperR:     &Key{Name: "UpperR", Code: "R"},
	UpperS:     &Key{Name: "UpperS", Code: "S"},
	UpperT:     &Key{Name: "UpperT", Code: "T"},
	UpperU:     &Key{Name: "UpperU", Code: "U"},
	UpperV:     &Key{Name: "UpperV", Code: "V"},
	UpperW:     &Key{Name: "UpperW", Code: "W"},
	UpperX:     &Key{Name: "UpperX", Code: "X"},
	UpperY:     &Key{Name: "UpperY", Code: "Y"},
	UpperZ:     &Key{Name: "UpperZ", Code: "Z"},
	LowerA:     &Key{Name: "LowerA", Code: "a"},
	LowerB:     &Key{Name: "LowerB", Code: "b"},
	LowerC:     &Key{Name: "LowerC", Code: "c"},
	LowerD:     &Key{Name: "LowerD", Code: "d"},
	LowerE:     &Key{Name: "LowerE", Code: "e"},
	LowerF:     &Key{Name: "LowerF", Code: "f"},
	LowerG:     &Key{Name: "LowerG", Code: "g"},
	LowerH:     &Key{Name: "LowerH", Code: "h"},
	LowerI:     &Key{Name: "LowerI", Code: "i"},
	LowerJ:     &Key{Name: "LowerJ", Code: "j"},
	LowerK:     &Key{Name: "LowerK", Code: "k"},
	LowerL:     &Key{Name: "LowerL", Code: "l"},
	LowerM:     &Key{Name: "LowerM", Code: "m"},
	LowerN:     &Key{Name: "LowerN", Code: "n"},
	LowerO:     &Key{Name: "LowerO", Code: "o"},
	LowerP:     &Key{Name: "LowerP", Code: "p"},
	LowerQ:     &Key{Name: "LowerQ", Code: "q"},
	LowerR:     &Key{Name: "LowerR", Code: "r"},
	LowerS:     &Key{Name: "LowerS", Code: "s"},
	LowerT:     &Key{Name: "LowerT", Code: "t"},
	LowerU:     &Key{Name: "LowerU", Code: "u"},
	LowerV:     &Key{Name: "LowerV", Code: "v"},
	LowerW:     &Key{Name: "LowerW", Code: "w"},
	LowerX:     &Key{Name: "LowerX", Code: "x"},
	LowerY:     &Key{Name: "LowerY", Code: "y"},
	LowerZ:     &Key{Name: "LowerZ", Code: "z"},
	Num0:       &Key{Name: "Num0", Code: "0"},
	Num1:       &Key{Name: "Num1", Code: "1"},
	Num2:       &Key{Name: "Num2", Code: "2"},
	Num3:       &Key{Name: "Num3", Code: "3"},
	Num4:       &Key{Name: "Num4", Code: "4"},
	Num5:       &Key{Name: "Num5", Code: "5"},
	Num6:       &Key{Name: "Num6", Code: "6"},
	Num7:       &Key{Name: "Num7", Code: "7"},
	Num8:       &Key{Name: "Num8", Code: "8"},
	Num9:       &Key{Name: "Num9", Code: "9"},
	ArrowUp:    &Key{Name: "ArrowUp", Code: "ArrowUp"},
	ArrowDown:  &Key{Name: "ArrowDown", Code: "ArrowDown"},
	ArrowLeft:  &Key{Name: "ArrowLeft", Code: "ArrowLeft"},
	ArrowRight: &Key{Name: "ArrowRight", Code: "ArrowRight"},
}

func init() {
	Keys.Command = Keys.Super
}

type keys struct {
	Super      *Key
	Meta       *Key
	Alt        *Key
	Ctrl       *Key
	Command    *Key
	UpperA     *Key
	UpperB     *Key
	UpperC     *Key
	UpperD     *Key
	UpperE     *Key
	UpperF     *Key
	UpperG     *Key
	UpperH     *Key
	UpperI     *Key
	UpperJ     *Key
	UpperK     *Key
	UpperL     *Key
	UpperM     *Key
	UpperN     *Key
	UpperO     *Key
	UpperP     *Key
	UpperQ     *Key
	UpperR     *Key
	UpperS     *Key
	UpperT     *Key
	UpperU     *Key
	UpperV     *Key
	UpperW     *Key
	UpperX     *Key
	UpperY     *Key
	UpperZ     *Key
	LowerA     *Key
	LowerB     *Key
	LowerC     *Key
	LowerD     *Key
	LowerE     *Key
	LowerF     *Key
	LowerG     *Key
	LowerH     *Key
	LowerI     *Key
	LowerJ     *Key
	LowerK     *Key
	LowerL     *Key
	LowerM     *Key
	LowerN     *Key
	LowerO     *Key
	LowerP     *Key
	LowerQ     *Key
	LowerR     *Key
	LowerS     *Key
	LowerT     *Key
	LowerU     *Key
	LowerV     *Key
	LowerW     *Key
	LowerX     *Key
	LowerY     *Key
	LowerZ     *Key
	Num0       *Key
	Num1       *Key
	Num2       *Key
	Num3       *Key
	Num4       *Key
	Num5       *Key
	Num6       *Key
	Num7       *Key
	Num8       *Key
	Num9       *Key
	ArrowUp    *Key
	ArrowDown  *Key
	ArrowLeft  *Key
	ArrowRight *Key
}
