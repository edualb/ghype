package messageformat

import "github.com/edualb/ghype/internal/version"

func GetHTTPMessage() string {
	return ""
}

type HTTPMessage struct {
	startline startLine
}

// requestLine will get requestLine or statusLine
type startLine struct {
	requestLine requestLine
	statusLine  statusLine
}

func (sl startLine) Get() line {
	if sl.requestLine.status {
		return sl.requestLine
	}
	return sl.statusLine
}

type line interface {
	Get() string
}

type requestLine struct {
	method        string // token https://datatracker.ietf.org/doc/html/rfc7230#section-3.2.6
	requestTarget string //https://datatracker.ietf.org/doc/html/rfc7230#section-5.3
	httpVersion   string
	status        bool
}

func (l requestLine) Get() string {
	return " " + " " + version.HTTPVersion() + ""
}

type statusLine struct {
	status bool
}

func (l statusLine) Get() string {
	return ""
}
