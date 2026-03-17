package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/mamba/assist"
	"github.com/snivilised/mamba/locale"
	"github.com/spf13/pflag"
)

type I18nParameterSet struct {
	ParameterSetWithOverrides
	Language string
}

func (f *I18nParameterSet) BindAll(
	parent *assist.ParamSet[I18nParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --language
	//
	const (
		defaultLanguage = ""
	)

	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.LanguageParamUsageTemplData{}),
			defaultLanguage,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.Language,
	)
}
