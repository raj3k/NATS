package server

import (
	"fmt"
	"log"
	"net"
	"sync"
)

type Server struct {
	*Config
	clients map[uint64]*client
	mu      sync.RWMutex
}

type Config struct {
	ListenAddr string
}

func NewServer(cfg *Config) *Server {
	return &Server{Config: cfg, clients: make(map[uint64]*client)}
}

func (s *Server) acceptConn(conn net.Conn) {
	client := &client{
		nc: conn,
	}

	s.mu.Lock()
	clientID := uint64(len(s.clients) + 1)
	s.clients[clientID] = client
	s.mu.Unlock()

	fmt.Printf("Client %d connected\n", clientID)

	s.handleClient(clientID)
}

func (s *Server) getClient(cid uint64) *client {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.clients[cid]
}

func (s *Server) Run() {
	addr := s.Config.ListenAddr
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalln(err)
	}
	defer l.Close()

	log.Println("Server is running on:", addr)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Failed to accept conn.", err)
			continue
		}
		go s.acceptConn(conn)
	}
}

func (s *Server) handleClient(clientID uint64) {

	client := s.getClient(clientID)

	defer func() {
		s.mu.Lock()
		delete(s.clients, clientID)
		s.mu.Unlock()

		client.nc.Close()
		fmt.Printf("Client %d disconnected\n", clientID)
	}()

	buffer := make([]byte, 1024)

	for {
		n, err := client.nc.Read(buffer)
		if err != nil {
			log.Println("Failed to read data.", err)
			return
		}

		err = client.parse(buffer[:n])
		if err != nil {
			log.Println("Failed to parse data.", err)
			return
		}
		log.Printf("Client: %d; Server processed command: %s", clientID, buffer[:n])
	}
}
