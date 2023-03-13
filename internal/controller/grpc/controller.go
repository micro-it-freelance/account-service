package grpc_ctrl

import (
	"github.com/micro-it-freelance/account-service/internal/service"
	"github.com/micro-it-freelance/protoc/out/account_service"
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
