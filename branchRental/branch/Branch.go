package branch

type Branch struct {
	name string
	city string
}

func (b *Branch) Name() string {
	return b.name
}

func NewBranch(branchName string) Branch {
	return Branch{
		name: branchName,
		city: "Delhi",
	}
}
