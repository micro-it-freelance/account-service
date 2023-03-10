package grpc_ctrl

import (
	"context"

	"github.com/micro-it-freelance/account-service/internal/repo"
	"github.com/micro-it-freelance/account-service/submodules/protoc/go/account_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (a *Adapter) AccountCreate(ctx context.Context, in *account_service.AccountCreateRequest) (*account_service.AccountCreateReply, error) {
	account, err := a.Service.AccountCreate(ctx, repo.AccountCreateArgs{
		TelegramID: in.TelegramId,
		Username:   in.Username,
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &account_service.AccountCreateReply{
		Id: account.ID,
	}, nil
}
