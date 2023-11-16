package server

import (
	"fmt"
	"net"
	"sync"
	"sync/atomic"
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
	out  outbound
	nc   net.Conn
	cid  uint64
	srv  *Server
	subs map[string]*subscription
	mu   sync.RWMutex
}

type subscription struct {
	client *client
	topic  []byte
	sid    []byte
}

type outbound struct {
	nb net.Buffers
}

func NewClient(c net.Conn, srv *Server) *client {
	return &client{
		nc:   c,
		cid:  atomic.AddUint64(&srv.totalClients, 1),
		srv:  srv,
		subs: make(map[string]*subscription),
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
	c.pa.topic = args[0]
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
	c.out.nb = append(c.out.nb, []byte(pong))
	c.out.nb.WriteTo(c.nc)
}

func (c *client) processConnect(arg []byte) {
	c.out.nb = append(c.out.nb, []byte(ok))
	c.out.nb.WriteTo(c.nc)
}

func (c *client) parseSub(argo []byte) {

	args := splitArg(argo)
	t, sid := args[0], args[1]

	c.processSub(t, sid)
}

func (c *client) processSub(topic []byte, bsid []byte) {

	sub := &subscription{
		client: c,
		topic:  topic,
		sid:    bsid,
	}

	c.mu.Lock()

	sid := string(sub.sid)

	s := c.subs[sid]
	if s == nil {
		c.subs[sid] = sub
	}

	c.mu.Unlock()

	c.out.nb = append(c.out.nb, []byte(ok))
	c.out.nb.WriteTo(c.nc)
}
