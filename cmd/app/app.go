package main

import (
	"fmt"
	"net"
	"os"

	grpc_ctrl "github.com/place-chat/account-service/internal/controller/grpc"
	"github.com/place-chat/account-service/internal/repo"
	"github.com/place-chat/account-service/internal/service"
	"github.com/place-chat/account-service/submodules/protoc/go/account_service"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println(os.Environ())

	Repo := repo.NewAdapter(&repo.DBConfig{
		Database: "place-chat",
		User:     "app",
		Password: "12345678",
		Host:     "localhost",
		Port:     5432,
	})

	Service := service.NewAdapter(Repo)

	GRPCCtrl := grpc_ctrl.NewAdatper(Service)

	listener, err := net.Listen("tcp", "9827")
	if err != nil {
		panic(err)
	}

	GRPCServ := grpc.NewServer()
	account_service.RegisterAccountServiceServer(GRPCServ, GRPCCtrl)

	if err := GRPCServ.Serve(listener); err != nil {
		panic(err)
	}
}
