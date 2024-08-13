package people

import (
	"fmt"
	"reflect"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type suffixes []Suffix

func (t suffixes) IsZero(item Suffix) bool {
	return reflect.DeepEqual(item, zeroSuffix)
}

func (t suffixes) Find(s string) (Suffix, bool) {
	fn := func(_ int, suffix Suffix) bool {
		return stringer.EqualIgnoreCase(suffix.Text, s) ||
			stringer.ContainsAnyIgnoreCase(s, suffix.Abbreviation...)
	}

	return slicer.FindFn(fn, t...)
}

func (t suffixes) MustFind(s string) Suffix {
	if found, ok := t.Find(s); !ok {
		panic(fmt.Errorf("can't find suffix %v", s))
	} else {
		return found
	}
}

var (
	zeroSuffix Suffix
	Suffixes   = suffixes{
		{
			Id:           id.Ids.MustParse("d4033138-e658-4b34-a2fa-55aefeff1250"),
			Text:         "Junior",
			Abbreviation: []string{"Jr."},
		},
	}
)
