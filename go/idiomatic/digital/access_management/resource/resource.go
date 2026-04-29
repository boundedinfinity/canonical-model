package resource

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrResource = errorer.New("resource error")
)

type Resource interface {
	GetKind() Kind
	GetId() ider.Id
}
