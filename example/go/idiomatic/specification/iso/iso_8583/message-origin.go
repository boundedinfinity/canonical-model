package iso_8583

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

type MessageOrigin struct {
	Meaning string
	Code    string
}

var MessageOrigins = mssageOrigins{}

type mssageOrigins struct {
	all []MessageOrigin
}

func (this mssageOrigins) Find(code string) (MessageOrigin, bool) {
	return slicer.FindFn(
		func(_ int, item MessageOrigin) bool { return item.Code == code },
		this.all...,
	)
}
