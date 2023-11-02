package audit

import (
	"github.com/boundedinfinity/rfc3339date"
)

func New() Audit {
	a := Audit{
		CreatedAt: rfc3339date.DateTimeNow(),
	}
	a.MarkUpdate()
	return a
}

type Audit struct {
	CreatedAt rfc3339date.Rfc3339DateTime   `json:"created-at,omitempty"`
	ModifedOn []rfc3339date.Rfc3339DateTime `json:"modified-on,omitempty"`
}

func (t *Audit) MarkUpdate() {
	t.ModifedOn = append(t.ModifedOn, rfc3339date.DateTimeNow())
}
