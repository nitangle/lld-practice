package rentalStrategy

import (
	"branchRental/rentalPrice"
	"branchRental/vehicle"
	"math"
)

type LowestPrice struct {
}

func NewLowestPrice() *LowestPrice {
	return &LowestPrice{}
}

func (l LowestPrice) SelectVehicle(rentalPriceRepo *rentalPrice.Repo, availableVehicles *[]vehicle.Vehicle) *vehicle.Vehicle {
	minPrice := math.MaxInt
	var minPriceVehicle vehicle.Vehicle

	for _, availableVehicle := range *availableVehicles {
		for _, rentalPriceInfo := range rentalPriceRepo.RentalPrices {
			if rentalPriceInfo.VehicleType == availableVehicle.GetVType() && rentalPriceInfo.BranchName == availableVehicle.Branch().Name() && minPrice > rentalPriceInfo.Price {
				minPriceVehicle = availableVehicle
				minPrice = rentalPriceInfo.Price
			}
		}

	}
	return &minPriceVehicle
}
