package label

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func NewLabelManager(options LabelManagerOptions, labels ...LabelModel) *LabelManager {
	manager := &LabelManager{options: options}
	manager.Add(labels...)
	return manager
}

type LabelManagerOptions struct {
	ErrorOnDuplicate bool
}

type LabelManager struct {
	options  LabelManagerOptions
	id2label map[ider.Id]LabelModel
}

func (this LabelManager) Labels() []LabelModel {
	return mapper.Values(this.id2label)
}

func (this LabelManager) Names() []string {
	return slicer.Map(
		func(label LabelModel) string { return label.Name },
		this.Labels()...,
	)
}

var (
	ErrLabelManagerDuplicate   = errors.New("duplicate label")
	errLabelManagerDuplicatefn = func(label LabelModel) error {
		return fmt.Errorf("%s : %w", label, ErrLabelManagerDuplicate)
	}
)

func (this *LabelManager) Add(labels ...LabelModel) error {
	if this.id2label == nil {
		this.id2label = map[ider.Id]LabelModel{}
	}

	for _, label := range labels {
		if _, ok := this.id2label[label.Id]; ok {
			if this.options.ErrorOnDuplicate {
				return errLabelManagerDuplicatefn(label)
			}
		} else {
			this.id2label[label.Id] = label
		}
	}

	return nil
}

func (this *LabelManager) Remove(labels ...LabelModel) {
	for _, label := range labels {
		delete(this.id2label, label.Id)
	}
}

func (this LabelManager) Find(term string) (LabelModel, bool) {
	fn := func(label LabelModel) bool {
		return stringer.EqualIgnoreCase(label.Name, term) ||
			stringer.EqualIgnoreCase(label.Abbreviation, term)
	}

	return slicer.FindFn(fn, this.Labels()...)
}

func (this LabelManager) MustFind(name string) LabelModel {
	label, ok := this.Find(name)

	if !ok {
		panic(fmt.Errorf("no label found for name: %s", name))
	}

	return label
}
