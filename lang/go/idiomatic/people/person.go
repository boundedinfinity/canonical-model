package people

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

type Person struct {
	Id         id.Id                   `json:"id,omitempty"`
	Name       Name                    `json:"name,omitempty"`
	Pseudonyms []Name                  `json:"pseudonyms,omitempty"`
	BirthDate  rfc3339date.Rfc3339Date `json:"birth-date,omitempty"`
	DeathDate  rfc3339date.Rfc3339Date `json:"death-date,omitempty"`
}

func (t Person) String() string {
	return t.Name.String()
}

func (t Person) Pseudonym() string {
	var items []string

	for _, item := range t.Pseudonyms {
		items = append(items, item.String())
	}

	return stringer.Join(", ", items...)
}

// ///////////////////////////////////////////////////
// Builder
// ///////////////////////////////////////////////////

// type personFn func(*Person) error

// type personBuilder struct {
// 	fns []personFn
// }

// func BuildPerson() *personBuilder {
// 	return &personBuilder{
// 		fns: make([]personFn, 0),
// 	}
// }

// func (t *personBuilder) Build() (Person, error) {
// 	var person Person

// 	for _, fn := range t.fns {
// 		if err := fn(&person); err != nil {
// 			return person, err
// 		}
// 	}

// 	return person, nil
// }

// func (t *personBuilder) Must() Person {
// 	person, err := t.Build()

// 	if err != nil {
// 		panic(err)
// 	}

// 	return person
// }

// func (t *personBuilder) Id(v string) *personBuilder {
// 	t.fns = append(t.fns, func(p *Person) error {
// 		id, err := id.Parse(v)

// 		if err != nil {
// 			return err
// 		}

// 		p.Id = id
// 		return nil
// 	})

// 	return t
// }

// func (t *personBuilder) Name(v Name) error {
// 	t.fns = append(t.fns, func(p *Person) error {
// 		p.Name = v
// 		return nil
// 	})

// 	return nil
// }

// func (t *personBuilder) AlsoKnownAs(v Name) error {
// 	t.fns = append(t.fns, func(p *Person) error {
// 		p.Pseudonyms = append(p.Pseudonyms, v)
// 		return nil
// 	})

// 	return nil
// }
