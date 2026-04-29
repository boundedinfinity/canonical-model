package calendar

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/label"
)

type CalendarModel struct {
	Id          ider.Id        `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Labels      []label.Labels `json:"labels"`
}
