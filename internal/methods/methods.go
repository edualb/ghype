package methods

import "errors"

const (
	// Get is based on RFC 7231 - Transfer a current representation of the target resource.
	// more information in section 4.3.1
	Get = "GET"
	// Head is based on RFC 7231 - Same as GET, except that the server MUST NOT send a message body in the response.
	// more information in section 4.3.2
	Head = "HEAD"
	// Post is based on RFC 7231 - Perform resource-specific processing on the request payload.
	// more information in section 4.3.3
	Post = "POST"
	// Put is based on RFC 7231 - Replace all current representations of the target resource with the request payload.
	// more information in section 4.3.4
	Put = "PUT"
	// Delete is based on RFC 7231 - Remove all current representations of the target resource.
	// more information in section 4.3.5
	Delete = "DELETE"
	// Connect is based on RFC 7231 - Establish a tunnel to the server identified by the target resource.
	// more information in section 4.3.6
	Connect = "CONNECT"
	// Options is based on RFC 7231 - Describe the communication options for the target resource.
	// more information in section 4.3.7
	Options = "OPTIONS"
	// Trace is based on RFC 7231 - Perform a message loop-back test along the path to the target resource.
	// more information in section 4.3.8
	Trace = "TRACE"
)

var (
	ErrMethodNotAllowed = errors.New("method not allowed")
)

func IsValid(method string) bool {
	methods := []string{Get, Head, Post, Put, Delete, Connect, Options, Trace}
	for _, m := range methods {
		if method == m {
			return true
		}
	}
	return false
}
