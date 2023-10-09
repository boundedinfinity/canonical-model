package people_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/optioner"
	"github.com/boundedinfinity/schema/people"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	person1 = people.Person{
		Id: optioner.Some(uuid.MustParse("792ce7f5-0586-42a4-ac2f-7f3fd276c3c5")),
		Name: people.Name{
			Id: optioner.Some(uuid.MustParse("d34a65f3-9761-43cf-ae31-1711c307a355")),
			Prefixes: optioner.Some([]people.Prefix{
				{
					Id:           optioner.Some(uuid.MustParse("b5a5b96c-5955-44c4-8d53-2c7fd1090a02")),
					Text:         "Mister",
					Abbreviation: optioner.Some([]string{"Mr."}),
				},
			}),
			GivenNames:  []string{"James"},
			FamilyNames: []string{"Bond"},
			Suffixes: optioner.Some([]people.Suffix{
				{
					Id:           optioner.Some(uuid.MustParse("d4033138-e658-4b34-a2fa-55aefeff1250")),
					Text:         "Junior",
					Abbreviation: optioner.Some([]string{"Jr."}),
					Format:       optioner.Some(people.Abbreviation),
				},
			}),
			Ordering: optioner.Some(people.GivenNameFamilyName),
		},
	}
)

func Test_Person_JSON_String(t *testing.T) {
	actual := person1.String()

	assert.Equal(t, `Mr. James Bond Jr.`, actual)
}

func Test_Person_JSON_Serialize(t *testing.T) {
	bs, err := json.Marshal(person1)
	j := string(bs)

	// fmt.Println("======================")
	// fmt.Printf("%v", j)
	// fmt.Println("======================")

	assert.Nil(t, err)
	assert.JSONEq(t, j, `
	{
		"id":"792ce7f5-0586-42a4-ac2f-7f3fd276c3c5",
		"name":{
			"id":"d34a65f3-9761-43cf-ae31-1711c307a355",
			"prefixes":[
				{
					"id":"b5a5b96c-5955-44c4-8d53-2c7fd1090a02",
					"text":"Mister",
					"abbreviation":["Mr."],
					"description":null,
					"format":null
				}
			],
			"givenNames":["James"],
			"familyNames":["Bond"],
			"suffixes":[
				{
					"id":"d4033138-e658-4b34-a2fa-55aefeff1250",
					"text":"Junior",
					"abbreviation":["Jr."],
					"description":null,
					"format":"Abbreviation"
				}
			],
			"ordering":"Family name, Given name"
		},
		"alsoKnownAs":null
	}`)
}
