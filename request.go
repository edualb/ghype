package ghype

import (
	"bytes"

	"github.com/edualb/ghype/internal"
)

type Request struct {
	Method Method
}

func newRequest(from []byte) (*Request, error) {
	r, err := buildFromRequestline(from)
	if err != nil {
		return nil, err
	}
	return r, nil
}

func buildFromRequestline(from []byte) (*Request, error) {
	byteSplited := bytes.Split(from, internal.SP.ToBytes())

	if len(byteSplited) > 3 {
		return nil, ErrRequestParse
	}

	m, err := buildFromMethod(byteSplited[0])

	if err != nil {
		return nil, ErrRequestParse
	}

	r := &Request{}

	r.Method = m

	return r, nil
}
