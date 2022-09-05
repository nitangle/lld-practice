package Rate

type Rates []Rate

type Repo interface {
	GetBranchRates(branchName string) Rates
	GetRates() Rates
	AllocatePrice(branchName string, vehicleType string, price int) error
}
