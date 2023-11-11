package server

import (
	"log"
	"net"
)

func Run() {
	addr := "localhost:4222"
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
		go handleClient(conn)
	}
}

func handleClient(c net.Conn) {
	defer c.Close()

	client := NewClient(c)

	buffer := make([]byte, 1024)

	for {
		n, err := c.Read(buffer)
		if err != nil {
			log.Println("Failed to read data.", err)
			return
		}

		err = client.parse(buffer[:n])
		if err != nil {
			log.Println("Failed to parse data.", err)
			return
		}
	}
}
