package version_test

import (
	"testing"

	"github.com/edualb/ghype/internal/version"
)

func TestHTTPVersion(t *testing.T) {
	type testTemplate struct {
		name string
		exp  string
	}
	tests := []testTemplate{
		{
			name: "should return HTTP/1.1 when call version.HTTPVersion()",
			exp:  "HTTP/1.1",
		},
	}
	for _, tt := range tests {
		got := version.HTTPVersion()
		if got != tt.exp {
			t.Errorf("unexpected http version, got %s wants %s", got, tt.exp)
		}
	}
}
