package option_test

import (
	"fmt"
	"testing"

	"github.com/edualb/ghype"
	"github.com/edualb/ghype/option"
)

func TestWithAddress(t *testing.T) {

	type input struct {
		addr string
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
			name: "WithAddress with an empty address should return err = ErrInvalidAddress",
			i: input{
				addr: "",
			},
			o: output{
				gHypeErr: option.ErrInvalidAddress,
			},
		},
		{
			name: "WithAddress with \":1234\" address should return err = nil",
			i: input{
				addr: ":1234",
			},
			o: output{
				gHypeErr: nil,
			},
		},
	}

	for _, tt := range tests {
		o := []option.GHypeOptions{}

		o = append(o, option.WithAddress(tt.i.addr))

		_, err := ghype.New(o...)

		if err != tt.o.gHypeErr {
			t.Error(fmt.Sprintf("[%s]: want gHypeErr %v, got %v", tt.name, tt.o.gHypeErr, err))
			return
		}
	}
}
