package time

import "github.com/boundedinfinity/canonical_model/go/idiomatic/ider"

// https://en.wikipedia.org/wiki/Time_zone

type TimeZone struct {
	Id          ider.Id `json:"id"`
	Name        string  `json:"name"`
	Abreviation string  `json:"abbreviation"`
	UtfOffset   string  `json:"utf-offset"`
}

type timeZones struct {
	Unknown                  TimeZone
	UniversalCoordinatedTime TimeZone
	EasternStandardTime      TimeZone
	CentralStandardTime      TimeZone
	MountainStandardTime     TimeZone
	PacificStandardTime      TimeZone
	AtlanticStandardTime     TimeZone
	IndianStandardTime       TimeZone
	GreenwichMeanTime        TimeZone
}

func (this timeZones) All() []TimeZone {
	return []TimeZone{
		this.Unknown,
		this.UniversalCoordinatedTime,
		this.EasternStandardTime,
		this.CentralStandardTime,
		this.MountainStandardTime,
		this.PacificStandardTime,
		this.AtlanticStandardTime,
		this.IndianStandardTime,
		this.GreenwichMeanTime,
	}
}

func (this timeZones) Get(id ider.Id) (TimeZone, bool) {
	var found TimeZone
	var ok bool

	for _, timeZone := range this.All() {
		if timeZone.Id == id {
			found = timeZone
			ok = true
			break
		}
	}

	return found, ok
}

func (this timeZones) Find(term string) ([]TimeZone, bool) {
	var found []TimeZone
	var ok bool

	for _, timeZone := range this.All() {
		if timeZone.Name == term || timeZone.Abreviation == term || timeZone.UtfOffset == term {
			found = append(found, timeZone)
			ok = true
		}
	}

	return found, ok
}

var TimeZones = timeZones{
	Unknown: TimeZone{},
	UniversalCoordinatedTime: TimeZone{
		Name:        "Universal Coordinated Time",
		Abreviation: "UTC",
		UtfOffset:   "+00:00",
	},
	EasternStandardTime: TimeZone{
		Name:        "Eastern Standard Time",
		Abreviation: "EST",
		UtfOffset:   "-05:00",
	},
	CentralStandardTime: TimeZone{
		Name:        "Central Standard Time",
		Abreviation: "CST",
		UtfOffset:   "-06:00",
	},
	MountainStandardTime: TimeZone{
		Name:        "Mountain Standard Time",
		Abreviation: "MST",
		UtfOffset:   "-07:00",
	},
	PacificStandardTime: TimeZone{
		Name:        "Pacific Standard Time",
		Abreviation: "PST",
		UtfOffset:   "-08:00",
	},
	AtlanticStandardTime: TimeZone{
		Name:        "Atlantic Standard Time",
		Abreviation: "AST",
		UtfOffset:   "-04:00",
	},
	IndianStandardTime: TimeZone{
		Name:        "Indian Standard Time",
		Abreviation: "IST",
		UtfOffset:   "+05:30",
	},
	GreenwichMeanTime: TimeZone{
		Name:        "Greenwich Mean Time",
		Abreviation: "GMT",
		UtfOffset:   "+00:00",
	},
}
