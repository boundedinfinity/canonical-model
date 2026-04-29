package bookmark

import (
	"strings"

	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

type BookmarkRepository interface {
	Add(label Bookmark)
	Get(id string) (*Bookmark, bool)
	Find(term string) ([]*Bookmark, bool)
}

///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var _ BookmarkRepository = &memoryRepository{}

func NewMemoryRepository(labels []Bookmark) BookmarkRepository {
	r := &memoryRepository{}

	for _, label := range labels {
		r.Add(label)
	}

	return r
}

type memoryRepository struct {
	idMap   map[ider.Id]*Bookmark
	nameMap map[string]*Bookmark
}

func (this *memoryRepository) Add(label Bookmark) {
	if _, ok := this.idMap[label.Id]; !ok {
		this.idMap[label.Id] = &label
	}

	if _, ok := this.nameMap[label.Title]; !ok {
		this.nameMap[label.Title] = &label
	}
}

func (this memoryRepository) Find(term string) ([]*Bookmark, bool) {
	var found []*Bookmark

	if id, err := ider.Parse(term); err == nil {
		if label, ok := this.idMap[id]; ok {
			found = append(found, label)
		}
	}

	if label, ok := this.nameMap[term]; ok {
		found = append(found, label)
	}

	if len(found) <= 0 {
		for _, label := range this.idMap {
			if strings.Contains(term, label.Title) {
				found = append(found, label)
			}
		}
	}

	return found, len(found) > 0
}

func (this memoryRepository) Get(id string) (*Bookmark, bool) {
	parsed, err := ider.Parse(id)
	if err != nil {
		return nil, false
	}

	label, ok := this.idMap[parsed]
	return label, ok
}
