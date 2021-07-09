package option

import "errors"

var (
	ErrInvalidNetwork = errors.New("HTTP Accept reliable networks only. Try: tcp, tcp4, tcp6, unix or unixpackage")
)
