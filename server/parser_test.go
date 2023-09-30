package server

import (
	"testing"
)

func dummyClient() *client {
	return &client{}
}

func TestParsePing(t *testing.T) {
	c := dummyClient()
	if c.state != OP_START {
		t.Fatalf("Expected OP_START, got: %d\n", c.state)
	}
	ping := []byte("PING\r\n")
	err := c.parse(ping[:1])
	if err != nil || c.state != OP_P {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(ping[1:2])
	if err != nil || c.state != OP_PI {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(ping[2:3])
	if err != nil || c.state != OP_PIN {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(ping[3:4])
	if err != nil || c.state != OP_PING {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(ping[4:5])
	if err != nil || c.state != OP_PING {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(ping[5:6])
	if err != nil || c.state != OP_START {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(ping)
	if err != nil || c.state != OP_START {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
}

func TestParsePong(t *testing.T) {
	c := dummyClient()
	if c.state != OP_START {
		t.Fatalf("Expected OP_START, got: %d\n", c.state)
	}
	pong := []byte("PONG\r\n")
	err := c.parse(pong[:1])
	if err != nil || c.state != OP_P {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(pong[1:2])
	if err != nil || c.state != OP_PO {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(pong[2:3])
	if err != nil || c.state != OP_PON {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(pong[3:4])
	if err != nil || c.state != OP_PONG {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(pong[4:5])
	if err != nil || c.state != OP_PONG {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(pong[5:6])
	if err != nil || c.state != OP_START {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
	err = c.parse(pong)
	if err != nil || c.state != OP_START {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
}

func TestParseConnect(t *testing.T) {
	c := dummyClient()
	if c.state != OP_START {
		t.Fatalf("Expected OP_START, got: %v\n", c.state)
	}
	connect := []byte("CONNECT {}\r\n")
	err := c.parse(connect)
	if err != nil || c.state != OP_START {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}

	connectWithArg := []byte("CONNECT {\"start\":true}\r\n")
	err = c.parse(connectWithArg)
	if err != nil || c.state != OP_START {
		t.Fatalf("Unexpected: %d : %v\n", c.state, err)
	}
}
