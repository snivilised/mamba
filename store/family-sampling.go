package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/mamba/assist"
	"github.com/snivilised/mamba/locale"
	"github.com/spf13/pflag"
)

type SamplingParameterSet struct {
	ParameterSetWithOverrides
	IsSampling    bool
	NoFiles       uint
	NoDirectories uint
	Last          bool
}

func (f *SamplingParameterSet) BindAll(
	parent *assist.ParamSet[SamplingParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	const (
		defIsSampling = false
	)

	// --sample
	//
	parent.BindBool(
		resolveNewFlagInfo(
			li18ngo.Text(locale.SamplingSampleUsageTemplData{}),
			defIsSampling,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.IsSampling,
	)

	const (
		defFSItems = uint(3)
		minFSItems = uint(1)
		maxFSItems = uint(128)
	)

	// --no-files
	//
	parent.BindUint(
		resolveNewFlagInfo(
			li18ngo.Text(locale.SamplingNoFilesUsageTemplData{}),
			defFSItems,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.NoFiles,
	)

	// --no-dirs
	//
	parent.BindUint(
		resolveNewFlagInfo(
			li18ngo.Text(locale.SamplingNoDirectoriesUsageTemplData{}),
			defFSItems,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.NoDirectories,
	)

	const (
		defIsLast = false
	)

	// --last
	//
	parent.BindBool(
		resolveNewFlagInfo(
			li18ngo.Text(locale.SamplingLastUsageTemplData{}),
			defIsLast,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.Last,
	)
}
