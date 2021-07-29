package ghype

import (
	"bytes"

	"github.com/edualb/ghype/internal"
)

func buildFromRequestURI(from []byte) (string, error) {
	if bytes.Equal(from, []byte("*")) {
		return "*", nil
	}

	if bytes.ContainsAny(from, ":") {
		b := bytes.Split(from, []byte(":"))
		if len(b) != 2 {
			return "", ErrRequestURIParse
		}
		if len(b[0]) <= 0 || len(b[1]) <= 0 {
			return "", ErrRequestURIParse
		}
		scheme := b[0]
		if !isSchemeValid(scheme) {
			return "", ErrRequestURIParse
		}
	}

	return string(from), nil
}

func isSchemeValid(scheme []byte) bool {
	if len(scheme) <= 0 {
		return false
	}
	if !internal.IsByteAlpha(scheme[0]) {
		return false
	}

	for i := 1; i < len(scheme); i++ {
		if internal.IsByteAlpha(scheme[i]) || internal.IsByteDigit(scheme[i]) || bytes.ContainsAny([]byte{scheme[i]}, "+") || bytes.ContainsAny([]byte{scheme[i]}, "-") || bytes.ContainsAny([]byte{scheme[i]}, ".") {
			continue
		}
		return false
	}

	return true
}
