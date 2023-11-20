package server

import (
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
	args := splitArg(arg)

	c.pa.arg = arg
	c.pa.topic = args[0]
	c.pa.size = parseSize(args[1])
	return nil
}

func (c *client) processInboundMessage(msg []byte) {

	topic := c.pa.topic

	c.mu.Lock()
	c.srv.topics[string(topic)].Enqueue(msg)
	c.mu.Unlock()

	c.out.nb = append(c.out.nb, []byte(ok))
	c.out.nb.WriteTo(c.nc)

	for _, c := range c.srv.clients {
		for _, sub := range c.subs {
			c.deliverMsg(sub, topic, msg)
		}
	}
}

func (c *client) deliverMsg(sub *subscription, topic, msg []byte) {
	// client := sub.client

	c.out.nb = append(c.out.nb, msg)
	c.out.nb = append(c.out.nb, []byte(CRLF))
}

func (c *client) processPing() {
	c.sendPong()
}

func (c *client) sendPong() {
	c.out.nb = append(c.out.nb, []byte(pong))
}

func (c *client) processConnect(arg []byte) {
	c.out.nb = append(c.out.nb, []byte(ok))
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
		c.srv.createTopic(string(topic))
	}

	c.mu.Unlock()

	c.out.nb = append(c.out.nb, []byte(ok))
}

func (c *client) writeLoop() {
	for {
		c.mu.Lock()

		c.out.nb.WriteTo(c.nc)

		c.mu.Unlock()
	}
}
