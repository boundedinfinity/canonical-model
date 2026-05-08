package vehicle

type DriveTrain string

var DriveTrains = driveTrains{
	Unknown:         "vehical.drive-train.unknown",
	FourWheelDrive:  "vehical.drive-train.4-wheel-drive",
	AllWheelDrive:   "vehical.drive-train.all-wheel-drive",
	FrontWheelDrive: "vehical.drive-train.front-wheel-drive",
	RearWheelDrive:  "vehical.drive-train.rear-wheel-drive",
}

type driveTrains struct {
	Unknown         DriveTrain
	FourWheelDrive  DriveTrain
	AllWheelDrive   DriveTrain
	FrontWheelDrive DriveTrain
	RearWheelDrive  DriveTrain
}
