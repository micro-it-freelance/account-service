package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/micro-it-freelance/account-service/internal/app/account"
	core_db "github.com/micro-it-freelance/core/database"
	core_grpc "github.com/micro-it-freelance/core/grpc"
	"github.com/micro-it-freelance/protoc/out/account_service"
	"google.golang.org/grpc"
)

func main() {
	// connect to database
	db := core_db.NewDB()

	// add listener
	listener := core_grpc.NewGRPCListener()

	// create grpc server
	GRPCServer := grpc.NewServer()
	account_service.RegisterAccountServiceServer(GRPCServer,
		account.NewAccountGRPCHandler(
			account.NewAccountService(
				account.NewAccountRepository(db),
			),
		))

	// serve
	core_grpc.Serve(GRPCServer, listener)
}
