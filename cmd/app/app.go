package main

import (
	"net"

	_ "github.com/micro-it-freelance/account-service/internal/config"
	grpc_ctrl "github.com/micro-it-freelance/account-service/internal/controller/grpc"
	"github.com/micro-it-freelance/account-service/internal/repo"
	"github.com/micro-it-freelance/account-service/internal/service"
	"github.com/micro-it-freelance/account-service/submodules/protoc/out/account_service"
	"google.golang.org/grpc"
)

func main() {
	newRepo := repo.NewAdapter()
	newService := service.NewAdapter(newRepo)
	newGRPCCtrl := grpc_ctrl.NewAdatper(newService)

	listener, err := net.Listen("tcp", "9827")
	if err != nil {
		panic(err)
	}

	newGRPCServ := grpc.NewServer()
	account_service.RegisterAccountServiceServer(newGRPCServ, newGRPCCtrl)

	if err := newGRPCServ.Serve(listener); err != nil {
		panic(err)
	}
}
