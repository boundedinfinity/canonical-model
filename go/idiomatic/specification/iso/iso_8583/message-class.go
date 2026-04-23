package iso_8583

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

type MessageClass struct {
	Meaning string
	Code    string
}

var MessageClasses = messageClasss{}

type messageClasss struct {
	all []MessageClass
}

func (this messageClasss) Find(code string) (MessageClass, bool) {
	return slicer.FindFn(
		func(_ int, item MessageClass) bool { return item.Code == code },
		this.all...,
	)
}
