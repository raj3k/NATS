package server

import (
	"fmt"
)

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
	OP_C
	OP_CO
)

func (c *client) parse(buf []byte) error {
	var b byte
	var i int
	for i = 0; i < len(buf); i++ {
		b = buf[i]

		switch c.state {
		case OP_START:
			switch b {
			case 'P', 'p':
				c.state = OP_P
			case 'C', 'c':
				c.state = OP_C
			default:
				goto parseErr
			}
		case OP_P:
			switch b {
			case 'I', 'i':
				c.state = OP_PI
			case 'O', 'o':
				c.state = OP_PO
			default:
				goto parseErr
			}
		case OP_PI:
			switch b {
			case 'N', 'n':
				c.state = OP_PIN
			default:
				goto parseErr
			}
		case OP_PIN:
			switch b {
			case 'G', 'g':
				c.state = OP_PING
			default:
				goto parseErr
			}
		case OP_PING:
			switch b {
			case '\n':
				c.state = OP_START
			}
		case OP_PO:
			switch b {
			case 'N', 'n':
				c.state = OP_PON
			default:
				goto parseErr
			}
		case OP_PON:
			switch b {
			case 'G', 'g':
				c.state = OP_PONG
			default:
				goto parseErr
			}
		case OP_PONG:
			switch b {
			case '\n':
				c.state = OP_START
			}
		case OP_C:
			switch b {
			case 'O', 'o':
				c.state = OP_CO
			default:
				goto parseErr
			}
		default:
			goto parseErr
		}
	}

	return nil
parseErr:
	err := fmt.Errorf("parser ERROR, state=%d, i=%d", c.state, i)
	return err
}
