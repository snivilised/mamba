package assist

type (
	// LongFlagName
	LongFlagName = string

	// ShortFlagName
	ShortFlagName = string

	// FlagDefinitions
	FlagDefinitions map[LongFlagName]ShortFlagName
	// ShortFlagOverride allows a caller to override a specific short flag name
	// at bind time, without affecting the package-level definitions. Use this
	// when two parameter sets in the same command have competing short codes.
	//
	//	fam.Native.BindAll(fam, store.WithShortFlag("files-glob", "F"))
	ShortFlagOverride func(FlagDefinitions)
)

// WithShortFlag returns a ShortFlagOverride that reassigns the short code
// for a given long flag name.
func WithShortFlag(long LongFlagName, short ShortFlagName) ShortFlagOverride {
	return func(defs FlagDefinitions) {
		defs[long] = short
	}
}

func (fd FlagDefinitions) Clone() FlagDefinitions {
	c := make(FlagDefinitions, len(fd))
	for k, v := range fd {
		c[k] = v
	}
	return c
}

func MergeFlagDefinitions(families ...FlagDefinitions) FlagDefinitions {
	merged := make(FlagDefinitions)
	for _, family := range families {
		for long, short := range family {
			merged[long] = short
		}
	}
	return merged
}
