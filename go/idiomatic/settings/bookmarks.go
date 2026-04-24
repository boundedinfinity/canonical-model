package settings

type BookmarkManager struct {
	Id          ider.Id             `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Directories []BookmarkDirectory `json:"directories,omitempty"`
}

type BookmarkDirectory struct {
	Id          ider.Id             `json:"id,omitempty"`
	Name        string              `json:"name,omitempty"`
	Description string              `json:"description,omitempty"`
	Directories []BookmarkDirectory `json:"directories,omitempty"`
	Bookmarks   []Bookmark          `json:"bookmarks,omitempty"`
}

type Bookmark struct {
	Id          ider.Id `json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Url         string  `json:"url,omitempty"`
}
