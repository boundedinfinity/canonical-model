package label

import (
	"fmt"

	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func NewGroup(name, description string, labels ...LabelModel) LabelGroupModel {
	m := LabelGroupModel{
		Name:        name,
		Description: description,
	}

	m.Add(labels...)

	return m
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type LabelGroupModel struct {
	Id          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Labels      Labels  `json:"labels"`
}

func (this *LabelGroupModel) Add(labels ...LabelModel) {
	this.Labels.Add(labels...)
}

func (this LabelGroupModel) Has(label LabelModel) bool {
	return this.Labels.Has(label)
}

func (this LabelGroupModel) Intersection(group LabelGroupModel) LabelGroupModel {
	return LabelGroupModel{
		Name:   fmt.Sprintf("%s intersects with %s", this.Name, group.Name),
		Labels: this.Labels.Intersection(group.Labels...),
	}
}

func (this LabelGroupModel) Union(group LabelGroupModel) LabelGroupModel {
	return LabelGroupModel{
		Name:   fmt.Sprintf("%s unions with %s", this.Name, group.Name),
		Labels: this.Labels.Union(group.Labels...),
	}
}
