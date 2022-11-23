package rental

import (
	"branchRental/vehicle"
)

type Rental struct {
	BranchName    string
	VehicleRented *vehicle.Vehicle
	StartTime     int64
	EndTime       int64
}

func NewRental(branchName string, vehicleRented *vehicle.Vehicle, startTime int64, endTime int64) *Rental {
	return &Rental{BranchName: branchName, VehicleRented: vehicleRented, StartTime: startTime, EndTime: endTime}
}
