package Rental

import (
	"RentalService/RentalStrategy"
	"RentalService/domain/Booking"
	"RentalService/domain/Branch"
	"RentalService/domain/Rate"
	"RentalService/domain/Vehicle"
)

type Service struct {
	BranchRepo     Branch.Repo
	VehicleRepo    Vehicle.Repo
	RateRepo       Rate.Repo
	BookingRepo    Booking.Repo
	RentalStrategy RentalStrategy.RentalStrategy
}

/*
addBranch(branchName)
This will add a new branch for your Service.

allocatePrice(branchName, vehicleType, price);
This will be used to define price per vehicle type per branch

addVehicle(vehicleId, vehicleType, branchName);
This will add a new vehicle of the given vehicle type in a given existing branch.

bookVehicle(vehicleType, startTime, endTime)
This will be used to rent a vehicle for the given vehicle type for a given time slot defined by Start time and end time. You can expect the start time and endTime to be in hourly slots only.

*/
func (s *Service) AddBranch(branchName string) {
	s.BranchRepo.AddBranch(branchName)
}
func (s *Service) AllocatePrice(branchName string, vehicleType string, price int) {
	s.RateRepo.AllocatePrice(branchName, vehicleType, price)
}
func (s *Service) AddVehicle(vehicleId int, vehicleType string, branchName string) {
	s.VehicleRepo.AddVehicle(vehicleId, vehicleType, branchName)
}

func (s *Service) BookVehicle(vType string, branchName string, startTime int, endTime int) {
	booked := false
	for _, v := range s.VehicleRepo.GetVehicles(branchName) {
		if v.Model == vType {
			for _, booking := range s.BookingRepo.GetBookings() {
				if (startTime <= booking.StartTime && endTime <= booking.EndTime) || (startTime >= booking.EndTime) {
					s.BookingRepo.AddBooking(startTime, endTime, v.Id)
					booked = true
					break
				}
			}
		}
		if booked {
			break
		}
	}

}
