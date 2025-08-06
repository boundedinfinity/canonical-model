package insurance_test

import (
	"testing"

	"github.com/boundedinfinity/schema/idiomatic/id"
	"github.com/boundedinfinity/schema/idiomatic/insurance"
	"github.com/boundedinfinity/schema/idiomatic/vehicle"
	"github.com/stretchr/testify/assert"
)

func Test_AutoInsureancePolicy(t *testing.T) {
	p1 := insurance.AutoInsurancePolicy{
		Id:           id.Ids.MustParse("AAAAAAAA-BBBB-CCCC-DDDD-EEEEEEEEEEEE"),
		PolicyNumber: "999 9999-X99-99X",
		Class:        "XXXXXXXXXX",
		Vechicle: vehicle.Vehicle{
			Id:        id.Ids.MustParse("AAAAAAAA-BBBB-CCCC-DDDD-EEEEEEEEEEEE"),
			YearModel: 2018,
			Make:      "Mercedes",
			Model:     "GLA250",
			BodyStyle: vehicle.BodyStyles.SportWagon,
			Vin:       "XXXXXXXXXXXXXXXXX",
		},
		Coverages: []insurance.InsuranceCoverage{
			{
				Symbol:   "A",
				Desc:     "Liability Coverage",
				Premiums: 160.10,
				Limits: []insurance.InsureanceLimit{
					{
						Desc: "Bodily Injury Limits",
						Values: []insurance.InsureanceLimitValue{
							{
								Desc:  "Each Person",
								Limit: 250000,
							},
							{
								Desc:  "Each Accident",
								Limit: 500000,
							},
						},
					},
					{
						Desc: "Property Damage Limit",
						Values: []insurance.InsureanceLimitValue{
							{
								Desc:  "Each Accident",
								Limit: 250000,
							},
						},
					},
				},
			},
			{
				Symbol:     "C",
				Desc:       "Medical Payments Coverage",
				Deductible: 500,
			},
		},
	}

	assert.Equal(t, p1, true)
}
