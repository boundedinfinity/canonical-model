package label

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

func NewLabelManager(options LabelManagerOptions, labels ...Label) *LabelManager {
	manager := &LabelManager{options: options}
	manager.Add(labels...)
	return manager
}

type LabelManagerOptions struct {
	ErrorOnDuplicate bool
}

type LabelManager struct {
	options  LabelManagerOptions
	id2label map[id.Id]Label
}

func (this LabelManager) Labels() []Label {
	return mapper.Values(this.id2label)
}

func (this LabelManager) Names() []string {
	return slicer.Map(
		func(_ int, label Label) string { return label.Name },
		this.Labels()...,
	)
}

var (
	ErrLabelManagerDuplicate   = errors.New("duplicate label")
	errLabelManagerDuplicatefn = func(label Label) error {
		return fmt.Errorf("%s : %w", label, ErrLabelManagerDuplicate)
	}
)

func (this *LabelManager) Add(labels ...Label) error {
	if this.id2label == nil {
		this.id2label = map[id.Id]Label{}
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

func (this *LabelManager) Remove(labels ...Label) {
	for _, label := range labels {
		delete(this.id2label, label.Id)
	}
}

func (this LabelManager) Find(term string) (Label, bool) {
	fn := func(_ int, label Label) bool {
		return stringer.EqualIgnoreCase(label.Name, term) ||
			stringer.EqualIgnoreCase(label.Abbreviation, term)
	}

	return slicer.FindFn(fn, this.Labels()...)
}

func (this LabelManager) MustFind(name string) Label {
	label, ok := this.Find(name)

	if !ok {
		panic(fmt.Errorf("no label found for name: %s", name))
	}

	return label
}
