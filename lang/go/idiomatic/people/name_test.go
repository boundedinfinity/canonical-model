package people_test

// var (
// 	person1 = people.Person{
// 		Id: id.MustParse("792ce7f5-0586-42a4-ac2f-7f3fd276c3c5"),
// 		Name: people.Name{
// 			Id:          id.MustParse("d34a65f3-9761-43cf-ae31-1711c307a355"),
// 			Prefixes:    []people.Prefix{people.Prefixes.MustFind("Mr.")},
// 			Firsts:      []string{"James"},
// 			FamilyNames: []string{"Bond"},
// 			Suffixes:    []people.Suffix{people.Suffixes.MustFind("Jr.")},
// 			Format:      people.NameFormats.GivenNameFamilyName,
// 		},
// 	}
// )

// func Test_Person_JSON_String(t *testing.T) {
// 	actual := person1.String()

// 	assert.Equal(t, `Mr. James Bond Jr.`, actual)
// }

// func Test_PersonMessage_JSON_String(t *testing.T) {
// 	bs, err := json.MarshalIndent(people.NewNameMessage(person1.Name), "", "    ")
// 	j := string(bs)

// 	fmt.Println("======================")
// 	fmt.Printf("%v\n", j)
// 	fmt.Println("======================")

// 	assert.Nil(t, err)
// 	assert.JSONEq(t, j, `
// 	{
//         "id":"792ce7f5-0586-42a4-ac2f-7f3fd276c3c5",
//         "name":{
//             "id":"d34a65f3-9761-43cf-ae31-1711c307a355",
//             "prefixes":[
//                 {
//                     "id":"b5a5b96c-5955-44c4-8d53-2c7fd1090a02",
//                     "text":"Mister",
//                     "abbreviation":["Mr."]
//                 }
//             ],
//             "givenNames":["James"],
//             "familyNames":["Bond"],
//             "suffixes":[
//                 {
//                     "id":"d4033138-e658-4b34-a2fa-55aefeff1250",
//                     "text":"Junior",
//                     "abbreviation":["Jr."]
//                 }
//             ],
//             "ordering":"GivenNameFamilyName"
//         }
//     }`)
// }

// func Test_Person_JSON_Serialize(t *testing.T) {
// 	bs, err := json.Marshal(person1)
// 	j := string(bs)

// 	fmt.Println("======================")
// 	fmt.Printf("%v\n", j)
// 	fmt.Println("======================")

// 	assert.Nil(t, err)
// 	assert.JSONEq(t, j, `
// 	{
//         "id":"792ce7f5-0586-42a4-ac2f-7f3fd276c3c5",
//         "name":{
//             "id":"d34a65f3-9761-43cf-ae31-1711c307a355",
//             "prefixes":[
//                 {
//                     "id":"b5a5b96c-5955-44c4-8d53-2c7fd1090a02",
//                     "text":"Mister",
//                     "abbreviation":["Mr."]
//                 }
//             ],
//             "givenNames":["James"],
//             "familyNames":["Bond"],
//             "suffixes":[
//                 {
//                     "id":"d4033138-e658-4b34-a2fa-55aefeff1250",
//                     "text":"Junior",
//                     "abbreviation":["Jr."]
//                 }
//             ],
//             "ordering":"GivenNameFamilyName"
//         }
//     }`)
// }
