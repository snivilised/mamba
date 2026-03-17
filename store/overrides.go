package store

import (
	"github.com/snivilised/mamba/assist"
)

type (
	// ParameterSetWithOverrides allows the client to define short code overrides
	// applicable for the host parameter set.
	ParameterSetWithOverrides struct {
		Overrides []assist.ShortFlagOverride
	}
)
