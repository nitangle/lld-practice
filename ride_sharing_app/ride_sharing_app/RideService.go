package ride_sharing_app

import (
	"errors"
	"fmt"
	"strings"
)

type IRideService interface {
	setSelectionStrategy(strategy IRideSelectStrategy)
	getAllRides() map[uint64]Ride

	OfferRide(rideDetails *RideRequest) error
	RequestRide(rideDetails *RideRequest, strategy string) error
	EndRide(rideId uint64)
}

type RideService struct {
	rides    map[uint64]Ride
	strategy IRideSelectStrategy
}

func (rs *RideService) setSelectionStrategy(strategy IRideSelectStrategy) {
	rs.strategy = strategy
}

func (rs *RideService) getAllRides() map[uint64]Ride {
	return rs.rides
}

func (rs *RideService) printStats() {
	for _, ride := range rs.rides {
		fmt.Printf("Ride Details:%+v\n", ride)
	}
}

func (rs *RideService) RequestRide(rideDetails *RideRequest, strategy string) error {
	if strategy == "Most Vacant" {
		rs.setSelectionStrategy(&MostVacant{})
	} else if strings.Contains(strategy, "Preferred") {
		vehicleType := strings.Split(strategy, "=")[1]
		rideDetails.vehicleType = vehicleType
		rs.setSelectionStrategy(&PreferredVehicle{})
	}
	r := rs.strategy.selectRide(rideDetails, rs)
	if r == nil {
		return errors.New("could not assign ride")
	}
	r.seatsLeft -= rideDetails.seats
	r.status = "active"
	return nil
}

func (rs *RideService) EndRide(rideId uint64) {
	if _, ok := rs.rides[rideId]; !ok {
		fmt.Println("ride does not exist")

	}
	r := rs.rides[rideId]
	r.status = "end"
	rs.rides[rideId] = r
}

func (rs *RideService) OfferRide(rideDetails *RideRequest) error {
	for _, ride := range rs.rides {
		if ride.offeredBy == rideDetails.offeredBy && ride.vehicleId == rideDetails.vehicleId {
			err := errors.New("Invalid request. Already offered a ride by this user")
			return err
		}
		rId := GenerateId()
		ride.id = rId
		rs.rides[rId] = ride

	}
	return nil
}
