package rentalPrice

import (
	"branchRental/domain"
)

type Repo struct {
	RentalPrices []RentalPrice
}

func NewRentalPriceRepo() *Repo {
	return &Repo{}
}

func (r *Repo) AllocatePrice(branchName string, vehicleType domain.VehicleType, price int) error {
	for idx, rentalPriceInfo := range r.RentalPrices {
		if rentalPriceInfo.BranchName == branchName && rentalPriceInfo.VehicleType == vehicleType {
			r.RentalPrices[idx].Price = price
			return nil
		}
	}
	newRentalPrice := NewRentalPrice(price, branchName, vehicleType)
	r.RentalPrices = append(r.RentalPrices, newRentalPrice)
	return nil
}
