package Booking

type MemoryRepo struct {
	bookings Bookings
}

func (m *MemoryRepo) AddBooking(startTime int, endTime int, vehicleId int) {
	m.bookings = append(m.bookings, Booking{
		StartTime: startTime,
		EndTime:   endTime,
		VehicleId: vehicleId,
	})
	return
}

func (m *MemoryRepo) GetBookings() Bookings {
	return m.bookings
}
