package store

import (
	"github.com/snivilised/li18ngo"
	"github.com/snivilised/mamba/assist"
	"github.com/snivilised/mamba/locale"
	"github.com/spf13/pflag"
)

const (
	defaultFilterValue = ""
)

// NB: We don't want to use up too many of the letters of the alphabet
// on short flag names, because it leaves less for the client to use.
// Therefore, for compound filters typically used when we want to filter
// file system nodes by file name and directory name, we forego the
// ability to specify compound file names (when using the navigator
// in traverse with the FoldersWithFiles subscription type) with a
// short code, as this is seen as a niche feature. The more common
// scenarios would be to either filter files, directories or both
// by using an 'any' scope. With this compromise, the user would
// always have to spell the compound file filter in it full form:
// --files-regex or --files-glob. When using traverse nav, the directories
// with files subscription would have to be used, ie there is no
// standalone file file, so --files-regex and --files-glob are both free
// to use without ambiguity.
// For a regular files scenario, we would need to use the files
// subscription type and in this case, --files-regex(x) and --files-glob(g)
// are still free to be used without ambiguity.

// FilesFilterParameterSet represents a family of parameters that can be used
// to accept a file filter. files is considered the default as it is
// the most user friendly to use, as a glob is easier and more intuitive
// to use on the command line and supports (with te help of a delimiter)
// multiple extensions to be specified with a csv, in contrast to a regular glob.
type FilesFilterParameterSet struct {
	ParameterSetWithOverrides
	Files      string
	FilesGlob  string
	FilesRegEx string
}

func (f *FilesFilterParameterSet) BindAll(
	parent *assist.ParamSet[FilesFilterParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --files(f)
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.FilesExGlobParamUsageTemplData{}),
			defaultFilterValue,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.Files,
	)

	// --files-glob(b)
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.FilesGlobParamUsageTemplData{}),
			defaultFilterValue,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.FilesGlob,
	)

	// --files-regex(x)
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.FilesRegExParamUsageTemplData{}),
			defaultFilterValue,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.FilesRegEx,
	)

	parent.Command.MarkFlagsMutuallyExclusive("files", "files-glob", "files-regex")
}

// DirectoriesFilterParameterSet represents a family of parameters that can be used
// to accept a folder filter. In contrast to files, the directories family does
// not include an extended glob because directories do not contain extensions,
// so the regular glob will suffice.
type DirectoriesFilterParameterSet struct {
	ParameterSetWithOverrides
	DirectoriesGlob  string
	DirectoriesRegEx string
}

func (f *DirectoriesFilterParameterSet) BindAll(
	parent *assist.ParamSet[DirectoriesFilterParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --dirs-glob(g)
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.DirectoriesGlobParamUsageTemplData{}),
			defaultFilterValue,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.DirectoriesGlob,
	)

	// --dirs-regex(y)
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.DirectoriesRexExParamUsageTemplData{}),
			defaultFilterValue,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.DirectoriesRegEx,
	)

	parent.Command.MarkFlagsMutuallyExclusive("dirs-glob", "dirs-regex")
}

// PolyFilterParameterSet represents a family of parameters that can be used
// to accept file and folder filters. This family is composed of files and
// filters. For files, either an extended glob or regex is supported. For
// directories, either a regular glob or regex is supported.
type PolyFilterParameterSet struct {
	ParameterSetWithOverrides
	FilesExGlob      string
	FilesRegEx       string
	DirectoriesGlob  string
	DirectoriesRegEx string
}

func (f *PolyFilterParameterSet) BindAll(
	parent *assist.ParamSet[PolyFilterParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --files-glob(b)
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.FilesExGlobParamUsageTemplData{}),
			defaultFilterValue,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.FilesExGlob,
	)

	// --files-regex(x)
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.FilesRegExParamUsageTemplData{}),
			defaultFilterValue,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.FilesRegEx,
	)

	// --dirs-glob(g)
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.DirectoriesGlobParamUsageTemplData{}),
			defaultFilterValue,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.DirectoriesGlob,
	)

	// --dirs-regex(y)
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.DirectoriesRexExParamUsageTemplData{}),
			defaultFilterValue,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.DirectoriesRegEx,
	)

	parent.Command.MarkFlagsMutuallyExclusive("files", "files-regex")
	parent.Command.MarkFlagsMutuallyExclusive("dirs-glob", "dirs-regex")
}

// AlloyFilterParameterSet represents a family of parameters that can be used
// to accept file and folder filters. Files are represented by an extended glob
// and directories by a regular glob.
type AlloyFilterParameterSet struct {
	ParameterSetWithOverrides
	FilesExGlob     string
	DirectoriesGlob string
}

func (f *AlloyFilterParameterSet) BindAll(
	parent *assist.ParamSet[AlloyFilterParameterSet],
	flagSet ...*pflag.FlagSet,
) {
	// --files(f)
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.FilesExGlobParamUsageTemplData{}),
			defaultFilterValue,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.FilesExGlob,
	)

	// --dirs-glob(g)
	//
	parent.BindString(
		resolveNewFlagInfo(
			li18ngo.Text(locale.DirectoriesGlobParamUsageTemplData{}),
			defaultFilterValue,
			f.Overrides,
			flagSet...,
		),
		&parent.Native.DirectoriesGlob,
	)
}
