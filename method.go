package ghype

import (
	"bytes"

	"github.com/edualb/ghype/internal"
)

type Method int

const (
	OPTION = iota
	GET
	HEAD
	POST
	PUT
	DELETE
	TRACE
	CONNECT
	ANOTHER
)

func (m Method) value() []byte {
	return [...][]byte{
		[]byte("OPTION"),
		[]byte("GET"),
		[]byte("HEAD"),
		[]byte("POST"),
		[]byte("PUT"),
		[]byte("DELETE"),
		[]byte("TRACE"),
		[]byte("CONNECT"),
		[]byte("ANOTHER"),
	}[m]
}

func (m Method) Value() string {
	return [...]string{
		"OPTION",
		"GET",
		"HEAD",
		"POST",
		"PUT",
		"DELETE",
		"TRACE",
		"CONNECT",
		"ANOTHER",
	}[m]
}

func buildFromMethod(from []byte) (Method, error) {
	var m Method

	switch {
	case bytes.Equal(from, []byte("GET")):
		m = GET
		break
	case bytes.Equal(from, []byte("PUT")):
		m = PUT
		break
	case bytes.Equal(from, []byte("DELETE")):
		m = DELETE
		break
	case bytes.Equal(from, []byte("POST")):
		m = POST
		break
	case bytes.Equal(from, []byte("CONNECT")):
		m = CONNECT
		break
	case bytes.Equal(from, []byte("HEAD")):
		m = HEAD
		break
	case bytes.Equal(from, []byte("TRACE")):
		m = TRACE
		break
	case bytes.Equal(from, []byte("OPTION")):
		m = OPTION
		break
	default:
		m = ANOTHER
		break
	}

	if !m.isMethodValid(from) {
		return m, ErrRequestParse
	}

	return m, nil
}

func (m Method) isMethodValid(from []byte) bool {
	if m != ANOTHER {
		return true
	}
	return internal.IsValidToken(from)
}
