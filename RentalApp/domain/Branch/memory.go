package Branch

type InMemoryRepo struct {
	branches []Branch
}

func (r *InMemoryRepo) AddBranch(name string) {
	r.branches = append(r.branches, NewBranch(name))
}
