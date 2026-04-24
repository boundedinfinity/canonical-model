package bookmark

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/digital/address/uri"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/model/specifications/rfc/rfc3339date"
)

type Bookmark struct {
	Id           ider.Id                     `json:"id"`
	Title        string                      `json:"title"`
	Description  string                      `json:"description"`
	Address      uri.Url                     `json:"address"`
	MigratedTo   *Bookmark                   `json:"migrated-to"`
	MigratedFrom *Bookmark                   `json:"migrated-from"`
	Created      rfc3339date.Rfc3339DateTime `json:"created"`
}
