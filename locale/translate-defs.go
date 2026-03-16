package locale

// CLIENT-TODO: Should be updated to use url of the implementing project,
// so should not be left as mamba. (this should be set by auto-check)
const MambaSourceID = "github.com/snivilised/mamba"

type mambaTemplData struct{}

func (td mambaTemplData) SourceID() string {
	return MambaSourceID
}
