package book_keeping

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/idiomatic/label"
)

// Assets = Liabilities
// Assets = Liabilities + Capital + Revenue + Gains - Expenses - Dividends - Losses
// 	=> Assets + Expenses + Dividends + Losses = Liabilities + Capital + Revenue + Gains
// DEALER
// 		- Debit-normal accounts: Dividends, Expenses, Assets
//		- Credit-normal accounts: Liabilities, Equity, Revenue
// https://en.wikipedia.org/wiki/Debits_and_credits

type TransactionModel struct {
	Id                ider.Id            `json:"id"`
	Name              string             `json:"name"`
	Description       string             `json:"description"`
	Source            AccountModel       `json:"source"`
	Destination       AccountModel       `json:"destination"`
	SourceAmount      float64            `json:"source-amount"`
	DestinationAmount float64            `json:"destination-amount"`
	Vendor            VendorModel        `json:"vendor"`
	Client            ClientModel        `json:"client"`
	Labels            []label.LabelModel `json:"labels"`
}
