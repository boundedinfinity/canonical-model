package measurement

// https://www.legislation.gov.uk/uksi/1994/2867/schedule/made
// https://en.wikipedia.org/wiki/Imperial_units

// # Square mile	=	640 acres.
// # Acre	        =	4840 square yards.
// # Rood	        =	1210 square yards.
// # Square yard	=	a superficial area equal to that of a square each side of which measures one yard.
// # Square foot	=	1/9 square yard.
// # Square inch	=	1/144; square foot.

var ImperialAreas = imperialAreas{
	factor: map[AreaUnit]float64{
		AreaUnits.SquareFoot: 1.0 / 144.0,
		AreaUnits.SquareFoot: 1.0,
		AreaUnits.SquareYard: 9.0,
		AreaUnits.Rood:       9.0 * 1210.0,
		AreaUnits.Acre:       9.0 * 4840.0,
		AreaUnits.SquareMile: 9.0 * 4840.0 * 640.0,
	},
	abbr: map[AreaUnit][]string{
		AreaUnits.Acre:       {AreaUnits.Acre.String()},
		AreaUnits.Rood:       {AreaUnits.Rood.String()},
		AreaUnits.SquareFoot: {AreaUnits.SquareFoot.String()},
		AreaUnits.SquareInch: {AreaUnits.SquareInch.String()},
		AreaUnits.SquareYard: {AreaUnits.SquareYard.String()},
		AreaUnits.SquareMile: {AreaUnits.SquareMile.String()},
	},
	system: map[AreaUnit]MeasurementSystem{
		AreaUnits.Acre:       MeasurementSystems.Imperial,
		AreaUnits.Rood:       MeasurementSystems.Imperial,
		AreaUnits.SquareFoot: MeasurementSystems.Imperial,
		AreaUnits.SquareInch: MeasurementSystems.Imperial,
		AreaUnits.SquareYard: MeasurementSystems.Imperial,
		AreaUnits.SquareMile: MeasurementSystems.Imperial,
	},
}

type imperialAreas struct {
	factor map[AreaUnit]float64
	abbr   map[AreaUnit][]string
	system map[AreaUnit]MeasurementSystem
}

type Area interface {
	Size()
}

type SquareArea[T ~int] struct {
	Side Length[T] `json:"size,omitempty"`
}

func (t SquareArea[T]) Size() {

}

var _ Area = &SquareArea[int]{}

type RectangleArea[T ~int] struct {
	Side1 Length[T] `json:"side-1,omitempty"`
	Side2 Length[T] `json:"side-2,omitempty"`
	Unit  AreaUnit  `json:"unit,omitempty"`
}
