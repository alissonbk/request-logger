package request

type Connection int8

const (
	KEEPALIVE = iota
	CLOSE
)

func (c Connection) String() string {
	switch c {
	case KEEPALIVE:
		return "keep-alive"
	case CLOSE:
		return "close"
	default:
		return "unkown connection"
	}
}
