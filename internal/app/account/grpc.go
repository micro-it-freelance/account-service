package account

import (
	"context"

	"github.com/micro-it-freelance/protoc/out/account_service"
)

type AccountGRPCHandler struct {
	s *AccountService
	account_service.UnimplementedAccountServiceServer
}

func NewAccountGRPCHandler(s *AccountService) *AccountGRPCHandler {
	return &AccountGRPCHandler{
		s: s,
	}
}

func (h *AccountGRPCHandler) AccountCreate(context.Context, *account_service.AccountCreateRequest) (*account_service.AccountCreateReply, error) {
	return &account_service.AccountCreateReply{}, nil
}
