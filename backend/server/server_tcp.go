package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

type CallbackFunc func(command string, args []string) <-chan string

type ServerTCP struct {
	Port    string
	Handler CallbackFunc
}

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
			_, err := conn.Write([]byte("ERROR: Invalid command format\n"))

			if err != nil {
				log.Println("Error writing to connection:", err)
				return
			}

			continue
		}

		command := strings.ToUpper(parts[0])

		args := parts[1:]

		responseChan := callback(command, args)

		response := <-responseChan

		_, err = conn.Write([]byte(response))

		if err != nil {
			log.Println("Error writing to connection:", err)
			return
		}
	}
}

func TCPServer(port string, callback CallbackFunc) {

	raw_port := fmt.Sprintf(":%s", port)

	listener, err := net.Listen("tcp", raw_port)

	if err != nil {
		log.Fatalf("Error starting server: %v\n", err)
	}

	defer listener.Close()

	log.Println("server is running on port ...", raw_port)

	for {

		conn, err := listener.Accept()

		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn, callback)
	}
}
