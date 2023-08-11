package request

import (
	"encoding/json"
	"log"
	"requestlogger/src/utils"
	"strconv"
	"strings"
)

type Content struct {
	Method           string
	HttpVersion      string
	Host             string
	Connection       string
	Server           string
	CacheControl     string
	Origin           string
	UserAgent        string
	Accept           []string
	AcceptEncoding   string
	AcceptLanguage   []string
	Referer          string
	TransferEncoding string
	XPoweredBy       string
	Date             string // FIXME: use date
	ContentType      []string
	ContentEncoding  string
	ContentLength    int32
	Cookies          []string
	Authorization    string // TODO
	ClientID         string
	ClientSessionID  string
	Payload          string // TODO
}

// BuildRequestContent Receive packet as a string and return the content
func BuildRequestContent(packet string) string {
	content := &Content{}
	content.IterateAndSetData(packet)
	jsonContent, err := json.MarshalIndent(content, "", " ")
	if err != nil {
		log.Fatal("Failed to parse JSON!")
	}
	return string(jsonContent)
}

func HasReadableContent(str string) bool {
	return strings.Contains(str, "HTTP")
}

func (content *Content) IterateAndSetData(packet string) {
	content.FindUserAgent(packet)
	content.FindDate(packet)
	content.FindServer(packet)
	content.FindAcceptEncoding(packet)
	words := strings.Fields(packet)
	for index, field := range words {
		content.FindMethod(field)
		content.FindHttpVersion(field)
		content.FindHost(field, index, words)
		content.FindConnection(field, index, words)
		content.FindReferer(field, index, words)
		content.FindCacheControl(field, index, words)
		content.FindXPoweredBy(field, index, words)
		content.FindContentEncoding(field, index, words)
		content.FindTransferEnconding(field, index, words)
		content.FindAccept(field, index, words)
		content.FindAcceptLanguage(field, index, words)
		content.FindContentLength(field, index, words)
		content.FindCookies(packet)
		content.FindClientSessionID(field, index, words)
		content.FindClientID(field, index, words)
		content.FindPayload(packet)
	}
}

func (content *Content) FindHttpVersion(field string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "HTTP/") || strings.Contains(f, "HTTPS/") {
		_, after, _ := strings.Cut(f, "HTTP")
		content.HttpVersion = "HTTP" + after
	}
}

func (content *Content) FindServer(packet string) {
	raw := strconv.Quote(packet)
	_, after, _ := strings.Cut(raw, "Server:")
	before, _, _ := strings.Cut(after, "\\r\\n")
	if before != "" {
		content.Server = strings.TrimSpace(before)
	}
}

func (content *Content) FindHost(field string, index int, words []string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "HOST:") {
		nextWord := words[index+1]
		content.Host = nextWord
	}
}

func (content *Content) FindConnection(field string, index int, words []string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "CONNECTION:") {
		nextWord := words[index+1]
		content.Connection = nextWord
	}
}

func (content *Content) FindCacheControl(field string, index int, words []string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "CACHE-CONTROL:") {
		nextWord := words[index+1]
		content.CacheControl = nextWord
	}
}

func (content *Content) FindReferer(field string, index int, words []string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "REFERER:") {
		nextWord := words[index+1]
		content.Referer = nextWord
	}
}

func (content *Content) FindDate(packet string) {
	raw := strconv.Quote(packet)
	_, after, _ := strings.Cut(raw, "Date:")
	before, _, _ := strings.Cut(after, "\\r\\n")
	if before != "" {
		content.Date = strings.TrimSpace(before)
	}
}

func (content *Content) FindXPoweredBy(field string, index int, words []string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "X-POWERED-BY:") {
		nextWord := words[index+1]
		content.XPoweredBy = nextWord
	}
}

func (content *Content) FindContentEncoding(field string, index int, words []string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "CONTENT-ENCODING:") {
		nextWord := words[index+1]
		content.ContentEncoding = nextWord
	}
}

func (content *Content) FindTransferEnconding(field string, index int, words []string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "TRANSFER-ENCODING:") {
		nextWord := words[index+1]
		content.TransferEncoding = nextWord
	}
}

func (content *Content) FindAccept(field string, index int, words []string) {
	f := field
	if strings.Contains(f, "Accept:") {
		acceptWords := strings.Split(words[index+1], ";")
		content.Accept = acceptWords
	}
}

func (content *Content) FindAcceptLanguage(field string, index int, words []string) {
	f := field
	if strings.Contains(f, "Accept-Language:") {
		acceptWords := strings.Split(words[index+1], ";")
		content.AcceptLanguage = acceptWords
	}
}

func (content *Content) FindContentLength(field string, index int, words []string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "Content-Length:") {
		nextWord := words[index+1]
		cLength, err := strconv.ParseInt(nextWord, 10, 32)
		if err != nil {
			log.Panic("Failed to convert Content Length to int32:", err)
		}
		content.ContentLength = int32(cLength)
	}
}

func (content *Content) FindUserAgent(packet string) {
	raw := strconv.Quote(packet)
	_, after, _ := strings.Cut(raw, "User-Agent:")
	before, _, _ := strings.Cut(after, "\\r\\n")
	if before != "" {
		content.UserAgent = strings.TrimSpace(before)
	}
}

func (content *Content) FindAcceptEncoding(packet string) {
	raw := strconv.Quote(packet)
	_, after, _ := strings.Cut(raw, "Accept-Encoding:")
	before, _, _ := strings.Cut(after, "\\r\\n")
	if before != "" {
		content.AcceptEncoding = strings.TrimSpace(before)
	}
}

func (content *Content) FindCookies(packet string) {
	raw := strconv.Quote(packet)
	_, after, _ := strings.Cut(raw, "Cookie:")
	before, _, _ := strings.Cut(after, "\\r\\n")
	if before != "" {
		split := strings.Split(before, ";")
		content.Cookies = split
	}
}

func (content *Content) FindClientID(field string, index int, words []string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "CLIENT-ID:") {
		nextWord := words[index+1]
		content.ClientID = nextWord
	}
}

func (content *Content) FindClientSessionID(field string, index int, words []string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "CLIENT-SESSION-ID:") {
		nextWord := words[index+1]
		content.ClientSessionID = nextWord
	}
}

func (content *Content) FindPayload(packet string) {
	raw := strconv.Quote(packet)
	_, after, _ := strings.Cut(raw, "\\r\\n\\r")
	after = strings.Replace(after, "\\n", "", -1)
	after = strings.Replace(after, "\\u", "", -1)
	after = utils.RemoveLastQuote(after)
	if after != "" {
		content.Payload = after
	}
}

// FindMethod Find the used method from the given string
func (content *Content) FindMethod(field string) {
	var method Method
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
	content.Method = method.String()
}
