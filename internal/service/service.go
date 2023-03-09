package service

import "github.com/place-chat/account-service/internal/repo"

type Adapter struct {
	Repo *repo.Adapter
}

func NewAdapter(Repo *repo.Adapter) *Adapter {
	return &Adapter{
		Repo: Repo,
	}
}
