package subject

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person"
)

var _ Subject = &Person{}

type Person struct {
	Person person.Person
}

func (this Person) GetId() ider.Id {
	return this.Person.Id
}

func (_ Person) GetKind() Kind {
	return Kinds.Person
}
