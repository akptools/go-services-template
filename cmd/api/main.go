package main

import (
	"github.com/akptools/go-services-template/internal/server"
	"github.com/joho/godotenv"
)

func init() {
	// load the env variables into the app
	godotenv.Load()
}

func main() {
	// creates a new server and runs it
	s := server.NewServer(":8443", nil)
	s.Run()
}
