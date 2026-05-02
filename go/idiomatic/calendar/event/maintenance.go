package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

// MaintenanceEventModel represents a maintenance event.  This is similar to the todo but requires more explaination
// and has a number steps that need to be completed. Think of this like and Standard Operating Procedure (SOP).
func Maintenance(name string, abbreviations []string, description string) Event {
	return Event{
		Id:   ider.New(),
		Kind: Kinds.Maintenance,
	}
}
