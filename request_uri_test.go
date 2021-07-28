package ghype

import (
	"fmt"
	"testing"
)

func TestBuildFromRequestURI(t *testing.T) {
	type input struct {
		from string
	}

	type output struct {
		rURI     string
		gHypeErr error
	}

	type test struct {
		name string
		i    input
		o    output
	}

	tests := []test{
		{
			name: "buildFromRequestURI with * value should return * string and err = nil",
			i: input{
				from: "*",
			},
			o: output{
				rURI:     "*",
				gHypeErr: nil,
			},
		},
	}

	for _, tt := range tests {
		rURI, err := buildFromRequestURI([]byte(tt.i.from))
		if err != tt.o.gHypeErr {
			t.Error(fmt.Sprintf("[%s]: want err %v, got %v", tt.name, tt.o.gHypeErr, err))
			return
		}
		if tt.o.rURI != rURI {
			t.Error(fmt.Sprintf("[%s]: want rURI %s, got %v", tt.name, tt.o.rURI, rURI))
		}
	}
}
