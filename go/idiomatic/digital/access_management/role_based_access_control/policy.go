package role_based_access_control

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/digital/access_management/resource"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type PolicyModel struct {
	ID          ider.Id             `json:"id"`
	Name        string              `json:"name"`
	Description string              `json:"description,omitempty"`
	Resources   []resource.Resource `json:"resources,omitempty"`
	Roles       []RoleModel         `json:"roles,omitempty"`
	Permissions []PermissionModel   `json:"permissions,omitempty"`
}
