package resource

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/idiomatic/person"
)

var _ Resource = &Person{}

type Person struct {
	Person person.PersonModel
}

func (this Person) GetId() ider.Id {
	return this.Person.Id
}

func (_ Person) GetKind() Kind {
	return Kinds.Person
}
