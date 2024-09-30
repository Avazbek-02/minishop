package main

import (
	"github.com/minishop/internal/logger"
	"github.com/minishop/internal/storage/postgres"
)

func main() {
	log, err := logger.NewLogger()
	if err != nil {
		return
	}
	_, err = postgres.ConnectPostgres()
	if err != nil {
		log.Error("Error connecting to database on main")
		return
	}
	log.Info("Successfully connected to database on main")
}
