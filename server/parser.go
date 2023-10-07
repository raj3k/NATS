package server

import (
	"fmt"
)

type parserState int
type parseState struct {
	state   parserState
	argBuf  []byte
	msgBuff []byte
	pa      PubArg
}

type PubArg struct {
	arg     []byte
	subject []byte
	size    int
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
	OP_CON
	OP_CONN
	OP_CONNE
	OP_CONNEC
	OP_CONNECT
	CONNECT_ARG
	OP_PU
	OP_PUB
	OP_PUB_SPACE
	PUB_ARG
	OP_S
	OP_SU
	OP_SUB
	SUB_ARG
	MSG_PAYLOAD
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
			case 'U', 'u':
				c.state = OP_PU
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
				//TODO: process ping command
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
				//TODO: process PONG command
				c.state = OP_START
			}
		case OP_C:
			switch b {
			case 'O', 'o':
				c.state = OP_CO
			default:
				goto parseErr
			}
		case OP_CO:
			switch b {
			case 'N', 'n':
				c.state = OP_CON
			default:
				goto parseErr
			}
		case OP_CON:
			switch b {
			case 'N', 'n':
				c.state = OP_CONN
			default:
				goto parseErr
			}
		case OP_CONN:
			switch b {
			case 'E', 'e':
				c.state = OP_CONNE
			default:
				goto parseErr
			}
		case OP_CONNE:
			switch b {
			case 'C', 'c':
				c.state = OP_CONNEC
			default:
				goto parseErr
			}
		case OP_CONNEC:
			switch b {
			case 'T', 't':
				c.state = OP_CONNECT
			default:
				goto parseErr
			}
		case OP_CONNECT:
			switch b {
			case ' ', '\t':
				continue
			default:
				c.state = CONNECT_ARG
			}
		case CONNECT_ARG:
			switch b {
			case '\r':
				continue
			case '{':
				continue
			case '}':
				continue
			case '\n':
				//TODO: process CONNECT {} command
				c.state = OP_START
			default:
				c.argBuf = append(c.argBuf, b)
			}
		case OP_PU:
			switch b {
			case 'B', 'b':
				c.state = OP_PUB
			default:
				goto parseErr
			}
		case OP_PUB:
			switch b {
			case ' ', '\t':
				c.state = OP_PUB_SPACE
			default:
				goto parseErr
			}
		case OP_PUB_SPACE:
			switch b {
			case ' ', '\t':
				continue
			default:
				c.state = PUB_ARG
				// TODO: for now, think how to implement this in other way
				c.argBuf = append(c.argBuf, b)
			}
		case PUB_ARG:
			switch b {
			case '\r':
			case '\n':
				var arg []byte
				if c.argBuf != nil {
					arg = c.argBuf
					c.argBuf = nil
				} else {
					arg = buf[:]
				}
				err := c.processPub(arg)
				if err != nil {
					return err
				}
				//TODO: think if this logic is sufficient for now
				c.state = MSG_PAYLOAD
			default:
				if c.argBuf != nil {
					c.argBuf = append(c.argBuf, b)
				}
			}
		case MSG_PAYLOAD:
			switch b {
			default:
				print("test")
			}
		// TODO: implement login for MSG payload
		default:
			goto parseErr
		}
	}

	return nil
parseErr:
	err := fmt.Errorf("parser ERROR, state=%d, i=%d", c.state, i)
	return err
}
