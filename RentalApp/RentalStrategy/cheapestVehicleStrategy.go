package RentalStrategy

import (
	"RentalService/services/Rental"
)

type CheapestVehicle struct {
}

func (c *CheapestVehicle) SelectVehicleType(RentalService Rental.Service, branchName string) string {
	mnPrice := 999999
	selectedVehicleType := ""
	for _, rate := range RentalService.RateRepo.GetBranchRates(branchName) {
		if rate.Price < mnPrice {
			mnPrice = rate.Price
			selectedVehicleType = rate.VehicleType
		}
	}
	return selectedVehicleType

}
