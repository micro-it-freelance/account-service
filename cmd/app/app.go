// test
package main

import (
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/micro-it-freelance/account-service/internal/app/account"
	dbconn "github.com/micro-it-freelance/core/db-conn"
	grpc_core "github.com/micro-it-freelance/core/grpc"
	"github.com/micro-it-freelance/protoc/out/account_service"
	"google.golang.org/grpc"
)

func main() {
	// connect to database
	db := dbconn.NewDBConn()

	// add listener
	listener := grpc_core.NewGRPCListener()

	// create grpc server
	GRPCServer := grpc.NewServer()
	account_service.RegisterAccountServiceServer(GRPCServer,
		account.NewAccountGRPCHandler(
			account.NewAccountService(
				account.NewAccountRepository(db),
			),
		))

	// serve
	grpc_core.Serve(GRPCServer, listener)
}
