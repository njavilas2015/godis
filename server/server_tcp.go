package server

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
	queue "github.com/njavilas2015/godis/queue"
)

type TCPServer struct {
	Addr  string
	Queue *queue.Queue
}

func NewTCPServer(addr string, q *queue.Queue) *TCPServer {
	return &TCPServer{
		Addr:  addr,
		Queue: q,
	}
}

func (s *TCPServer) handleConnection(conn net.Conn) {

	defer conn.Close()

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {

		request := scanner.Text()

		jobID := uuid.New().String()

		responseChan := make(chan string)

		job := &queue.Job{
			ID:       jobID,
			Request:  request,
			Response: responseChan,
		}

		s.Queue.Add(job)

		select {
		case response := <-responseChan:
			conn.Write([]byte(response + "\n"))
		case <-time.After(5 * time.Second): // Timeout opcional
			conn.Write([]byte("Error: timeout en la respuesta\n"))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error leyendo conexión: %v\n", err)
	}
}

func (s *TCPServer) Start() error {

	listener, err := net.Listen("tcp", s.Addr)

	if err != nil {
		return fmt.Errorf("error iniciando el servidor: %v", err)
	}

	defer listener.Close()

	fmt.Printf("Servidor escuchando en %s\n", s.Addr)

	/*for {

		conn, err := listener.Accept()

		if err != nil {
			fmt.Printf("Error aceptando conexión: %v\n", err)
			continue
		}

		go s.handleConnection(conn)
	}*/

	for {

	}
}
