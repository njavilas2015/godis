package server

import (
	"bufio"
	"fmt"
	"net"
	"strings"

	queue "github.com/njavilas2015/godis/queue"
)

func Proxy(command string) string {

	parts := strings.Fields(command)

	if len(parts) < 1 {
		return "ERROR: empty command"
	}

	queue.NewCommandQueue().Enqueue(queue.CommandQueue{
		ID:        "1",
		Operation: "SET",
		Key:       "foo",
		Value:     "bar",
	})

	return queue.Start(parts[0], parts[1:])
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

		response := Proxy(command)

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
