package ghype_test

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/edualb/ghype"
	"github.com/edualb/ghype/option"
)

func TestGHypeServeAndClose(t *testing.T) {

	type input struct {
		addr        string
		network     string
		dialAddr    string
		dialNetwork string
	}

	type output struct {
		gHypeErr   error
		hasConnErr bool
	}

	type test struct {
		name string
		i    input
		o    output
	}

	tests := []test{
		{
			name: "Default GHype should connect",
			i: input{
				addr:        "",
				network:     "",
				dialAddr:    ":8080",
				dialNetwork: "tcp",
			},
			o: output{
				gHypeErr:   nil,
				hasConnErr: false,
			},
		},
		{
			name: "GHype WithAddress(:1234) should connect",
			i: input{
				addr:        ":1234",
				network:     "",
				dialAddr:    ":1234",
				dialNetwork: "tcp",
			},
			o: output{
				gHypeErr:   nil,
				hasConnErr: false,
			},
		},
		{
			name: "GHype WithAddress(:1234) and WithNetwork(\"tcp6\") should connect",
			i: input{
				addr:        ":1234",
				network:     "tcp6",
				dialAddr:    ":1234",
				dialNetwork: "tcp6",
			},
			o: output{
				gHypeErr:   nil,
				hasConnErr: false,
			},
		},
		{
			name: "GHype WithNetwork(\"tcp4\") should connect",
			i: input{
				addr:        "",
				network:     "tcp4",
				dialAddr:    ":8080",
				dialNetwork: "tcp4",
			},
			o: output{
				gHypeErr:   nil,
				hasConnErr: false,
			},
		},
		{
			name: "GHype WithNetwork(\"tcp6\") should connect",
			i: input{
				addr:        "",
				network:     "tcp6",
				dialAddr:    ":8080",
				dialNetwork: "tcp6",
			},
			o: output{
				gHypeErr:   nil,
				hasConnErr: false,
			},
		},
		{
			name: "GHype WithNetwork(\"unix\") should connect",
			i: input{
				addr:        "",
				network:     "unix",
				dialAddr:    ":8080",
				dialNetwork: "unix",
			},
			o: output{
				gHypeErr:   nil,
				hasConnErr: false,
			},
		},
		{
			name: "GHype WithNetwork(\"unixpacket\") should connect",
			i: input{
				addr:        "",
				network:     "unixpacket",
				dialAddr:    ":8080",
				dialNetwork: "unix",
			},
			o: output{
				gHypeErr:   nil,
				hasConnErr: false,
			},
		},
		{
			name: "Trying GHype UDP connection should return ErrInvalidNetwork",
			i: input{
				addr:        "",
				network:     "udp",
				dialAddr:    ":8080",
				dialNetwork: "udp",
			},
			o: output{
				gHypeErr:   option.ErrInvalidNetwork,
				hasConnErr: false,
			},
		},
		{
			name: "Trying GHype wrong connection should return ErrInvalidNetwork",
			i: input{
				addr:        "",
				network:     "",
				dialAddr:    ":8080",
				dialNetwork: "unix",
			},
			o: output{
				gHypeErr:   nil,
				hasConnErr: true,
			},
		},
	}

	for _, tt := range tests {
		o := []option.GHypeOptions{}

		if len(tt.i.addr) > 0 {
			o = append(o, option.WithAddress(tt.i.addr))
		}

		if len(tt.i.network) > 0 {
			o = append(o, option.WithNetwork(tt.i.network))
		}

		g, err := ghype.New(o...)

		if err != tt.o.gHypeErr {
			t.Error(fmt.Sprintf("[%s]: want gHypeErr %v, got %v", tt.name, tt.o.gHypeErr, err))
			return
		}

		if tt.o.gHypeErr != nil {
			break
		}

		go g.Serve()
		// Wait server start
		time.Sleep(10 * time.Millisecond)

		defer g.Close()

		conn, err := net.Dial(tt.i.dialNetwork, tt.i.dialAddr)
		if err != nil {
			if !tt.o.hasConnErr {
				t.Error(fmt.Sprintf("[%s]: Does not want err, got %v", tt.name, err))
			}
			return
		}
		conn.Close()
	}
}
