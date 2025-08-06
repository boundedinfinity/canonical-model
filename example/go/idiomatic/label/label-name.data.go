package label

import (
	"fmt"
	"reflect"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type labelNames []LabelName

func (t labelNames) IsZero(item LabelName) bool {
	return reflect.DeepEqual(item, zeroLabelName)
}

func (t labelNames) Find(s string) (LabelName, bool) {
	fn := func(_ int, name LabelName) bool {
		return stringer.EqualIgnoreCase(name.Name, s) ||
			stringer.EqualIgnoreCase(s, name.Abbreviation)
	}

	return slicer.FindFn(fn, t...)
}

func (t labelNames) MustFind(s string) LabelName {
	if found, ok := t.Find(s); !ok {
		panic(fmt.Errorf("can't find label name %v", s))
	} else {
		return found
	}
}

var (
	zeroLabelName LabelName

	LabelNames = labelNames{
		{
			Id:           id.Ids.MustParse("B12E7628-35B0-4AE5-B2D9-0F34F4828229"),
			Abbreviation: "dob",
			Name:         "Date of Birth",
		},
		{
			Id:           id.Ids.MustParse("453F12F8-90C3-487C-BF96-53E281ABE143"),
			Abbreviation: "dod",
			Name:         "Date of Death",
		},
		{
			Id:   id.Ids.MustParse("9102C190-4093-4C73-BFE5-1FBA1CA41DE7"),
			Name: "Expiration Date",
		},
		{
			Id:   id.Ids.MustParse("B64458ED-7B79-483E-9087-613BABB7A165"),
			Name: "Formation Date",
		},
		{
			Id:   id.Ids.MustParse("FA64EF80-C3BC-4197-9A87-1FE1728A1100"),
			Name: "Dissolution Date",
		},
		{
			Id:           id.Ids.MustParse("51ED6E00-D565-471F-87FD-479773C1382B"),
			Name:         "End of Life",
			Abbreviation: "eol",
		},
		{
			Id:   id.Ids.MustParse("51ED6E00-D565-471F-87FD-479773C1382B"),
			Name: "Warranty Begin",
		},
		{
			Id:   id.Ids.MustParse("B6CB6E80-6179-42AA-B582-755FE640EA02"),
			Name: "Warranty End",
		},
	}

	Lookup = map[string]LabelName{}
)

func init() {
	for _, name := range LabelNames {
		if !stringer.IsEmpty(name.Abbreviation) {
			if _, ok := Lookup[name.Abbreviation]; !ok {
				Lookup[name.Abbreviation] = name
			} else {
				panic(fmt.Errorf("already contains %v", name.Abbreviation))
			}
		}

		if _, ok := Lookup[name.Name]; !ok {
			Lookup[name.Name] = name
		} else {
			panic(fmt.Errorf("already contains %v", name.Name))
		}

		if _, ok := Lookup[stringer.Lowercase(name.Name)]; !ok {
			Lookup[stringer.Lowercase(name.Name)] = name
		} else {
			panic(fmt.Errorf("already contains %v", stringer.Lowercase(name.Name)))
		}
	}
}
