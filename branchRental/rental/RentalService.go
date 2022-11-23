package rental

import (
	"branchRental/branch"
	"branchRental/domain"
	"branchRental/rentalPrice"
	"branchRental/rentalStrategy"
	"branchRental/utils"
	"branchRental/vehicle"
	"fmt"
)

type Service interface {
	AddBranch(branchName string)
	AllocatePrice(branchName string, vehicleType domain.VehicleType, price int)
	AddVehicle(vehicleId string, vehicleType domain.VehicleType, branchName string) error
	BookVehicle(vehicleType domain.VehicleType, startTime string, endTime string) error
}

type RentalService struct {
	vehicleRepo     *vehicle.Repo
	rentalPriceRepo *rentalPrice.Repo
	rentalRepo      *Repo
	branchRepo      *branch.Repo
	rentalStrategy  rentalStrategy.RentalStrategy
}

func NewRentalService(vehicleRepo *vehicle.Repo, rentalPriceRepo *rentalPrice.Repo, rentalRepo *Repo, branchRepo *branch.Repo, rentalStrategy rentalStrategy.RentalStrategy) *RentalService {
	return &RentalService{vehicleRepo: vehicleRepo, rentalPriceRepo: rentalPriceRepo, rentalRepo: rentalRepo, branchRepo: branchRepo, rentalStrategy: rentalStrategy}
}

func (r RentalService) AddBranch(branchName string) {
	r.branchRepo.AddBranch(branchName)
}

func (r RentalService) AllocatePrice(branchName string, vehicleType domain.VehicleType, price int) {
	r.rentalPriceRepo.AllocatePrice(branchName, vehicleType, price)
}

func (r RentalService) AddVehicle(vehicleId string, vehicleType domain.VehicleType, branchName string) error {
	br := r.branchRepo.Branch(branchName)
	if br == nil {
		fmt.Errorf("branch does not exist")
	}
	r.vehicleRepo.AddVehicle(vehicleId, vehicleType, br)
	return nil
}

func (r RentalService) BookVehicle(vehicleType domain.VehicleType, startTime string, endTime string) error {
	// convert string time to int
	startTimeEpoch, err := utils.ConvertToEpoch(startTime)
	if err != nil {
		return err
	}
	endEpochTime, err := utils.ConvertToEpoch(endTime)
	if err != nil {
		return err
	}

	// find all vehicles of vehicleType
	vehiclesOfGivenType := r.vehicleRepo.VehiclesOfType(vehicleType)
	if vehiclesOfGivenType == nil {
		return fmt.Errorf("no vehicles of type %s present", vehicleType)
	}
	availableVehicles := make([]vehicle.Vehicle, 0)
	for idx1, _ := range vehiclesOfGivenType {
		found := false
		for idx2, _ := range r.rentalRepo.Rentals {
			bookedStartTime := r.rentalRepo.Rentals[idx2].StartTime
			bookedEndTime := r.rentalRepo.Rentals[idx2].EndTime
			vehicleId := r.rentalRepo.Rentals[idx2].VehicleRented.GetId()
			if vehicleId == vehiclesOfGivenType[idx1].GetId() &&
				((bookedStartTime >= startTimeEpoch && bookedStartTime <= endEpochTime) || (bookedEndTime >= startTimeEpoch && bookedEndTime <= endEpochTime)) {
				found = true
				break
			}
		}
		if !found {
			availableVehicles = append(availableVehicles, vehiclesOfGivenType[idx1])
		}
	}
	if len(availableVehicles) == 0 {
		return fmt.Errorf("no available %s found", vehicleType)
	}
	selectedVehicle := r.rentalStrategy.SelectVehicle(r.rentalPriceRepo, &availableVehicles)
	if selectedVehicle == nil {
		return fmt.Errorf("no %s found", vehicleType)
	}

	fmt.Printf("%v from %s booked \n", selectedVehicle.GetId(), selectedVehicle.Branch())
	newRental := NewRental(selectedVehicle.Branch().Name(), selectedVehicle, startTimeEpoch, endEpochTime)
	r.rentalRepo.AddRental(newRental)
	return nil
}
