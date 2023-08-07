package request

import (
	"encoding/json"
	"log"
	"strings"
)

type Content struct {
	Method         string
	HttpVersion    string
	Host           string
	Connection     string
	Server         string
	ContentLength  int32
	CacheControl   string
	Origin         string
	ContentType    []string
	UserAgent      string
	Accept         string
	Referer        string
	AcceptEncoding string
	AcceptLanguage string
	Payload        string
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
	words := strings.Fields(packet)
	for index, field := range words {
		content.FindMethod(field)
		content.FindHttpVersion(field)
		content.FindServer(field, index, words)
		content.FindHost(field, index, words)
		content.FindConnection(field, index, words)
		content.FindReferer(field, index, words)
		content.FindCacheControl(field, index, words)
	}
}

func (content *Content) FindHttpVersion(field string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "HTTP/") || strings.Contains(f, "HTTPS/") {
		_, after, _ := strings.Cut(f, "HTTP")
		content.HttpVersion = "HTTP" + after
	}
}

func (content *Content) FindServer(field string, index int, words []string) {
	f := strings.ToUpper(field)
	if strings.Contains(f, "SERVER:") {
		nextWord := words[index+1]
		content.Server = nextWord
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
