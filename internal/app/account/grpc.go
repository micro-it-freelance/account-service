package account

import (
	"context"

	"github.com/micro-it-freelance/protoc/out/account_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type AccountGRPCHandler struct {
	service Service
	account_service.UnimplementedAccountServiceServer
}

func NewAccountGRPCHandler(s Service) account_service.AccountServiceServer {
	return &AccountGRPCHandler{
		service: s,
	}
}

func (h *AccountGRPCHandler) Create(ctx context.Context, in *account_service.CreateRequest) (*account_service.CreateReply, error) {
	account, err := h.service.Create(ctx, in.TelegramId, in.Username)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &account_service.CreateReply{
		Id: account.ID,
	}, nil
}

func (h *AccountGRPCHandler) GetByTelegramID(ctx context.Context, in *account_service.GetByTelegramIDRequest) (*account_service.GetByTelegramIDReply, error) {
	account, err := h.service.GetByTelegramID(ctx, in.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &account_service.GetByTelegramIDReply{
		Account: &account_service.Account{
			Id:         account.ID,
			TelegramId: account.TelegramID,
			Username:   account.Username,
			CreatedAt:  timestamppb.New(account.CreatedAt),
		},
	}, nil
}
