package subject

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrSubject = errorer.New("subject error")
)

type Subject interface {
	GetKind() Kind
	GetId() ider.Id
}
