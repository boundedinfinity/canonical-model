package number

import "github.com/boundedinfinity/canonical_model/idiomatic/util/errorer"

var (
	ErrNumber = errorer.New("number")
)

type Number interface {
	GetKind() Kind
}
