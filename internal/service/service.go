package service

import "github.com/micro-it-freelance/account-service/internal/repo"

type Adapter struct {
	Repo *repo.Adapter
}

func NewAdapter(Repo *repo.Adapter) *Adapter {
	return &Adapter{
		Repo: Repo,
	}
}
