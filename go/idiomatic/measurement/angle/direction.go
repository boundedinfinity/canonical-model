package angle

type Direction string

var Directions = directions{
	Clockwise:        "measurement.angle.direction.clockwise",
	CounterClockwise: "measurement.angle.direction.counter-clockwise",
}

type directions struct {
	Clockwise        Direction
	CounterClockwise Direction
}
