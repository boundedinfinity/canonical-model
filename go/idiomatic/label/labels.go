package label

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/util/slicer"
)

type Labels []LabelModel

func (this Labels) Names() []string {
	return slicer.Map(this, getName)
}

func (this Labels) String() string {
	return stringer.JoinFunc(this, ", ", getName)
}

func (this Labels) Has(label LabelModel) bool {
	return slicer.ContainsFunc(this, equalsAny(label))
}

func (this *Labels) Add(labels ...LabelModel) {
	*this = this.Union(labels...)
}

func (this Labels) Intersection(labels ...LabelModel) []LabelModel {
	return slicer.IntersectionFunc(this, labels, getId)
}

func (this Labels) Difference(labels ...LabelModel) []LabelModel {
	return slicer.DifferenceFunc(this, labels, getId)
}

func (this Labels) Union(labels ...LabelModel) []LabelModel {
	return slicer.UnionFunc(this, labels, getId)
}
