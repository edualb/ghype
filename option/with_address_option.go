package option

import "github.com/edualb/ghype/internal"

type withAddress struct{ address string }

// WithAddress is responsable to configure the address value in GHypeSettings.
//
// More information: https://tools.ietf.org/html/rfc2616
func WithAddress(addr string) GHypeOptions {
	return withAddress{
		address: addr,
	}
}

func (w withAddress) Apply(g *internal.GHypeSettings) error {
	g.Address = w.address
	return nil
}
