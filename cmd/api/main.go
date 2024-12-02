package main

import (
	"log"
	"os"

	"github.com/akptools/go-services-template/internal/db"
	"github.com/akptools/go-services-template/internal/server"
	"github.com/akptools/go-services-template/internal/util"
	"github.com/joho/godotenv"
)

func init() {
	// load the env variables into the app
	err := godotenv.Load()
	if err != nil {
		log.Fatal("There was an error while loading env variables")
	}
}

func main() {
	db, err := db.NewDatabase(util.PostgresDB, os.Getenv("DB_URI"))
	if err != nil {
		log.Fatal("There was an error while connecting to the DB", "err:", err)
	}
	log.Println("Connected to the DB")
	// creates a new server and runs it
	s := server.NewServer(":8443", db)
	s.Run()
}
