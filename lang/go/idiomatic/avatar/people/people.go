package people

import (
	"github.com/boundedinfinity/schema/idiomatic/avatar/persistence"
	"github.com/boundedinfinity/schema/idiomatic/people"
)

type Name persistence.Item[people.Name]
type Person persistence.Item[people.Person]
