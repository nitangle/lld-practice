package main

import (
	"branchRental/branch"
	"branchRental/domain"
	"branchRental/rental"
	"branchRental/rentalPrice"
	"branchRental/rentalStrategy"
	"branchRental/vehicle"
	"fmt"
)

func main() {
	branchRepo := branch.NewBranchRepo()
	rentalRepo := rental.NewRepo()
	rentalPriceRepo := rentalPrice.NewRentalPriceRepo()
	vehicleRepo := vehicle.NewVehicleRepo()
	st := rentalStrategy.NewLowestPrice()
	rs := rental.NewRentalService(vehicleRepo, rentalPriceRepo, rentalRepo, branchRepo, st)
	rs.AddBranch("Vasanth Vihar")
	rs.AddBranch("Cyber City")
	rs.AllocatePrice("Vasanth Vihar", domain.Sedan, 100)
	rs.AllocatePrice("Vasanth Vihar", domain.Hatchback, 80)
	rs.AllocatePrice("Cyber City", domain.Sedan, 200)
	rs.AllocatePrice("Cyber City", domain.Hatchback, 50)

	rs.AddVehicle("DL 01 MR 9310", domain.Sedan, "Vasanth Vihar")
	rs.AddVehicle("DL 01 MR 9311", domain.Sedan, "Cyber City")
	rs.AddVehicle("DL 01 MR 9312", domain.Hatchback, "Cyber City")
	err := rs.BookVehicle(domain.Sedan, "2020-02-29T10:00:00+00:00", "2020-02-29T13:00:00+00:00")
	if err != nil {
		fmt.Println(err)
	}
	err = rs.BookVehicle(domain.Sedan, "2020-02-29T14:00:00+00:00", "2020-02-29T15:00:00+00:00")
	if err != nil {
		fmt.Println(err)
	}
	err = rs.BookVehicle(domain.Sedan, "2020-02-29T14:00:00+00:00", "2020-02-29T15:00:00+00:00")
	if err != nil {
		fmt.Println(err)
	}
	err = rs.BookVehicle(domain.Sedan, "2020-02-29T14:00:00+00:00", "2020-02-29T15:00:00+00:00")
	if err != nil {
		fmt.Println(err)
	}
	err = rs.BookVehicle(domain.Hatchback, "2020-02-29T14:00:00+00:00", "2020-02-29T15:00:00+00:00")
	if err != nil {
		fmt.Println(err)
	}
}
