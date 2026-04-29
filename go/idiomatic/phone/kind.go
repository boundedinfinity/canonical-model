package phone

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/marshal"
	"github.com/boundedinfinity/go-commoner/idiomatic/errorer"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var (
	ErrKind   = errorer.New("phone.kind error")
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
	Unknown  Kind
	Mobile   Kind
	Fax      Kind
	Landline Kind
}

func (this kinds) All() []Kind {
	return []Kind{
		this.Unknown,
		this.Mobile,
		this.Fax,
		this.Landline,
	}
}

var Kinds = kinds{
	Unknown:  "phone.kind.unknown",
	Mobile:   "phone.kind.mobile",
	Fax:      "phone.kind.fax",
	Landline: "phone.kind.landline",
}
