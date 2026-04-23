package loan

type Period string

type periods struct {
	Annual  Period
	Yearly  Period
	Mensual Period
	Monthly Period
}

var Periods = periods{
	Annual:  "annual",
	Yearly:  "yearly",
	Mensual: "mensual",
	Monthly: "monthly",
}
