package event

import "github.com/boundedinfinity/canonical-model/go/idiomatic/ider"

type Event interface {
	GetKind() Kind
	GetId() ider.Id
}
