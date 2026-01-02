package associations

import (
	"github.com/boundedinfinity/schema/idiomatic/label"
	"github.com/boundedinfinity/schema/idiomatic/people"
)

type LabelToPerson struct {
	Label  label.Label
	Person people.Person
}
