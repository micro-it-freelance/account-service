// test
package main

import (
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/micro-it-freelance/account-service/internal/app/account"
	"github.com/micro-it-freelance/core/config"
	dbconn "github.com/micro-it-freelance/core/db-conn"
	grpccore "github.com/micro-it-freelance/core/grpc"
	"github.com/micro-it-freelance/protoc/out/account_service"
	"google.golang.org/grpc"
)

func main() {
	// connect to database
	db := dbconn.NewDBConn()

	// add listener
	listener := grpccore.NewGRPCListener()

	// create grpc server
	GRPCServer := grpc.NewServer()
	account_service.RegisterAccountServiceServer(GRPCServer,
		account.NewAccountGRPCHandler(
			account.NewAccountService(
				account.NewAccountRepository(db),
			),
		))

	// serve
	fmt.Printf("[account-service] Listen on :%d\n", config.GRPC.Port)
	if err := GRPCServer.Serve(listener); err != nil {
		panic(err)
	}
}
