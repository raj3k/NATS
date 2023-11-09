package server

import (
	"io"
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

	io.Copy(c, c)

}
