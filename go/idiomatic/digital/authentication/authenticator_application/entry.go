package authenticator_application

import (
	"github.com/boundedinfinity/canonical_model/go/idiomatic/digital/bookmark"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/phone"
	"github.com/boundedinfinity/canonical_model/go/idiomatic/specification/rfc/rfc3339date"
)

type Entry struct {
	Application Application                 `json:"application"`
	BackupCodes []BackupCode                `json:"backup-codes"`
	Phone       phone.PhoneModel            `json:"phone"`
	Bookmarks   []bookmark.Bookmark         `json:"bookmarks"`
	CreatedDate rfc3339date.Rfc3339DateTime `json:"created-date"`
}
