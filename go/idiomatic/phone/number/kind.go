package number

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/marshal"
	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

const (
	MarshalId = "phone.number"
)

var (
	ErrKind   = errorer.New("phone.number.kind error")
	errKindFn = errorer.Func(ErrKind)
)

type Kind string

func (this *Kind) UnmarshalJSON(bytes []byte) error {
	var err error
	*this, err = marshal.UnmarshalJsonFromSlice(bytes, Kinds.All(), errKindFn)
	return err
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type kinds struct {
	Unknown Kind
	Nanp    Kind
}

func (this kinds) All() []Kind {
	return []Kind{
		this.Unknown,
		this.Nanp,
	}
}

var Kinds = kinds{
	Unknown: "phone.number.unknown",
	Nanp:    "phone.number.nanp",
}
