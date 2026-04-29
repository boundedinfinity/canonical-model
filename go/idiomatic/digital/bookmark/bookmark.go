package bookmark

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/digital/address/uri"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/ider"
)

type Bookmark struct {
	Id           ider.Id   `json:"id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Address      uri.Url   `json:"address"`
	MigratedTo   *Bookmark `json:"migrated-to"`
	MigratedFrom *Bookmark `json:"migrated-from"`
	// Created      time.Date `json:"created"`
}
