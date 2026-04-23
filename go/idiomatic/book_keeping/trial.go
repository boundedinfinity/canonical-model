package book_keeping

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/idiomatic/label"
)

type TrialModel struct {
	Id           ider.Id            `json:"id"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Transactions []TransactionModel `json:"transactions"`
	Labels       []label.LabelModel `json:"labels"`
	debitNormal  []TransactionModel
	creditNormal []TransactionModel
}

func (this TrialModel) Sort() {

	for _, transaction := range this.Transactions {
		switch transaction.Source.Kind {
		case Kinds.Debit:
			this.debitNormal = append(this.debitNormal, transaction)
		case Kinds.Credit:
			this.creditNormal = append(this.creditNormal, transaction)
		}
	}
}

func (this TrialModel) Calculate() {
	this.Sort()

	var debitNormalTotal float64
	for _, transaction := range this.debitNormal {
		debitNormalTotal += transaction.SourceAmount
	}

	var creditNormalTotal float64
	for _, transaction := range this.creditNormal {
		creditNormalTotal += transaction.SourceAmount
	}

	if debitNormalTotal != creditNormalTotal {
		panic("trial balance is not balanced")
	}
}
