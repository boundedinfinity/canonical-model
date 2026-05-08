package vehicle

type TransmissionType string

type transmissionTypes struct {
	Automatic TransmissionType
	Manual    TransmissionType
}

var TransmissionTypes = transmissionTypes{
	Automatic: "automatic",
	Manual:    "manual",
}
