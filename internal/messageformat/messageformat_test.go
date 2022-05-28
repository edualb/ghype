package messageformat_test

import (
	"testing"

	"github.com/edualb/ghype/internal/messageformat"
)

func TestGetHTTPMessage(t *testing.T) {
	type testTemplate struct {
		name string
		exp  string
	}
	tests := []testTemplate{
		{
			name: "should return ok when call GetHTTPMessage",
			exp:  "",
		},
	}
	for _, tt := range tests {
		got := messageformat.GetHTTPMessage()
		if tt.exp != got {
			t.Errorf("unexpected HTTPMessage, got %s, wants, %s", got, tt.exp)
		}
	}
}
