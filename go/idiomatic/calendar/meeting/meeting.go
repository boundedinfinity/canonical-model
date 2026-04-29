package meeting

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/calendar/event"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

var _ event.Event = MeetingEventModel{}

// MeetingEventModel represents a meeting.  Typically used for things like a work meetings or similar.
type MeetingEventModel struct {
}

func (this MeetingEventModel) GetId() ider.Id {
	panic("unimplemented")
}

func (_ MeetingEventModel) GetKind() event.Kind {
	return event.Kinds.Meeting
}
