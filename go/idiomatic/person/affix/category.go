package affix

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/util/slicer"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/util/stringer"
)

type Category struct {
	Id          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}

func (this Category) String() string {
	return this.Name
}

func (this Category) Matches(names []string) bool {
	names = stringer.SliceToLower(names)
	return slicer.Contains(names, stringer.ToLower(this.Name))
}
