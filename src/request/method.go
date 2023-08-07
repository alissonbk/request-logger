package request

type Method uint8

const (
	GET = iota
	POST
	PUT
	PATCH
	DELETE
	HEAD
	CONNECT
	OPTIONS
	TRACE
	UNKNOWN
)

func (m Method) String() string {
	switch m {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case PATCH:
		return "PATCH"
	case DELETE:
		return "DELETE"
	case HEAD:
		return "HEAD"
	case CONNECT:
		return "CONNECT"
	case OPTIONS:
		return "OPTIONS"
	case TRACE:
		return "TRACE"
	case UNKNOWN:
		return "UNKOWN METHOD"
	default:
		return "unkown method"
	}
}
