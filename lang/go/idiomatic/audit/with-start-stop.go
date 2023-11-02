package audit

import "github.com/boundedinfinity/rfc3339date"

type StartStop struct {
	StartDate rfc3339date.Rfc3339Date `json:"start-date,omitempty"`
	EndDate   rfc3339date.Rfc3339Date `json:"end-date,omitempty"`
}

func (t *StartStop) MarkStart() {
	t.StartDate = rfc3339date.DateNow()
}

func (t *StartStop) MarkEnd() {
	t.EndDate = rfc3339date.DateNow()
}
