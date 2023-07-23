package request

import "strings"

type RequestContent struct {
	method         string
	httpVersion    string
	host           string
	connection     Connection
	contentLength  int32
	cacheControl   string
	origin         string
	contentType    []string
	userAgent      string
	accept         string
	refer          string
	acceptEncoding string
	acceptLanguage string
	payload        string
}

// BuildRequestContent Receive packet as a string and return the content
func BuildRequestContent(packet string) RequestContent {
	return RequestContent{
		method: FindMethod(packet).String(),
	}
}

func HasReadableContent(str string) bool {
	return strings.Contains(str, "HTTP")
}
