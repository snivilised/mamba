package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/mamba/assist"
	"github.com/snivilised/mamba/locale"
	"github.com/spf13/pflag"
)

type TextualInteractionParameterSet struct {
	ParameterSetWithOverrides
	IsNoTui bool
}

func (f *TextualInteractionParameterSet) BindAll(
	parent *assist.ParamSet[TextualInteractionParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --no-tui
	//
	const (
		defNoTUI = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			li18ngo.Text(locale.TextualInteractionIsNoTUIUsageTemplData{}),
			defNoTUI,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.IsNoTui,
	)
}

type CliInteractionParameterSet struct {
	ParameterSetWithOverrides
	IsTUI bool
}

func (f *CliInteractionParameterSet) BindAll(
	parent *assist.ParamSet[CliInteractionParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --tui
	//
	const (
		defIsTUI = false
	)

	parent.BindBool(
		resolveNewFlagInfo(
			li18ngo.Text(locale.CliInteractionIsTUIUsageTemplData{}),
			defIsTUI,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.IsTUI,
	)
}
