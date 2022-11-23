package vehicle

import (
	branch2 "branchRental/branch"
	"branchRental/domain"
)

type Vehicle struct {
	id     string
	vType  domain.VehicleType
	branch *branch2.Branch
}

func (v Vehicle) GetId() string {
	return v.id
}

func (v Vehicle) GetVType() domain.VehicleType {
	return v.vType
}

func (v Vehicle) Branch() *branch2.Branch {
	return v.branch
}

func NewVehicle(id string, VType domain.VehicleType, branch *branch2.Branch) *Vehicle {
	return &Vehicle{id: id, vType: VType, branch: branch}
}
