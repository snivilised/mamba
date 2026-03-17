package store

import (
	"strings"

	"github.com/snivilised/mamba/assist"
	"github.com/spf13/pflag"
)

// A note to remember: when flags are bound in

var (
	workerPoolFlags = assist.FlagDefinitions{
		"cpu": "",
		"now": "",
	}

	previewFlags = assist.FlagDefinitions{
		"dry-run": "D",
	}

	// NB: files and files-glob can both use f as a short code,
	// because only one of them will be active in a CLI. It
	// doesn't makes sense to define both. Since the extended glob
	// filter is a very useful feature, it is being established as
	// the default, hence the name just 'files'.
	filterFlags = assist.FlagDefinitions{
		"files":         "f",
		"files-glob":    "b",
		"files-regex":   "x",
		"folders-glob":  "g",
		"folders-regex": "y",
	}

	profileFlags = assist.FlagDefinitions{
		"profile": "P",
		"scheme":  "S",
	}

	i18nFlags = assist.FlagDefinitions{
		"language": "",
	}

	depthFlags = assist.FlagDefinitions{
		"depth":      "",
		"no-recurse": "N",
	}

	samplingFlags = assist.FlagDefinitions{
		"sample":     "",
		"no-files":   "",
		"no-folders": "",
		"last":       "",
	}

	interactionFlags = assist.FlagDefinitions{
		"tui": "",
	}

	// ShortFlags is the unified lookup, composed from all families.
	ShortFlags = mergeFlagDefinitions(
		workerPoolFlags,
		previewFlags,
		filterFlags,
		profileFlags,
		i18nFlags,
		depthFlags,
		samplingFlags,
		interactionFlags,
	)
)

func mergeFlagDefinitions(families ...assist.FlagDefinitions) assist.FlagDefinitions {
	merged := make(assist.FlagDefinitions)
	for _, family := range families {
		for long, short := range family {
			merged[long] = short
		}
	}
	return merged
}

func longName(usage string) assist.LongFlagName {
	return strings.Split(usage, " ")[0]
}

// resolveNewFlagInfo resolves a FlagInfo for the given usage string and default
// value. Pass ShortFlagOverride values to reassign short codes locally without
// affecting the package-level ShortFlags map.
// Remember, when flags are bound in using BindString for example, the
// name of the flag is derived as the first token as retrieved from the
// i18n text definition. So if flag fails a unit test, then it possibly
// because of the text defined in locale.
func resolveNewFlagInfo[T any](usage string, defaultValue T,
	overrides []assist.ShortFlagOverride,
	alternativeFlagSet ...*pflag.FlagSet,
) *assist.FlagInfo {
	name := longName(usage)
	defs := ShortFlags.Clone()

	for _, o := range overrides {
		o(defs)
	}

	short := defs[name]

	if len(alternativeFlagSet) > 0 {
		return assist.NewFlagInfoOnFlagSet(usage, short, defaultValue, alternativeFlagSet[0])
	}

	return assist.NewFlagInfo(usage, short, defaultValue)
}
