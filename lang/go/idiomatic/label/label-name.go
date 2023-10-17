package label

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/google/uuid"
)

type LabelName struct {
	Id           uuid.UUID `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Abbreviation string    `json:"abbreviation,omitempty"`
	Description  string    `json:"description,omitempty"`
}

var (
	Data   = []LabelName{}
	Lookup = map[string]LabelName{}
)

func init() {
	Data = append(Data, LabelName{
		Id:           uuid.MustParse("B12E7628-35B0-4AE5-B2D9-0F34F4828229"),
		Abbreviation: "dob",
		Name:         "Date of Birth",
	})

	Data = append(Data, LabelName{
		Id:           uuid.MustParse("453F12F8-90C3-487C-BF96-53E281ABE143"),
		Abbreviation: "dod",
		Name:         "Date of Death",
	})

	Data = append(Data, LabelName{
		Id:   uuid.MustParse("9102C190-4093-4C73-BFE5-1FBA1CA41DE7"),
		Name: "Expiration Date",
	})

	Data = append(Data, LabelName{
		Id:   uuid.MustParse("B64458ED-7B79-483E-9087-613BABB7A165"),
		Name: "Business Formation Date",
	})

	processNames()
}

func processNames() {
	for _, name := range Data {
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
