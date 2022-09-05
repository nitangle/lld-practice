package Vehicle

type Vehicle struct {
	Id    int
	Model string
}

func NewVehicle(id int, model string) Vehicle {
	return Vehicle{
		Id:    id,
		Model: model,
	}
}
