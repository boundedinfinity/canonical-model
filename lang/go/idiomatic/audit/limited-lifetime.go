package audit

import "github.com/boundedinfinity/rfc3339date"

type LimitedLifetime struct {
	StartDate rfc3339date.Rfc3339Date `json:"start-date,omitempty"`
	EndDate   rfc3339date.Rfc3339Date `json:"end-date,omitempty"`
}

func (t *LimitedLifetime) MarkStart() {
	t.StartDate = rfc3339date.DateNow()
}

func (t *LimitedLifetime) MarkEnd() {
	t.EndDate = rfc3339date.DateNow()
}
