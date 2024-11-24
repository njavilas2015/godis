package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type CallbackFunc func(command string, args []string) <-chan string

func handleConnection(conn net.Conn, callback CallbackFunc) {

	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {

		message, err := reader.ReadString('\n')

		if err != nil {
			log.Println("Client disconnected:", err)
			return
		}

		parts := strings.Fields(strings.TrimSpace(message))

		if len(parts) < 1 {
			conn.Write([]byte("ERROR: Invalid command format\n"))
			continue
		}

		command := strings.ToUpper(parts[0])
		args := parts[1:]

		responseChan := callback(command, args)

		response := <-responseChan

		conn.Write([]byte(response + "\n"))
	}
}

func TCPServer(port string, callback CallbackFunc) {

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", port))

	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}

	defer listener.Close()

	log.Println("Server is running on port 6379...")

	for {

		conn, err := listener.Accept()

		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn, callback)
	}
}
