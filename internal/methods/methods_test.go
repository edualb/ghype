package methods_test

import (
	"testing"

	"github.com/edualb/ghype/internal/methods"
)

func TestIsValid(t *testing.T) {
	type testTemplate struct {
		name   string
		method string
		exp    bool
	}
	tests := []testTemplate{
		{
			name:   "should return true when has GET",
			method: methods.Get,
			exp:    true,
		},
		{
			name:   "should return false when has Get",
			method: "Get",
			exp:    false,
		},
		{
			name:   "should return true when has POST",
			method: methods.Post,
			exp:    true,
		},
		{
			name:   "should return false when has Post",
			method: "Post",
			exp:    false,
		},
		{
			name:   "should return true when has HEAD",
			method: methods.Head,
			exp:    true,
		},
		{
			name:   "should return false when has Head",
			method: "Head",
			exp:    false,
		},
		{
			name:   "should return true when has PUT",
			method: methods.Put,
			exp:    true,
		},
		{
			name:   "should return false when has Put",
			method: "Put",
			exp:    false,
		},
		{
			name:   "should return true when has DELETE",
			method: methods.Delete,
			exp:    true,
		},
		{
			name:   "should return false when has delete",
			method: "delete",
			exp:    false,
		},
		{
			name:   "should return true when has OPTIONS",
			method: methods.Options,
			exp:    true,
		},
		{
			name:   "should return false when has OptionS",
			method: "OptionS",
			exp:    false,
		},
		{
			name:   "should return true when has TRACE",
			method: methods.Trace,
			exp:    true,
		},
		{
			name:   "should return false when has trace",
			method: "trace",
			exp:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := methods.IsValid(tt.method)
			if got != tt.exp {
				t.Errorf("unexpected validation, got %t, wants %t", got, tt.exp)
			}
		})
	}
}
