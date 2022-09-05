package Booking

type Bookings []Booking

type Repo interface {
	GetBookings() Bookings
	AddBooking(startTime int, endTime int, vehicleId int)
}
