package iso_8583

type Message string

func (this Message) TypeIndicator() (MessageTypeIndicator, bool) {
	code := string(this[0])
	return MessageTypeIndicators.Find(code)
}

func (this Message) Class() (MessageClass, bool) {
	code := string(this[1])
	return MessageClasses.Find(code)
}

func (this Message) Function() (MessageFunction, bool) {
	code := string(this[2])
	return MessageFunctions.Find(code)
}

func (this Message) Origin() (MessageOrigin, bool) {
	code := string(this[3])
	return MessageOrigins.Find(code)
}
