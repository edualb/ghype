package messageformat_test

import (
	"testing"

	"github.com/edualb/ghype/internal/messageformat"
	"github.com/edualb/ghype/internal/methods"
	"github.com/edualb/ghype/internal/version"
	"github.com/google/go-cmp/cmp"
)

func TestGetHTTPMessage(t *testing.T) {
	type testTemplate struct {
		name    string
		httpMsg string
		exp     messageformat.HTTPMessage
		expErr  bool
	}
	tests := []testTemplate{
		{
			name:    "should return ok when have right http message",
			httpMsg: "GET /where?q=now HTTP/1.1",
			exp: messageformat.HTTPMessage{
				Startline: messageformat.StartLine{
					RequestLine: messageformat.RequestLine{
						Method:        methods.Get,
						RequestTarget: "",
						HTTPVersion:   version.HTTP(),
						Status:        true,
					},
				},
			},
			expErr: false,
		},
		{
			name:    "should return error when have wrong http version",
			httpMsg: "GET /where?q=now HTTP/2.1",
			exp: messageformat.HTTPMessage{
				Startline: messageformat.StartLine{
					RequestLine: messageformat.RequestLine{
						Method:        "",
						RequestTarget: "",
						HTTPVersion:   "",
						Status:        false,
					},
				},
			},
			expErr: true,
		},
		{
			name:    "should return error when have right http message",
			httpMsg: "get /where?q=now HTTP/1.1",
			exp: messageformat.HTTPMessage{
				Startline: messageformat.StartLine{
					RequestLine: messageformat.RequestLine{
						Method:        "",
						RequestTarget: "",
						HTTPVersion:   "",
						Status:        false,
					},
				},
			},
			expErr: true,
		},
		{
			name:    "should return error when have space in the final of the message",
			httpMsg: "GET /where?q=now HTTP/1.1 ",
			exp: messageformat.HTTPMessage{
				Startline: messageformat.StartLine{
					RequestLine: messageformat.RequestLine{
						Method:        "",
						RequestTarget: "",
						HTTPVersion:   "",
						Status:        false,
					},
				},
			},
			expErr: true,
		},
		{
			name:    "should return error when have space in the beginning of the message",
			httpMsg: " GET /where?q=now HTTP/1.1",
			exp: messageformat.HTTPMessage{
				Startline: messageformat.StartLine{
					RequestLine: messageformat.RequestLine{
						Method:        "",
						RequestTarget: "",
						HTTPVersion:   "",
						Status:        false,
					},
				},
			},
			expErr: true,
		},
		{
			name:    "should return error when don't have request-target in the message",
			httpMsg: "GET HTTP/1.1",
			exp: messageformat.HTTPMessage{
				Startline: messageformat.StartLine{
					RequestLine: messageformat.RequestLine{
						Method:        "",
						RequestTarget: "",
						HTTPVersion:   "",
						Status:        false,
					},
				},
			},
			expErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := messageformat.GetHTTPMessage(tt.httpMsg)
			if tt.expErr {
				if err == nil {
					t.Error("unexpected err equals nil")
				}
			} else {
				if err != nil {
					t.Errorf("unexpected err, got %v, wants, nil", got)
				}
			}
			if !cmp.Equal(got, tt.exp) {
				t.Errorf("unexpected HTTPMessage, got %v, wants, %v", got, tt.exp)
			}
		})
	}
}
