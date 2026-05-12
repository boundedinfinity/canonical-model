package authenticator_application

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/digital/bookmark"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/phone"
)

type Entry struct {
	Application Application         `json:"application"`
	BackupCodes []BackupCode        `json:"backup-codes"`
	Phone       phone.Phone         `json:"phone"`
	Bookmarks   []bookmark.Bookmark `json:"bookmarks"`
	// CreatedDate time.Date           `json:"created-date"`
}
