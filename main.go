package main

// package main

import (
	"log"

	"github.com/bauerc/GoBlockchain/server"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(server.Run())
}
