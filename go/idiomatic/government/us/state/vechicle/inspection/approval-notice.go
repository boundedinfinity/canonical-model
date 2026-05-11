package vehicle_inspection

import (
	"github.com/boundedinfinity/canonical-model/go/idiomatic/business"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/measurement/time"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/payment"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/person"
	"github.com/boundedinfinity/canonical-model/go/idiomatic/vehicle"
)

// ID / OD Verification

type ApprovalNotice struct {
	Vehicle                 vehicle.Vehicle   `json:"vehicle,omitempty"`
	Person                  person.Person     `json:"person,omitempty"`
	PermitNumber            string            `json:"permit-number,omitempty"`
	Inspecter               person.Person     `json:"inspector,omitempty"`
	InspectionStation       business.Business `json:"inspection-station,omitempty"`
	InspectionStationNumber string            `json:"inspection-station-number,omitempty"`
	InspectionDate          time.Date         `json:"inspection-date,omitempty"`
	ControlNumber           string            `json:"control-number,omitempty"`
	OdometerReading         int               `json:"odometer-reading,omitempty"`
	PaidWith                payment.Payment   `json:"paid-with,omitempty"`
	BrakeInspected          struct {
		FrontLeft  bool `json:"front-left,omitempty"`
		FrontRight bool `json:"front-right,omitempty"`
		RearLeft   bool `json:"rear-left,omitempty"`
		RearRight  bool `json:"rear-right,omitempty"`
	} `json:"brake-inspected,omitempty"`
}
