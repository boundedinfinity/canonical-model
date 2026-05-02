package event

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

func Free(name string, abbreviations []string, description string) Event {
	return Event{
		Id:   ider.New(),
		Kind: Kinds.Free,
	}
}
