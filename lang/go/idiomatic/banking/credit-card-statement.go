package banking

import (
	"github.com/boundedinfinity/rfc3339date"
	"github.com/boundedinfinity/schema/idiomatic/currency"
	"github.com/boundedinfinity/schema/idiomatic/digital_document"
	"github.com/boundedinfinity/schema/idiomatic/finanical"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type CreditCardStatement struct {
	Id          id.Id                      `json:"id,omitempty"`
	Account     BankAccount                `json:"account,omitempty"`
	Summary     CreditCardStatementSummary `json:"summary,omitempty"`
	Payments    []CreditCardTransaction    `json:"payments,omitempty"`
	Purchases   []CreditCardTransaction    `json:"purchases,omitempty"`
	Redemptions []CreditCardTransaction    `json:"redemptions,omitempty"`
}

type CreditCardTransaction struct {
	Amount currency.CurrencyAmount `json:"amount,omitempty"`
	Date   rfc3339date.Rfc3339Date `json:"date,omitempty"`
	Memo   string                  `json:"memo,omitempty"`
}

type CreditCardStatementSummary struct {
	PurchaseInterestRate        finanical.InterestRate      `json:"purchase-interest-rate,omitempty"`
	CashAdvanceInterestRate     finanical.InterestRate      `json:"cash-advance-interest-rate,omitempty"`
	BalanceTransferInterestRate finanical.InterestRate      `json:"balance-transfer-interest-rate,omitempty"`
	BeginOnDate                 rfc3339date.Rfc3339Date     `json:"begin-on-date,omitempty"`
	EndOnDate                   rfc3339date.Rfc3339Date     `json:"end-on-date,omitempty"`
	StatementDate               rfc3339date.Rfc3339Date     `json:"statement-date,omitempty"`
	PaymentDate                 rfc3339date.Rfc3339Date     `json:"payment-date,omitempty"`
	MinimumPaymentDue           currency.CurrencyAmount     `json:"minimum-payment-due,omitempty"`
	Capture                     digital_document.DocumentV1 `json:"capture,omitempty"`
	OpeningBalance              currency.CurrencyAmount     `json:"opening-balance,omitempty"`
	ClosingBalance              currency.CurrencyAmount     `json:"closing-balance,omitempty"`
	Purchases                   currency.CurrencyAmount     `json:"purchases,omitempty"`
	CashAdvances                currency.CurrencyAmount     `json:"cash-advances,omitempty"`
	BanlanceTransfers           currency.CurrencyAmount     `json:"balance-transfers,omitempty"`
	FeesCharged                 currency.CurrencyAmount     `json:"fees-charged,omitempty"`
	InterestCharged             currency.CurrencyAmount     `json:"interest-charged,omitempty"`
	AvailableCredit             currency.CurrencyAmount     `json:"available-credit,omitempty"`
	CashAccessLine              currency.CurrencyAmount     `json:"cash-access-line,omitempty"`
	AvailableForCash            currency.CurrencyAmount     `json:"available-for-cash,omitempty"`
	PastDueAmount               currency.CurrencyAmount     `json:"past-due-amount,omitempty"`
}
