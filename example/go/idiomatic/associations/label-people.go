package associations

import (
	"github.com/boundedinfinity/schema/idiomatic/label"
	"github.com/boundedinfinity/schema/idiomatic/people"
)

type LabelToPerson struct {
	Label  label.NameOnlyLabel
	Person people.Person
}
