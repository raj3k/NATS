package main

import (
	"NATS/server"
)

const (
	Red int = iota
	Orange
	Yellow
	Green
	Blue
	Indigo
	Violet
)

func main() {
	server.Run()
}
