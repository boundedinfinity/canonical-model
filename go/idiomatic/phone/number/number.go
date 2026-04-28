package number

import "github.com/boundedinfinity/go-commoner/idiomatic/errorer"

var (
	ErrNumber = errorer.New("number")
)

type Number interface {
	GetKind() Kind
}
