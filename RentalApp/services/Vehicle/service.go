package Vehicle

import (
	"RentalService/domain/Vehicle"
)

type Service struct {
	VehicleRepo Vehicle.Repo
}

func NewService(VehicleRepo Vehicle.Repo) Service {
	return Service{
		VehicleRepo,
	}
}
