package website

import (
	"net/url"

	"github.com/boundedinfinity/canonical_model/idiomatic/id"
)

type PortalWebSite struct {
	Id  id.Id   `json:"id,omitempty"`
	Url url.URL `json:"url,omitempty"`
}
