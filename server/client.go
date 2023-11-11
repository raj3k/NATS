package server

import (
	"fmt"
	"net"
)

const (
	CRLF = "\r\n"
)

const (
	pong = "PONG" + CRLF
	ok   = "+OK" + CRLF
)

type client struct {
	parseState
	out outbound
	nc  net.Conn
}

type outbound struct {
	nb net.Buffers
}

func NewClient(c net.Conn) *client {
	return &client{
		nc: c,
	}
}

func (c *client) processPub(arg []byte) error {
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
	c.out.nb = append(c.out.nb, []byte(ok))
	c.out.nb.WriteTo(c.nc)
}

func (c *client) processPing() {
	c.sendPong()
}

func (c *client) sendPong() {
	// TODO: work on outbound & processing messages to client
	c.out.nb = append(c.out.nb, []byte(pong))
	c.out.nb.WriteTo(c.nc)
}

func (c *client) processConnect(arg []byte) {
	fmt.Println(string(arg))
	c.out.nb = append(c.out.nb, []byte(ok))
	c.out.nb.WriteTo(c.nc)
}
