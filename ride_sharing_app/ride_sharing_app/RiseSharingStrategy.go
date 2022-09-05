package ride_sharing_app

type IRideSelectStrategy interface {
	selectRide(rideDetails *RideRequest, service IRideService) *Ride
}

type MostVacant struct {
}

type PreferredVehicle struct {
}

func (m *PreferredVehicle) selectRide(rideDetails *RideRequest, service IRideService) *Ride {
	rides := service.getAllRides()
	for _, ride := range rides {
		if rideDetails.vehicleType == ride.vehicleType && rideDetails.source == ride.source && ride.dest == rideDetails.dest && ride.seatsLeft >= rideDetails.seats {
			return &ride
		}
	}
	return nil
}

func (m *MostVacant) selectRide(rideDetails *RideRequest, service IRideService) *Ride {
	rides := service.getAllRides()
	var mxSeatsLeft uint64 = 0
	var r *Ride
	for _, ride := range rides {
		if rideDetails.vehicleType == ride.vehicleType &&
			rideDetails.source == ride.source &&
			ride.dest == rideDetails.dest && ride.seatsLeft >= rideDetails.seats &&
			ride.status == "active" {
			if mxSeatsLeft < uint64(ride.seatsLeft) {
				mxSeatsLeft = uint64(ride.seatsLeft)
				r = &ride
			}
		}
	}

	return r
}
