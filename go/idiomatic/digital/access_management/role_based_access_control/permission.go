package role_based_access_control

import "github.com/boundedinfinity/canonical_model/idiomatic/ider"

type PermissionModel struct {
	ID          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
}
