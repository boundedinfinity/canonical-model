package label

import (
	"strings"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type LabelRepository interface {
	Add(label Label)
	Get(id ider.Id) (*Label, bool)
	Find(term string) ([]*Label, bool)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ LabelRepository = &memoryRepository{}

func NewMemoryRepository(labels []Label) LabelRepository {
	r := &memoryRepository{}

	for _, label := range labels {
		r.Add(label)
	}

	return r
}

type memoryRepository struct {
	idMap   map[ider.Id]*Label
	nameMap map[string]*Label
}

func (this *memoryRepository) Add(label Label) {
	if _, ok := this.idMap[label.Id]; !ok {
		this.idMap[label.Id] = &label
	}

	if _, ok := this.nameMap[label.Name]; !ok {
		this.nameMap[label.Name] = &label
	}
}

func (this memoryRepository) Find(term string) ([]*Label, bool) {
	var found []*Label

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

func (this memoryRepository) Get(id ider.Id) (*Label, bool) {
	label, ok := this.idMap[id]
	return label, ok
}
