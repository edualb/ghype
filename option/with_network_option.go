package option

import "github.com/edualb/ghype/internal"

type withNetwork struct{ network string }

// WithNetwork is responsable to configure the network value in GHypeSettings.
//
// The network value must be a reliable protocol: "tcp", "tcp4", "tcp6", "unix" or "unixpacket".
//
// More information: https://tools.ietf.org/html/rfc2616
func WithNetwork(n string) GHypeOptions {
	return withNetwork{
		network: n,
	}
}

func (w withNetwork) Apply(g *internal.GHypeSettings) error {
	if w.network != "tcp" && w.network != "tcp4" && w.network != "tcp6" && w.network != "unix" && w.network != "unixpacket" {
		return ErrInvalidNetwork
	}
	g.Network = w.network
	return nil
}
