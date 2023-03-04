package main

import (
	grpc_ctrl "github.com/place-chat/user-service/internal/controller/grpc"
	"github.com/place-chat/user-service/internal/repo"
	"github.com/place-chat/user-service/internal/service"
)

func main() {
	_repo := repo.NewAdapter(&repo.DBConfig{
		Database: "place-chat",
		User:     "app",
		Password: "12345678",
		Host:     "localhost",
		Port:     5432,
	})

	_service := service.NewAdapter(_repo)

	_grpc := grpc_ctrl.NewAdatper(_service)
}
