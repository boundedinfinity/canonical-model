package span

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/calendar/event"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

var _ event.Event = SpanEventModel{}

// SpanEventModel represents a span of time.  Typically used for things like a vacation or a trip which spans
// multiple days or even weeks.
type SpanEventModel struct {
}

func (this SpanEventModel) GetId() ider.Id {
	panic("unimplemented")
}

func (_ SpanEventModel) GetKind() event.Kind {
	return event.Kinds.Span
}
