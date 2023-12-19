package settings

import (
	"github.com/boundedinfinity/schema/idiomatic/contact"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type BrowserType struct {
	Id     id.Id           `json:"id,omitempty"`
	Name   string          `json:"name,omitempty"`
	Vendor contact.Contact `json:"vendor,omitempty"`
}

type BrowserProfile struct {
	Id              id.Id             `json:"id,omitempty"`
	Name            string            `json:"name,omitempty"`
	MachineId       id.MachineId      `json:"machine-id,omitempty"`
	Type            BrowserType       `json:"browser-type,omitempty"`
	Appearance      BrowserAppearance `json:"appearance,omitempty"`
	Security        BrowserSecurity   `json:"security,omitempty"`
	BookmarkManager BookmarkManager   `json:"bookmark-manager,omitempty"`
}

type BrowserAppearance struct {
	Theme    string `json:"theme,omitempty"`
	SideBar  bool   `json:"side-bar,omitempty"`
	InfoBar  bool   `json:"info-bar,omitempty"`
	UrlBar   bool   `json:"url-bar,omitempty"`
	FontSize string `json:"font-size,omitempty"`
}

type BrowserSecurity struct {
	SitesAllowedCookies           []string `json:"sites-allowed-cookies,omitempty"`
	SitesAllowedCookiesUtilClosed []string `json:"sites-allowed-cookies-until-closed,omitempty"`
	SitesNeverUseCookies          []string `json:"sites-never-use-cookies,omitempty"`
}
