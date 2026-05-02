package event

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

// SpanEventModel represents a span of time.  Typically used for things like a vacation or a trip which spans
// multiple days or even weeks.
func Span(name string, abbreviations []string, description string) Event {
	return Event{
		Id:   ider.New(),
		Kind: Kinds.Span,
	}
}
