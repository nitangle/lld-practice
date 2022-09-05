package Branch

import (
	"RentalService/domain/Branch"
	"RentalService/services/Vehicle"
)

type Service struct {
	BranchRepo Branch.Repo
}

func NewService(BranchRepo Branch.Repo, VehicleService Vehicle.Service) Service {
	return Service{
		BranchRepo: BranchRepo,
	}
}
