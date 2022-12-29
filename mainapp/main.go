package main

// Simple API starter for AWS lambda function and AWS API Gateway using api key
// @author Olli <oliver@cyberlnk.dev>
// Example setup with mariadb and lambda router
// Embedded env file

import (
	"embed"
	"github.com/driftprogramming/godotenv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kedric/lambdarouter"
	"log"
	"mobile/internal/mariadb"
	"mobile/pk/search"
	"os"
)

//go:embed envs/*
var envs embed.FS

func main() {

	// get access to embedded env file
	err := godotenv.Load(envs, "envs/.env")
	if err != nil {
		log.Fatal("No env data")
	}

	// set up db
	dbError := mariadb.SetUpDatabase()
	if dbError != nil {
		os.Exit(1)
	}

	// start router.
	router := lambdarouter.New()

	// routes
	search.SetupRoutes(router)
	// run server
	router.Serve(":"+os.Getenv("SERVER_LISTEN_PORT"), nil)
}
