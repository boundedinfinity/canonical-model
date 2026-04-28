package label

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type Labels []LabelModel

func (this Labels) Names() []string {
	return slicer.Map(getName, this...)
}

func (this Labels) String() string {
	return stringer.JoinFunc(this, ", ", getName)
}

func (this Labels) Has(label LabelModel) bool {
	return slicer.ContainsFunc(equalsAny(label), this...)
}

func (this *Labels) Add(labels ...LabelModel) {
	// *this = this.Union(labels...)
}

func (this Labels) Intersection(labels ...LabelModel) []LabelModel {
	// return slicer.IntersectionFunc(this, labels, getId)
	return nil
}

func (this Labels) Difference(labels ...LabelModel) []LabelModel {
	// return slicer.DifferenceFunc(this, labels, getId)
	return nil
}

func (this Labels) Union(labels ...LabelModel) []LabelModel {
	// return slicer.UnionFunc(this, labels, getId)
	return nil
}
