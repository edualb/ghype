package ghype

import (
	"net"
)

type GHype interface {
	Serve(string) error
	Close() error
}

type GHypeImpl struct {
	listener net.Listener
}

func New() GHype {
	return &GHypeImpl{}
}

func (g *GHypeImpl) Serve(port string) error {
	err := g.listen(port)
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
	return g.listener.Close()
}

func (g *GHypeImpl) listen(port string) error {
	if len(port) <= 0 {
		ln, err := net.Listen("tcp", ":8080")
		if err != nil {
			return err
		}
		g.listener = ln
		return nil
	}

	ln, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	g.listener = ln
	return nil
}
