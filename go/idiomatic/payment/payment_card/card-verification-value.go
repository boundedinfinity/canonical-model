package payment_card

// https://en.wikipedia.org/wiki/Card_security_code

type CardVerificationValue string

func (this CardVerificationValue) Validate() error {

	return nil
}
