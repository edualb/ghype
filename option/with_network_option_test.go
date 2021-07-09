package option_test

import (
	"fmt"
	"testing"

	"github.com/edualb/ghype"
	"github.com/edualb/ghype/option"
)

func TestWithNetwork(t *testing.T) {

	type input struct {
		network string
	}

	type output struct {
		gHypeErr error
	}

	type test struct {
		name string
		i    input
		o    output
	}

	tests := []test{
		{
			name: "WithNetwork with an empty network should return err = ErrInvalidNetwork",
			i: input{
				network: "",
			},
			o: output{
				gHypeErr: option.ErrInvalidNetwork,
			},
		},
		{
			name: "WithNetwork with \"udp\" network should return err = nil",
			i: input{
				network: "udp",
			},
			o: output{
				gHypeErr: option.ErrInvalidNetwork,
			},
		},
		{
			name: "WithNetwork with \"tcp\" network should return err = nil",
			i: input{
				network: "tcp",
			},
			o: output{
				gHypeErr: nil,
			},
		},
	}

	for _, tt := range tests {
		o := []option.GHypeOptions{}

		o = append(o, option.WithNetwork(tt.i.network))

		_, err := ghype.New(o...)

		if err != tt.o.gHypeErr {
			t.Error(fmt.Sprintf("[%s]: want gHypeErr %v, got %v", tt.name, tt.o.gHypeErr, err))
			return
		}
	}
}
