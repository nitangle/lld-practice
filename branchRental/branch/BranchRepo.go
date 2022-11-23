package branch

import (
	"fmt"
)

type Repo struct {
	branches map[string]Branch
}

func (b *Repo) Branch(branchName string) *Branch {
	if val, ok := b.branches[branchName]; !ok {
		return nil
	} else {
		return &val
	}
}

func (b *Repo) AddBranch(branchName string) error {
	if _, ok := b.branches[branchName]; ok {
		return fmt.Errorf("branch already exists")
	}
	newBranch := NewBranch(branchName)
	b.branches[branchName] = newBranch
	return nil
}

func NewBranchRepo() *Repo {
	branches := make(map[string]Branch)
	return &Repo{branches: branches}
}
