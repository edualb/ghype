package ghype

import (
	"fmt"
	"testing"
)

func TestGHypeListen(t *testing.T) {
	ports := []string{
		"",
		":3000",
	}

	for _, port := range ports {
		g := &GHypeImpl{}
		err := g.listen(port)

		if err != nil {
			t.Error(err.Error())
		}

		if g.listener == nil {
			t.Error("Listener nil")
		}

		if g.listener.Addr().String() != fmt.Sprintf("127.0.0.1:%s", port) {
			t.Error(fmt.Sprintf("Wrong port: %s. %s", port, g.listener.Addr().String()))
		}

		if g.listener.Addr().Network() != "tcp" {
			t.Error("Wrong network")
		}
	}
}
