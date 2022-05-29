package messageformat

import (
	"errors"
	"strings"

	"github.com/edualb/ghype/internal/methods"
	"github.com/edualb/ghype/internal/version"
)

const (
	cr   = string(rune(13))
	lf   = string(rune(10))
	CRLF = cr + lf
)

var (
	ErrMethodNotAllowed = methods.ErrMethodNotAllowed
	ErrInvalidMessage   = errors.New("invalid HTTP message to process")
)

func GetHTTPMessage(message string) (HTTPMessage, error) {
	httpMsg := &HTTPMessage{}
	startLine := &StartLine{}

	requestLine, err := getRequestLine(message)
	if err != nil {
		return *httpMsg, err
	}
	startLine.RequestLine = requestLine

	httpMsg.Startline = *startLine
	return *httpMsg, err
}

type HTTPMessage struct {
	Startline StartLine
}

// requestLine will get requestLine or statusLine
type StartLine struct {
	RequestLine RequestLine
}

type RequestLine struct {
	Method        string // token https://datatracker.ietf.org/doc/html/rfc7230#section-3.2.6
	RequestTarget string //https://datatracker.ietf.org/doc/html/rfc7230#section-5.3
	HTTPVersion   string
	Status        bool
}

func getRequestLine(message string) (RequestLine, error) {
	r := &RequestLine{}

	lines := strings.Split(message, " ")
	if len(lines) != 3 {
		return *r, ErrInvalidMessage
	}

	ok := methods.IsValid(lines[0])
	if !ok {
		return *r, ErrMethodNotAllowed // A server that receives a method longer than any that it implements SHOULD respond with a 501 (Not Implemented) status code.
	}
	r.Method = lines[0]

	ok = version.IsValid(lines[2])
	if !ok {
		return *r, errors.New("wrong HTTP version, this protocol just accept " + version.HTTP())
	}
	r.HTTPVersion = lines[2]

	r.Status = true
	r.RequestTarget = "" // request-target longer than any URI it wishes to parse MUST respond with a 414 (URI Too Long) status code

	return *r, nil
}
