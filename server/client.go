package server

import "fmt"

type client struct {
	parseState
}

func (c *client) processPub(arg []byte) error {
	//a := [2][]byte{}
	var args [][]byte
	start := -1
	for i, b := range arg {
		switch b {
		case ' ', '\t':
			if start >= 0 {
				args = append(args, arg[start:i])
				start = -1
			}
		default:
			if start < 0 {
				start = i
			}
		}
	}
	if start >= 0 {
		args = append(args, arg[start:])
	}
	c.pa.arg = arg
	c.pa.subject = args[0]
	c.pa.size = parseSize(args[1])
	return nil
}

func (c *client) processInboundMessage(msg []byte) {
	// TODO: To be implemented
	fmt.Println(string(msg))
}

func (c *client) processPing() {
	c.sendPong()
}

func (c *client) sendPong() {
	fmt.Println("PONG")
}
