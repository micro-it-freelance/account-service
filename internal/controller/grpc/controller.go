package grpc_ctrl

import "github.com/place-chat/account-service/internal/service"

type Adapter struct {
	service *service.Adapter
}

func NewAdatper(service *service.Adapter) *Adapter {
	return &Adapter{service}
}