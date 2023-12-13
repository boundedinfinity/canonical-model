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
	Amount currency.Amount         `json:"amount,omitempty"`
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
	MinimumPaymentDue           currency.Amount             `json:"minimum-payment-due,omitempty"`
	Capture                     digital_document.DocumentV1 `json:"capture,omitempty"`
	OpeningBalance              currency.Amount             `json:"opening-balance,omitempty"`
	ClosingBalance              currency.Amount             `json:"closing-balance,omitempty"`
	Purchases                   currency.Amount             `json:"purchases,omitempty"`
	CashAdvances                currency.Amount             `json:"cash-advances,omitempty"`
	BanlanceTransfers           currency.Amount             `json:"balance-transfers,omitempty"`
	FeesCharged                 currency.Amount             `json:"fees-charged,omitempty"`
	InterestCharged             currency.Amount             `json:"interest-charged,omitempty"`
	AvailableCredit             currency.Amount             `json:"available-credit,omitempty"`
	CashAccessLine              currency.Amount             `json:"cash-access-line,omitempty"`
	AvailableForCash            currency.Amount             `json:"available-for-cash,omitempty"`
	PastDueAmount               currency.Amount             `json:"past-due-amount,omitempty"`
}
