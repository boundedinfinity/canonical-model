package dental

import (
	"github.com/boundedinfinity/canonical_model/idiomatic/currency"
	"github.com/boundedinfinity/canonical_model/idiomatic/ider"
	"github.com/boundedinfinity/canonical_model/idiomatic/person"
	"github.com/boundedinfinity/rfc3339date"
)

type PaymentSummary struct {
	Id                              ider.Id                 `json:"id"`
	TotalApprovedCharges            currency.Amount         `json:"total-approved-charges"`
	DentistPayment                  currency.Amount         `json:"dentist-payment"`
	InsurancePayment                currency.Amount         `json:"insurance-payment"`
	AppliedToDeductible             currency.Amount         `json:"applied-to-deductible"`
	DentistNonBillableAmount        currency.Amount         `json:"dentist-non-billable-amount"`
	OutOfPocketAmount               currency.Amount         `json:"out-of-pocket-amount"`
	ClaimNumber                     string                  `json:"claim-number"`
	PaymentDate                     rfc3339date.Rfc3339Date `json:"payment-date"`
	PayorCheckNumber                string                  `json:"payor-check-number"`
	DentistIdNumber                 string                  `json:"dentist-id-number"`
	DentistName                     string                  `json:"dentist-name"`
	ParState                        string                  `json:"par-state"`
	BenefitPeriodStartDate          rfc3339date.Rfc3339Date `json:"benefit-period-start-date"`
	BenefitPeriodEndDate            rfc3339date.Rfc3339Date `json:"benefit-period-end-date"`
	AnnualPlanMaximum               currency.Amount         `json:"annual-plan-maximum"`
	AnnualPlanUsedToDate            currency.Amount         `json:"annual-plan-used-to-date"`
	AnnualImplantMaximum            currency.Amount         `json:"annual-implant-maximum"`
	AnnualImplantUsedToDate         currency.Amount         `json:"annual-implant-used-to-date"`
	AnnualOrthodonticMaximum        currency.Amount         `json:"annual-orthodontic-maximum"`
	AnnualOrthodonticUsedToDate     currency.Amount         `json:"annual-orthodontic-used-to-date"`
	GroupId                         string                  `json:"group-id"`
	GroupName                       string                  `json:"group-name"`
	Member                          person.PersonModel      `json:"member"`
	Patient                         person.PersonModel      `json:"patient"`
	Relationship                    string                  `json:"relationship"`
	PlanType                        string                  `json:"plan-type"`
	ToothNoOrLetter                 string                  `json:"tooth-no-or-letter"`
	Surface                         string                  `json:"surface"`
	DateOfService                   rfc3339date.Rfc3339Date `json:"date-of-service"`
	SubmittedProcedureNumber        string                  `json:"submitted-procedure-number"`
	PaidProcedureNumber             string                  `json:"paid-procedure-number"`
	SubmittedAmount                 currency.Amount         `json:"submitted-amount"`
	ApprovedAmount                  currency.Amount         `json:"approved-amount"`
	AmountUsedForBenefitCalculation currency.Amount         `json:"amount-used-for-benefit-calculation"`
	DeductibleAmount                currency.Amount         `json:"deductible-amount"`
	CopayAmount                     currency.Amount         `json:"copay-amount"`
	InsurancePaymentAmount          currency.Amount         `json:"insurance-payment-amount"`
}
