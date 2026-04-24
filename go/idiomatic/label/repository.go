package label

import (
	"strings"

	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type LabelRepository interface {
	Add(label LabelModel)
	Get(id ider.Id) (*LabelModel, bool)
	Find(term string) ([]*LabelModel, bool)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ LabelRepository = &memoryRepository{}

func NewMemoryRepository(labels []LabelModel) LabelRepository {
	r := &memoryRepository{}

	for _, label := range labels {
		r.Add(label)
	}

	return r
}

type memoryRepository struct {
	idMap   map[ider.Id]*LabelModel
	nameMap map[string]*LabelModel
}

func (this *memoryRepository) Add(label LabelModel) {
	if _, ok := this.idMap[label.Id]; !ok {
		this.idMap[label.Id] = &label
	}

	if _, ok := this.nameMap[label.Name]; !ok {
		this.nameMap[label.Name] = &label
	}
}

func (this memoryRepository) Find(term string) ([]*LabelModel, bool) {
	var found []*LabelModel

	if label, ok := this.nameMap[term]; ok {
		found = append(found, label)
	}

	if id, err := ider.Parse(term); err != nil {
		if label, ok := this.idMap[id]; ok {
			found = append(found, label)
		}
	}

	if len(found) <= 0 {
		for _, label := range this.idMap {
			if strings.Contains(term, label.Name) {
				found = append(found, label)
			}
		}
	}

	return found, len(found) > 0
}

func (this memoryRepository) Get(id ider.Id) (*LabelModel, bool) {
	label, ok := this.idMap[id]
	return label, ok
}
