package event

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

func Holiday() Event {
	return Event{
		Id:   ider.New(),
		Kind: Kinds.Holiday,
	}
}
