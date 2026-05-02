package visibility

type Visibility string

type visibilities struct {
	Unknown Visibility
	Public  Visibility
	Private Visibility
	Blocked Visibility
}

var Visibilities = visibilities{
	Unknown: "calendar.visibility.unknown",
	Public:  "calendar.visibility.public",
	Private: "calendar.visibility.private",
	Blocked: "calendar.visibility.blocked",
}
