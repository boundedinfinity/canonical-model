package business

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/contact"
)

type DateRange struct {
	Start rfc3339date.Rfc3339Date
	End   rfc3339date.Rfc3339Date
}

type PaySlipStatement struct {
	Period              DateRange
	Person              contact.Contact
	Taxes               []PaySlipStatementLineItem
	TaxableWages        []PaySlipStatementLineItem
	BeforeTaxDeductions []PaySlipStatementLineItem
	AfterTaxDeductions  []PaySlipStatementLineItem
}

type PaySlipStatementLineItem struct {
	Description string
	Current     float32
	YearToDate  float32
}
