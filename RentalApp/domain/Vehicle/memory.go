package Vehicle

type InMemoryRepo struct {
	//branchwise vehicle info
	vehicles map[string][]Vehicle
}

func (r *InMemoryRepo) GetVehicles(branchName string) []Vehicle {
	return r.vehicles[branchName]
}

func (r *InMemoryRepo) AddVehicle(id int, model string, branch string) {
	newV := NewVehicle(id, model)
	r.vehicles[branch] = []Vehicle{newV}
}
