package rentalStrategy

import (
	"branchRental/rentalPrice"
	"branchRental/vehicle"
)

type RentalStrategy interface {
	SelectVehicle(rentalPriceRepo *rentalPrice.Repo, availableVehicles *[]vehicle.Vehicle) *vehicle.Vehicle
}
