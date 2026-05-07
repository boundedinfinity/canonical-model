package grantor

type Kind string

var Kinds = kinds{
	IndividualPerson: "legal.grantor.individual-person",
	MultiplePeople:   "legal.grantor.multiple-people",
	Business:         "legal.grantor.business",
}

type kinds struct {
	IndividualPerson Kind
	MultiplePeople   Kind
	Business         Kind
}
