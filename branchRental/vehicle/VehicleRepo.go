package vehicle

import (
	branch2 "branchRental/branch"
	"branchRental/domain"
)

type Repo struct {
	Vehicles []Vehicle
}

func NewVehicleRepo() *Repo {
	return &Repo{}
}

func (r *Repo) VehiclesOfType(vehicleType domain.VehicleType) []Vehicle {
	var vehicles []Vehicle
	for idx, _ := range r.Vehicles {
		if r.Vehicles[idx].vType == vehicleType {
			vehicles = append(vehicles, r.Vehicles[idx])
		}
	}
	return vehicles
}

func (r *Repo) AddVehicle(vehicleId string, vehicleType domain.VehicleType, branch *branch2.Branch) {
	newVehicle := NewVehicle(vehicleId, vehicleType, branch)
	r.Vehicles = append(r.Vehicles, *newVehicle)
}
