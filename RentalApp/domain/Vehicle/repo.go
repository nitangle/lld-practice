package Vehicle

type Repo interface {
	AddVehicle(id int, model string, branch string)
	GetVehicles(branchName string) []Vehicle
}
