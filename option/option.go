package option

import (
	"github.com/edualb/ghype/internal"
)

type GHypeOptions interface {
	Apply(g *internal.GHypeSettings) error
}
