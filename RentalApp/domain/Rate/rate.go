package Rate

type Rate struct {
	VehicleType string
	Branch      string
	Price       int
}

func NewRate(vehicleType string,
	branch string,
	price int) Rate {
	return Rate{VehicleType: vehicleType, Branch: branch, Price: price}
}
