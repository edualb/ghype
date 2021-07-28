package ghype

import "bytes"

func buildFromRequestURI(from []byte) (string, error) {
	if bytes.Equal(from, []byte("*")) {
		return "*", nil
	}
	return "", nil
}
