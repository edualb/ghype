package ghype

import (
	"net"

	"github.com/edualb/ghype/internal"
	"github.com/edualb/ghype/option"
)

type GHype interface {
	// Serve is responsable to start a new GHype server.
	Serve() error
	// Close is responsable to close a GHype server.
	Close() error
}

type GHypeImpl struct {
	settings *internal.GHypeSettings
	listener net.Listener
}

// New creates GHype struct responsable for start and close a server.
//
// It is possible set a lot of configurations using github.com/edualb/ghype/option package.
//
// EX: WithNetwork, WithAddress...
func New(o ...option.GHypeOptions) (GHype, error) {
	g := &GHypeImpl{
		settings: &internal.GHypeSettings{
			Address: ":8080",
			Network: "tcp",
		},
	}

	for _, opt := range o {
		err := opt.Apply(g.settings)
		if err != nil {
			return nil, err
		}
	}

	return g, nil
}

func (g *GHypeImpl) Serve() error {
	err := g.listen()
	if err != nil {
		return err
	}

	for {
		_, err := g.listener.Accept()
		if err != nil {
			return err
		}
	}
}

func (g *GHypeImpl) Close() error {
	if g.listener == nil {
		return ErrServerNotStarted
	}
	return g.listener.Close()
}

func (g *GHypeImpl) listen() error {
	ln, err := net.Listen(g.settings.Network, g.settings.Address)
	if err != nil {
		return err
	}
	g.listener = ln
	return nil
}
