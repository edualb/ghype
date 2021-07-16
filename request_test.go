package ghype

import (
	"fmt"
	"testing"
)

func TestNewRequest(t *testing.T) {
	type input struct {
		from string
	}

	type output struct {
		r        *Request
		gHypeErr error
	}

	type test struct {
		name string
		i    input
		o    output
	}

	tests := []test{
		{
			name: "newRequest with an empty value should return nil request and err = ErrRequestParse",
			i: input{
				from: "",
			},
			o: output{
				r:        nil,
				gHypeErr: ErrRequestParse,
			},
		},
		{
			name: "newRequest with GET method should return a valid request and err = nil",
			i: input{
				from: "GET * something",
			},
			o: output{
				r: &Request{
					Method: GET,
				},
				gHypeErr: nil,
			},
		},
		{
			name: "newRequest with POST method should return a valid request and err = nil",
			i: input{
				from: "POST * something",
			},
			o: output{
				r: &Request{
					Method: POST,
				},
				gHypeErr: nil,
			},
		},
		{
			name: "newRequest with PUT method should return a valid request and err = nil",
			i: input{
				from: "PUT * something",
			},
			o: output{
				r: &Request{
					Method: PUT,
				},
				gHypeErr: nil,
			},
		},
		{
			name: "newRequest with DELETE method should return a valid request and err = nil",
			i: input{
				from: "DELETE * something",
			},
			o: output{
				r: &Request{
					Method: DELETE,
				},
				gHypeErr: nil,
			},
		},
		{
			name: "newRequest with HEAD method should return a valid request and err = nil",
			i: input{
				from: "HEAD * something",
			},
			o: output{
				r: &Request{
					Method: HEAD,
				},
				gHypeErr: nil,
			},
		},
		{
			name: "newRequest with OPTION method should return a valid request and err = nil",
			i: input{
				from: "OPTION * something",
			},
			o: output{
				r: &Request{
					Method: OPTION,
				},
				gHypeErr: nil,
			},
		},
		{
			name: "newRequest with TRACE method should return a valid request and err = nil",
			i: input{
				from: "TRACE * something",
			},
			o: output{
				r: &Request{
					Method: TRACE,
				},
				gHypeErr: nil,
			},
		},
		{
			name: "newRequest with CONNECT method should return a valid request and err = nil",
			i: input{
				from: "CONNECT * something",
			},
			o: output{
				r: &Request{
					Method: CONNECT,
				},
				gHypeErr: nil,
			},
		},
		{
			name: "newRequest with ANOTHER (valid) method should return a valid request and err = nil",
			i: input{
				from: "JOJO * something",
			},
			o: output{
				r: &Request{
					Method: ANOTHER,
				},
				gHypeErr: nil,
			},
		},
		{
			name: "newRequest with ANOTHER (invalid) method should return a valid request and err = nil",
			i: input{
				from: "JO<JO * something",
			},
			o: output{
				r:        nil,
				gHypeErr: ErrRequestParse,
			},
		},
		{
			name: "newRequest with * value in  should return a valid request and err = nil",
			i: input{
				from: "JO<JO * something",
			},
			o: output{
				r:        nil,
				gHypeErr: ErrRequestParse,
			},
		},
	}

	for _, tt := range tests {
		r, err := newRequest([]byte(tt.i.from))

		if err != tt.o.gHypeErr {
			t.Error(fmt.Sprintf("[%s]: want err %v, got %v", tt.name, tt.o.gHypeErr, err))
			return
		}

		if tt.o.r != nil {
			if r.Method != tt.o.r.Method {
				t.Error(fmt.Sprintf("[%s]: want method %s, got %v", tt.name, tt.o.r.Method.Value(), r.Method.Value()))
			}
		} else {
			if r != nil {
				t.Error(fmt.Sprintf("[%s]: want Request %v, got %v", tt.name, tt.o.r, r))
			}
		}
	}
}
