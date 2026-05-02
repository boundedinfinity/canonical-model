package label

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type Labels []Label

func (this Labels) Names() []string {
	return slicer.Map(getName, this...)
}

func (this Labels) String() string {
	return stringer.JoinFunc(this, ", ", getName)
}

func (this Labels) Has(label Label) bool {
	return slicer.ContainsFunc(equalsAny(label), this...)
}

func (this *Labels) Add(labels ...Label) {
	for _, label := range labels {
		if !this.Has(label) {
			*this = append(*this, label)
		}
	}
}

func (this *Labels) Group(group Group) {
	this.Add(group.Labels...)
}
