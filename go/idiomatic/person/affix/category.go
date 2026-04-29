package affix

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
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
	return slicer.Contains(stringer.ToLower(this.Name), names...)
}
