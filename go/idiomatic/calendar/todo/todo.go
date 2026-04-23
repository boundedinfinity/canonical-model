package todo

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/calendar/event"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

var _ event.Event = TodoEventModel{}

// TodoEventModel represents a to-do item.  Typically used for things like tasks or reminders.
type TodoEventModel struct {
}

func (this TodoEventModel) GetId() ider.Id {
	panic("unimplemented")
}

func (_ TodoEventModel) GetKind() event.Kind {
	return event.Kinds.Todo
}
