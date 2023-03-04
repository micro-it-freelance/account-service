package main

import "github.com/iivkis/place-chat/user-service/internal/repo"

func main() {
	_repo := repo.NewRepo(&repo.DBConfig{
		Database: "place-chat",
		User:     "app",
		Password: "12345678",
		Host:     "localhost",
		Port:     5432,
	})
}
