package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/mamba/assist"
	"github.com/snivilised/mamba/locale"
	"github.com/spf13/pflag"
)

type PreviewParameterSet struct {
	ParameterSetWithOverrides
	DryRun bool
}

func (f *PreviewParameterSet) BindAll(
	parent *assist.ParamSet[PreviewParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --dry-run(D)
	//
	const (
		defaultDryRun = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			li18ngo.Text(locale.DryRunParamUsageTemplData{}),
			defaultDryRun,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.DryRun,
	)
}
