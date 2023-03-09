package grpc_ctrl

import (
	"github.com/place-chat/account-service/internal/service"
	"github.com/place-chat/account-service/submodules/protoc/go/account_service"
)

type Adapter struct {
	Service *service.Adapter
	account_service.UnimplementedAccountServiceServer
}

func NewAdatper(service *service.Adapter) *Adapter {
	return &Adapter{
		Service: service,
	}
}
