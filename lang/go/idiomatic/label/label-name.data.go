package label

import (
	"fmt"
	"reflect"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type labelNames []LabelName

func (t labelNames) IsZero(item LabelName) bool {
	return reflect.DeepEqual(item, zeroLabelName)
}

var (
	zeroLabelName LabelName

	LabelNames = labelNames{
		{
			Id:           id.MustParse("B12E7628-35B0-4AE5-B2D9-0F34F4828229"),
			Abbreviation: "dob",
			Name:         "Date of Birth",
		},
		{
			Id:           id.MustParse("453F12F8-90C3-487C-BF96-53E281ABE143"),
			Abbreviation: "dod",
			Name:         "Date of Death",
		},
		{
			Id:   id.MustParse("9102C190-4093-4C73-BFE5-1FBA1CA41DE7"),
			Name: "Expiration Date",
		},
		{
			Id:   id.MustParse("B64458ED-7B79-483E-9087-613BABB7A165"),
			Name: "Business Formation Date",
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

		if _, ok := Lookup[stringer.ToLower(name.Name)]; !ok {
			Lookup[stringer.ToLower(name.Name)] = name
		} else {
			panic(fmt.Errorf("already contains %v", stringer.ToLower(name.Name)))
		}
	}
}
