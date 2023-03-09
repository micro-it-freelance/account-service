package service

import (
	"context"

	"github.com/place-chat/account-service/internal/repo"
)

func (a *Adapter) AccountCreate(ctx context.Context, args repo.AccountCreateArgs) (repo.AccountModel, error) {
	return a.Repo.AccountCreate(ctx, &args)
}
