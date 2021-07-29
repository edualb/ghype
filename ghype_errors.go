package ghype

import "errors"

var (
	ErrServerNotStarted = errors.New("Server not started. Run GHype.Serve() before.")
	ErrRequestParse     = errors.New("It is not possible to create a request from this value. Follow the correct message. See more in https://datatracker.ietf.org/doc/html/rfc2616#section-5")
	ErrRequestURIParse  = errors.New("It is not possible to create a requestURI from this value. Follow the correct message. See more in https://datatracker.ietf.org/doc/html/rfc2396#section-3")
)
