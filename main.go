package main

import (
	"log"

	"github.com/njavilas2015/godis/queue"
	"github.com/njavilas2015/godis/server"
)

func main() {

	port := "6379"

	log.Printf("Welcome to Godis! starting in %v", port)

	q := queue.NewQueue(100)

	err := server.NewTCPServer(port, q)

	if err != nil {
		log.Fatalf("error starting Godis :( %v", err)
	}
}
