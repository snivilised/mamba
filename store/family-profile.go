package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/mamba/assist"
	"github.com/snivilised/mamba/locale"
	"github.com/spf13/pflag"
)

type ProfileParameterSet struct {
	ParameterSetWithOverrides
	Profile string
	Scheme  string
}

func (f *ProfileParameterSet) BindAll(
	parent *assist.ParamSet[ProfileParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --profile(P)
	//
	const (
		defaultProfile = ""
	)

	// should match: `^[\w-]+$`,
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.ProfileParamUsageTemplData{}),
			defaultProfile,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.Profile,
	)

	// -- scheme(S)
	//
	const (
		defaultScheme = ""
	)

	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.SchemeParamUsageTemplData{}),
			defaultScheme,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.Scheme,
	)
}
