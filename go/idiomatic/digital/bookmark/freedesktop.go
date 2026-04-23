package bookmark

import (
	"time"
)

// https://www.freedesktop.org/wiki/Specifications/desktop-bookmark-spec/
// https://pyxml.sourceforge.net/topics/xbel/docs/html/xbel.html
// https://www.freedesktop.org/wiki/Specifications/shared-mime-info-spec/

type FreeDesktopBookmark struct {
	Name      string
	Exec      string
	Count     int
	Modified  time.Time
	Timestamp time.Time
}

type FreeDesktopBookmarkIcon struct {
	Type string
	Href string
	Name string
}

type FreeDesktopBookmarkGroup struct {
	Name        string
	Description string
}
