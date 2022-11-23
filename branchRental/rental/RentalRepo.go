package rental

type Repo struct {
	Rentals []Rental
}

func NewRepo() *Repo {
	return &Repo{}
}

func (r *Repo) GetRentals() []Rental {
	return r.Rentals
}

func (r *Repo) AddRental(rental *Rental) {
	r.Rentals = append(r.Rentals, *rental)
}
