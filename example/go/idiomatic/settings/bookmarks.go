package settings

import "github.com/boundedinfinity/schema/idiomatic/id"

type BookmarkManager struct {
	Id          id.Id               `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Directories []BookmarkDirectory `json:"directories,omitempty"`
}

type BookmarkDirectory struct {
	Id          id.Id               `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Directories []BookmarkDirectory `json:"directories,omitempty"`
	Bookmarks   []Bookmark          `json:"bookmarks,omitempty"`
}

type Bookmark struct {
	Id          id.Id  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Url         string `json:"url,omitempty"`
}
