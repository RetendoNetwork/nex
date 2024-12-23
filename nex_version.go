package nex

import (
	"fmt"
)

type nexVersion struct {
	Maj int
	Min int
	Pat int
	Gsp string
	Smv string
}

// For get the NEX version
func (n *nexVersion) Copy() *nexVersion {
	return &nexVersion{
		Maj: n.Maj,
		Min: n.Min,
		Pat: n.Pat,
		Gsp: n.Gsp,
		Smv: fmt.Sprintf("v%d.%d.%d", n.Maj, n.Min, n.Pat),
	}
}

// For set the NEX version
func NewNexVersion(maj, min, pat int) *nexVersion {
	return &nexVersion{
		Maj: maj,
		Min: min,
		Pat: pat,
		Smv: fmt.Sprintf("v%d.%d.%d", maj, min, pat),
	}
}
