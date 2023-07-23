package request

import (
	"strings"
)

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

// FindMethod Find the used method from the given string
func FindMethod(s string) Method {
	words := strings.Fields(s)
	var method Method
	for _, field := range words {
		f := strings.ToUpper(field)
		switch {
		case strings.Contains(f, "POST"):
			method = POST
		case strings.Contains(f, "GET"):
			method = GET
		case strings.Contains(f, "PUT"):
			method = PUT
		case strings.Contains(f, "PATCH"):
			method = PATCH
		case strings.Contains(f, "DELETE"):
			method = DELETE
		case strings.Contains(f, "HEAD"):
			method = HEAD
		case strings.Contains(f, "CONNECT"):
			if !strings.Contains(f, "CONNECTION") {
				method = CONNECT
			}
		case strings.Contains(f, "OPTIONS"):
			method = OPTIONS
		case strings.Contains(f, "TRACE"):
			method = TRACE
		}
	}
	return method
}

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
