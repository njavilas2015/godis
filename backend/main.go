package main

import (
	internal "github.com/njavilas2015/godis/internal"
	server "github.com/njavilas2015/godis/server"
)

func main() {
	servers := []server.ServerTCP{
		{Port: "6480", Handler: internal.HandlerHashStore},
		{Port: "6481", Handler: internal.HandlerKvStore},
		{Port: "6482", Handler: internal.HandlerListStore},
		{Port: "6483", Handler: internal.HandlerSetStorage},
	}

	for _, s := range servers {
		go server.TCPServer(s.Port, s.Handler)
	}

	go server.HTTPServer()

	select {}
}
