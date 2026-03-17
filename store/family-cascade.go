package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/mamba/assist"
	"github.com/snivilised/mamba/locale"

	"github.com/spf13/pflag"
)

type CascadeParameterSet struct {
	ParameterSetWithOverrides
	Depth     uint
	NoRecurse bool
}

func (f *CascadeParameterSet) BindAll(
	parent *assist.ParamSet[CascadeParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --depth
	//
	const (
		defaultDepth = uint(0)
	)

	parent.BindUint(
		resolveNewFlagInfo(
			li18ngo.Text(locale.CascadeDepthParamUsageTemplData{}),
			defaultDepth,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.Depth,
	)

	// --no-recurse(N)
	//
	const (
		defaultNoRecurse = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			li18ngo.Text(locale.CascadeNoRecurseParamUsageTemplData{}),
			defaultNoRecurse,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.NoRecurse,
	)

	parent.Command.MarkFlagsMutuallyExclusive("depth", "no-recurse")
}
