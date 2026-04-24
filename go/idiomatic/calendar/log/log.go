package log

import (
	"time"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/calendar/event"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/rfc3339date"
)

var _ event.Event = LogEventModel{}

// LogEventModel represents an event that happened which needs to be tracked or remembered.  Typically used for things
// like the last time an air filter was changed or when you spoke with someome on the phone last.
type LogEventModel struct {
	Id            ider.Id  `json:"id"`
	Name          string   `json:"name"`
	Abbreviations []string `json:"abbreviations"`
	Description   string   `json:"description"`
	When          rfc3339date.Rfc3339DateTime
}

func (this LogEventModel) String() string {
	return this.Name
}

func Happened(name string, abbreviations []string, description string) LogEventModel {
	return LogEventModel{
		Name:          name,
		Abbreviations: abbreviations,
		Description:   description,
		When:          rfc3339date.NewDateTime(time.Now()),
	}
}

// TODO: Want a way to track a thread of related log events.  Like when you're logging
// a process over time, and want to track how long it took.  This may be things like
// a home repair project with where you need to speak with mulitple contractors and need
// to track when and who you spoke with and maybe what was discussed.

func (this LogEventModel) GetId() ider.Id {
	return this.Id
}

func (_ LogEventModel) GetKind() event.Kind {
	return event.Kinds.Log
}
