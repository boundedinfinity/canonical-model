package people

import (
	"github.com/boundedinfinity/optioner"
	"github.com/google/uuid"
)

// ///////////////////////////////////////////////////
// Model
// ///////////////////////////////////////////////////

type Person struct {
	Id      optioner.Option[uuid.UUID] `json:"id"`
	Name    Name                       `json:"name"`
	Aliases optioner.Option[[]Name]    `json:"aliases"`
}

func (t Person) String() string {
	return t.Name.String()
}

// ///////////////////////////////////////////////////
// Builder
// ///////////////////////////////////////////////////

type personFn func(*Person) error

type personBuilder struct {
	fns []personFn
}

func BuildPerson() *personBuilder {
	return &personBuilder{
		fns: make([]personFn, 0),
	}
}

func (t *personBuilder) Build() (Person, error) {
	var person Person

	for _, fn := range t.fns {
		if err := fn(&person); err != nil {
			return person, err
		}
	}

	return person, nil
}

func (t *personBuilder) Must() Person {
	person, err := t.Build()

	if err != nil {
		panic(err)
	}

	return person
}

func (t *personBuilder) Id(v string) *personBuilder {
	t.fns = append(t.fns, func(p *Person) error {
		id, err := uuid.Parse(v)

		if err != nil {
			return err
		}

		p.Id = optioner.Some(id)
		return nil
	})

	return t
}

func (t *personBuilder) Name(v Name) error {
	t.fns = append(t.fns, func(p *Person) error {
		p.Name = v
		return nil
	})

	return nil
}

func (t *personBuilder) AlsoKnownAs(v Name) error {
	t.fns = append(t.fns, func(p *Person) error {
		if p.Aliases.Empty() {
			p.Aliases = optioner.Some(make([]Name, 0))
		}

		p.Aliases = optioner.Some(append(p.Aliases.Get(), v))

		return nil
	})

	return nil
}
