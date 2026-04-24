package access_control_list

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/model/digital/access_management/permission"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/model/digital/access_management/resource"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/model/digital/access_management/subject"
)

type AccessControlListModel struct {
	ID          ider.Id                `json:"id"`
	Name        string                 `json:"name"`
	Description string                 `json:"description,omitempty"`
	Resource    resource.Resource      `json:"resource"`
	Items       []AccessControlLisItem `json:"items,omitempty"`
}

type AccessControlLisItem struct {
	Subject     subject.Subject              `json:"subjects"`
	Permissions []permission.PermissionModel `json:"permissions"`
}
