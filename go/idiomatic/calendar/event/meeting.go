package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

// MeetingEventModel represents a meeting.  Typically used for things like a work meetings or similar.
func Meeting(name string, abbreviations []string, description string) Event {
	return Event{
		Id:   ider.New(),
		Kind: Kinds.Meeting,
	}
}
