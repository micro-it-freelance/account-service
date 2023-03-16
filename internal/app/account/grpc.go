package account

import (
	"context"

	"github.com/micro-it-freelance/protoc/out/account_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountGRPCHandler struct {
	service *AccountService
	account_service.UnimplementedAccountServiceServer
}

func NewAccountGRPCHandler(s *AccountService) *AccountGRPCHandler {
	return &AccountGRPCHandler{
		service: s,
	}
}

func (h *AccountGRPCHandler) AccountCreate(ctx context.Context, in *account_service.AccountCreateRequest) (*account_service.AccountCreateReply, error) {
	id, err := h.service.Create(ctx, in.TelegramId, in.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &account_service.AccountCreateReply{
		Id: id,
	}, nil
}
