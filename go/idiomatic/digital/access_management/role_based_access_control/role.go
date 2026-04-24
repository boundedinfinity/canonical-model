package role_based_access_control

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/digital/access_management/subject"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
)

type RoleModel struct {
	ID          ider.Id           `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Subjects    []subject.Subject `json:"subjects,omitempty"`
}
