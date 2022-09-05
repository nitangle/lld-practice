package ride_sharing_app

type Ride struct {
	vehicleId   string
	source      Place
	dest        Place
	status      string
	seatsLeft   uint8
	offeredBy   uint64
	vehicleType string
	requestedBy uint64
	id          uint64
}

type RideRequest struct {
	vehicleId   string
	source      Place
	dest        Place
	seats       uint8
	offeredBy   uint64
	vehicleType string
	requestedBy uint64
}
