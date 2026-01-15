package iso_8583

import "github.com/boundedinfinity/go-commoner/idiomatic/slicer"

type MessageTypeIndicator struct {
	Meaning string
	Code    string
}

var MessageTypeIndicators = messageTypeIndicators{}

type messageTypeIndicators struct {
	all []MessageTypeIndicator
}

func (this messageTypeIndicators) Find(code string) (MessageTypeIndicator, bool) {
	return slicer.FindFn(
		func(_ int, item MessageTypeIndicator) bool { return item.Code == code },
		this.all...,
	)
}
