package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/mamba/assist"
	"github.com/snivilised/mamba/locale"
	"github.com/spf13/pflag"
)

type WorkerPoolParameterSet struct {
	ParameterSetWithOverrides
	CPU       bool
	NoWorkers int
}

func (f *WorkerPoolParameterSet) BindAll(
	parent *assist.ParamSet[WorkerPoolParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --cpu(C)
	//
	const (
		defaultCPU = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			li18ngo.Text(locale.WorkerPoolCPUParamUsageTemplData{}),
			defaultCPU,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.CPU,
	)

	// --now(N)
	//
	const (
		defaultNoW = -1
		minNow     = -1
		maxNow     = 100
	)

	parent.BindInt(
		resolveNewFlagInfo(
			li18ngo.Text(locale.WorkerPoolNoWParamUsageTemplData{}),
			defaultNoW,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.NoWorkers,
	)

	parent.Command.MarkFlagsMutuallyExclusive("cpu", "now")
}
