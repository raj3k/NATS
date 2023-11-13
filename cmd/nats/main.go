package main

import (
	"NATS/server"
)

func main() {
	cfg := server.Config{
		ListenAddr: "localhost:4222",
	}
	s := server.NewServer(&cfg)
	s.Run()
}
