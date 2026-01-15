package iso_8583

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

type MessageFunction struct {
	Meaning string
	Code    string
}

var MessageFunctions = messageFunctions{}

type messageFunctions struct {
	all []MessageFunction
}

func (this messageFunctions) Find(code string) (MessageFunction, bool) {
	return slicer.FindFn(
		func(_ int, item MessageFunction) bool { return item.Code == code },
		this.all...,
	)
}
