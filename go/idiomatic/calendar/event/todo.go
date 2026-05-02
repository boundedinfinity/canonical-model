package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

// TodoEventModel represents a to-do item.  Typically used for things like tasks or reminders.
func Todo(name string, abbreviations []string, description string) Event {
	return Event{
		Id:   ider.New(),
		Kind: Kinds.Todo,
	}
}
