package insurance

import (
	"github.com/boundedinfinity/schema/idiomatic/business"
	"github.com/boundedinfinity/schema/idiomatic/id"
)

type InsuranceCoverage struct {
	Id    id.Id   `json:"id,omitempty"`
	Limit float32 `json:"limit,omitempty"`
}

type InsurancePolicy struct {
	Id              id.Id             `json:"id,omitempty"`
	Number          string            `json:"number,omitempty"`
	MonthlyPremium  float32           `json:"monthly-premium,omitempty"`
	Company         business.Business `json:"company,omitempty"`
	Deductible      id.Id             `json:"deductible,omitempty"`
	Home            InsuranceCoverage `json:"home,omitempty"`
	Things          InsuranceCoverage `json:"things,omitempty"`
	OtherStructures InsuranceCoverage `json:"other-structures,omitempty"`
	Accidents       InsuranceCoverage `json:"accidents,omitempty"`
	LivingExpenses  InsuranceCoverage `json:"living-expenses,omitempty"`
	MedicalExpenses InsuranceCoverage `json:"medical-expenses,omitempty"`
}
