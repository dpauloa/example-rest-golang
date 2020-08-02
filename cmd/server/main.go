package main

import (
	"os"

	"dpauloa/example-rest-golang/database"
	"dpauloa/example-rest-golang/database/postgres"
	"dpauloa/example-rest-golang/domain/usecase"
	"dpauloa/example-rest-golang/log"
	"dpauloa/example-rest-golang/settings"
	"dpauloa/example-rest-golang/transport/http"
)

func main() {
	logger := log.NewStdout()

	config, errs := settings.LoadConfig()
	if errs != nil {
		settings.LogErrsAndExit(logger, errs)
	}

	db, err := database.ConnectionDatabase(config.DatabaseURL)
	if err != nil {
		logger.Critical("unable to create database: %v", err)
		os.Exit(1)
	}

	phoneBookRepo := postgres.NewPhoneBookRepo(db)
	phoneBookUC := usecase.NewCreatePhoneBook(phoneBookRepo)
	r := http.NewRouter(phoneBookUC)

	port := config.Port
	logger.Info("Running at port %s...", port)
	http.ListenAndServe(":"+port, r)
}
