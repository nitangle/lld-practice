package rentalPrice

import "branchRental/domain"

type RentalPrice struct {
	Price       int
	BranchName  string
	VehicleType domain.VehicleType
}

func NewRentalPrice(price int, branchName string, vehicleType domain.VehicleType) RentalPrice {
	return RentalPrice{Price: price, BranchName: branchName, VehicleType: vehicleType}
}
