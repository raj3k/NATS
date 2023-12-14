package main

import (
	"NATS/server"
)

func main() {
	cfg := server.Config{
		ListenAddr: "0.0.0.0:4222",
	}
	s := server.NewServer(&cfg)
	s.Run()
}
