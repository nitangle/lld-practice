package RentalStrategy

import (
	"RentalService/services/Rental"
)

type RentalStrategy interface {
	SelectVehicleType(RentalService Rental.Service, branchName string) string
}
