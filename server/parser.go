package server

type parserState int
type parseState struct {
	state parserState
}

const (
	OP_START parserState = iota
	OP_P
	OP_PI
	OP_PIN
	OP_PING
	OP_PO
	OP_PON
	OP_PONG
)

func (c *client) parse(buf []byte) error {
	for i := 0; i < len(buf); i++ {
		b := buf[i]
		switch c.state {
		case OP_START:
			switch b {
			case 'P', 'p':
				c.state = OP_P
			}
		case OP_P:
			switch b {
			case 'I', 'i':
				c.state = OP_PI
			case 'O', 'o':
				c.state = OP_PO
			}
		case OP_PI:
			switch b {
			case 'N', 'n':
				c.state = OP_PIN
			}
		case OP_PO:
			switch b {
			case 'N', 'n':
				c.state = OP_PON
			}
		case OP_PIN:
			switch b {
			case 'G', 'g':
				c.state = OP_PING
			}
		case OP_PON:
			switch b {
			case 'G', 'g':
				c.state = OP_PONG
			}
		case OP_PING:
			switch b {
			case '\n':
				c.state = OP_START
			}
		case OP_PONG:
			switch b {
			case '\n':
				c.state = OP_START
			}
		}
	}
	return nil
}
