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

	client := NewClient()

	// io.Copy(c, c)

	buffer := make([]byte, 1024)

	for {
		n, err := c.Read(buffer)
		if err != nil {
			log.Println("Failed to read data.", err)
			return
		}

		if string(buffer[:n-2]) == "PING" {
			client.parse(buffer[:n])

			pong := []byte("PONG\r\n")
			_, err := c.Write(pong)
			if err != nil {
				log.Println("Failed to send PONG response.", err)
				return
			}
			continue
		}

		// fmt.Printf("Received: %s\n", buffer[:n])

		data := buffer[:n]
		_, err = c.Write(data)
		if err != nil {
			log.Println("Failed to send data.", err)
			return
		}
	}

}
