package main

import (
	internal "github.com/njavilas2015/godis/internal"
	server "github.com/njavilas2015/godis/server"
)

func main() {

	go internal.Hs.Process()

	server.TCPServer("6379", internal.Handler)
}
