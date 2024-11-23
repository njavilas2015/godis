package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	"github.com/njavilas2015/godis/storage"
)

var store *storage.Storage = storage.NewStorage()

func processCommand(command string) string {

	parts := strings.Fields(command)

	if len(parts) < 1 {
		return "ERROR: empty command"
	}

	switch strings.ToUpper(parts[0]) {

	case "SET":

		if len(parts) != 3 {
			return "ERROR: Correct use: SET <key> <value>"
		}

		store.Set(parts[1], parts[2])

		return "OK"

	case "GET":

		if len(parts) != 2 {
			return "ERROR: Correct use: GET <key>"
		}

		value, err := store.Get(parts[1])

		if !err {
			return "(nil)"
		}

		return value

	case "DEL":

		if len(parts) != 2 {
			return "ERROR: Correct use: DEL <key>"
		}

		store.Delete(parts[1])

		return "OK"

	default:
		return "ERROR: Unrecognized command"
	}
}

func handleConnection(socket net.Conn) {

	defer socket.Close()

	reader := bufio.NewReader(socket)

	for {

		raw, err := reader.ReadString('\n')

		if err != nil {

			fmt.Printf("Error reading command: %v\n", err)
		}

		command := strings.TrimSpace(raw)

		response := processCommand(command)

		socket.Write([]byte(response + "\n"))
	}
}

func Start(port string) error {

	listener, err := net.Listen("tcp", ":"+port)

	if err != nil {
		return fmt.Errorf("could not start the server %v", err)
	}

	defer listener.Close()

	for {

		socket, err := listener.Accept()

		if err != nil {
			fmt.Printf("Error accepting connection %v\n", err)
			continue
		}

		go handleConnection(socket)
	}
}
