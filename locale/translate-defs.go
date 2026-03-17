package locale

import (
	"github.com/snivilised/li18ngo"
)

// MambaSourceID is the source ID for mamba, to which i18n localisation
// messages are bound.
const (
	MambaSourceID = "github.com/snivilised/mamba"
)

// Use initialises i18n translation
func Use(options ...li18ngo.UseOptionFn) error {
	return li18ngo.Use(options...)
}

type mambaTemplData struct{}

// SourceID is the source ID for mamba, to which i18n localisation
// messages are bound.
func (td mambaTemplData) SourceID() string {
	return MambaSourceID
}
