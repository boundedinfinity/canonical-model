package people

import (
	"github.com/boundedinfinity/schema/idiomatic/avatar/model"
	"github.com/boundedinfinity/schema/idiomatic/people"
)

type PrefixItem model.Item[people.Prefix]
type SuffixItem model.Item[people.Suffix]
type NameItem model.Item[people.Name]
type PersonItem model.Item[people.Person]

type PersonEvent = model.ConcreteEvent[people.Person]
type PrefixEvent = model.ConcreteEvent[people.Prefix]
type SuffixEvent = model.ConcreteEvent[people.Suffix]
type NameEvent = model.ConcreteEvent[people.Name]
