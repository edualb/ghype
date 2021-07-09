package ghype

import "errors"

var (
	ErrServerNotStarted = errors.New("Server not started. Run GHype.Serve() before.")
)
