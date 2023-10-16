package people_test

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/people"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	// person2 = people.BuildPerson().
	// 	Id("792ce7f5-0586-42a4-ac2f-7f3fd276c3c5").
	// 	Name(people.BuildName().
	// 		Id("d34a65f3-9761-43cf-ae31-1711c307a355").
	// 		Must(),
	// 	).Must()

	person1 = people.Person{
		Id: uuid.MustParse("792ce7f5-0586-42a4-ac2f-7f3fd276c3c5"),
		Name: people.Name{
			Id: uuid.MustParse("d34a65f3-9761-43cf-ae31-1711c307a355"),
			Prefixes: []people.Prefix{
				{
					Id:           uuid.MustParse("b5a5b96c-5955-44c4-8d53-2c7fd1090a02"),
					Text:         "Mister",
					Abbreviation: []string{"Mr."},
				},
			},
			GivenNames:  []string{"James"},
			FamilyNames: []string{"Bond"},
			Suffixes: []people.Suffix{
				{
					Id:           uuid.MustParse("d4033138-e658-4b34-a2fa-55aefeff1250"),
					Text:         "Junior",
					Abbreviation: []string{"Jr."},
				},
			},
			Format: people.NameFormats.GivenNameFamilyName,
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

	fmt.Println("======================")
	fmt.Printf("%v\n", j)
	fmt.Println("======================")

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
                    "format":null
                }
            ],
            "ordering":"GivenNameFamilyName"
        },
        "aliases":null
    }`)
}
