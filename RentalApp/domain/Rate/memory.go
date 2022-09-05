package Rate

type InMemoryRepo struct {
	rates Rates
}

func (r *InMemoryRepo) GetBranchRates(branchName string) Rates {
	rates := make([]Rate, 0)
	for _, rate := range r.rates {
		if rate.Branch == branchName {
			rates = append(rates, rate)
		}
	}
	return rates
}

func (r *InMemoryRepo) GetRates() Rates {
	return r.rates
}

func (r *InMemoryRepo) AllocatePrice(branchName string, vehicleType string, price int) error {
	newRate := Rate{Branch: branchName, VehicleType: vehicleType, Price: price}
	for idx, rate := range r.rates {
		if rate.VehicleType == vehicleType && rate.Branch == branchName {
			r.rates[idx].Price = price
		}
	}
	r.rates = append(r.rates, newRate)
	return nil
}
