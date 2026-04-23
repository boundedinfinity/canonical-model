package routing_transit_number

// https://en.wikipedia.org/wiki/Magnetic_ink_character_recognition

type E13BSymbol string

const (
	Unknown          E13BSymbol = "unknown"
	TransitDelimiter E13BSymbol = "⑆"
	OnUsDelimiter    E13BSymbol = "⑈"
	AmountDelimiter  E13BSymbol = "⑇"
	DashDelimiter    E13BSymbol = "⑉"
)
