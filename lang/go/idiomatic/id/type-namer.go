package id

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/boundedinfinity/go-commoner/idiomatic/reflecter"
)

type TypeNamer interface {
	TypeName() string
}

var TypeNamers = typeNamers{}

type typeNamers struct{}

func (t typeNamers) Dotted(val any) string {
	qname := reflecter.Instances.QualifiedName(val)
	n1 := pather.Paths.Base(pather.Paths.Dir(qname))
	n2 := pather.Paths.Base(qname)
	return fmt.Sprintf("%v.%v", n1, n2)
}
