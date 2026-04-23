package maintenance

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/calendar/event"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

var _ event.Event = MaintenanceEventModel{}

// MaintenanceEventModel represents a maintenance event.  This is similar to the todo but requires more explaination
// and has a number steps that need to be completed. Think of this like and Standard Operating Procedure (SOP).
type MaintenanceEventModel struct {
}

// TODO: May be merged with todo

func (this MaintenanceEventModel) GetId() ider.Id {
	panic("unimplemented")
}

func (_ MaintenanceEventModel) GetKind() event.Kind {
	return event.Kinds.Maintenance
}
