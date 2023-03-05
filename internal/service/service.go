package service

import "github.com/place-chat/account-service/internal/repo"

type Adapter struct {
	repo *repo.Adapter
}

func NewAdapter(repo *repo.Adapter) *Adapter {
	return &Adapter{repo}
}
