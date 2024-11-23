package main

import (
	"log"

	"github.com/njavilas2015/godis/server"
)

func main() {

	port := "6379"

	log.Printf("Welcome to Godis! starting in %v", port)

	err := server.Start(port)

	if err != nil {
		log.Fatalf("Error starting Godis :(")
	}
}
