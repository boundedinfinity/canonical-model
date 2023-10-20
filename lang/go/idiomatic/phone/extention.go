package phone

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

// https://en.wikipedia.org/wiki/Extension_(telephone)

type Extention struct {
	Items []ExtentionItem
}

func (t *Extention) Digit(digit int) {
	d, _ := NewDigit(digit)
	t.Items = append(t.Items, d)
}

func (t *Extention) Pause2Seconds() {
	t.Items = append(t.Items, AtCommands.Pause2Seconds)
}

func (t Extention) Has() bool {
	return len(t.Items) > 0
}

func (t Extention) String() string {
	if !t.Has() {
		return ""
	}

	var items []string

	for _, item := range t.Items {
		items = append(items, item.String())
	}

	return "ext " + stringer.Join("", items...)
}

type ExtentionItem interface {
	String() string
}
