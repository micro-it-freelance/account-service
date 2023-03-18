package main

import (
	"fmt"
	"net"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/jmoiron/sqlx"
	"github.com/micro-it-freelance/account-service/internal/app/account"
	"github.com/micro-it-freelance/config"
	"github.com/micro-it-freelance/protoc/out/account_service"
	"google.golang.org/grpc"
)

func main() {
	

	// connect to database
	db, err := sqlx.Connect("pgx", fmt.Sprintf("dbname=%s user=%s password=%s host=%s port=%d sslmode=disable",
		config.DB.Name, config.DB.User, config.DB.Password, config.DB.Host, config.DB.Port))
	if err != nil {
		panic(err)
	}

	// add listener
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GRPC.Port))
	if err != nil {
		panic(err)
	}

	// create grpc server
	GRPCServer := grpc.NewServer()
	account_service.RegisterAccountServiceServer(GRPCServer,
		account.NewAccountGRPCHandler(
			account.NewAccountService(
				account.NewAccountRepository(db),
			),
		))

	// serve
	fmt.Printf("Listen on :%d\n", config.GRPC.Port)
	if err := GRPCServer.Serve(listener); err != nil {
		panic(err)
	}
}
