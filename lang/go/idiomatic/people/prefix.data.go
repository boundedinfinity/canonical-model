package people

import (
	"fmt"
	"reflect"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type prefixes []Prefix

func (t prefixes) IsZero(item Prefix) bool {
	return reflect.DeepEqual(item, zeroPrefix)
}

func (t prefixes) Find(s string) (Prefix, bool) {
	fn := func(prefix Prefix) bool {
		return stringer.EqualIgnoreCase(prefix.Text, s) ||
			stringer.ContainsAnyIgnoreCase(s, prefix.Abbreviation...)
	}

	return slicer.FindFn(fn, t...)
}

func (t prefixes) MustFind(s string) Prefix {
	if found, ok := t.Find(s); !ok {
		panic(fmt.Errorf("can't find prefix %v", s))
	} else {
		return found
	}
}

var (
	zeroPrefix Prefix
	Prefixes   = prefixes{
		{
			Id:           id.MustParse("b5a5b96c-5955-44c4-8d53-2c7fd1090a02"),
			Text:         "Mister",
			Abbreviation: []string{"Mr."},
		},
	}
)
