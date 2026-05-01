package locale

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

// FilesGlobParamUsageTemplData
// 🧊
type FilesGlobParamUsageTemplData struct {
	mambaTemplData
}

func (td FilesGlobParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "files-glob-filter.param-usage",
		Description: "files glob filter (negate-able with leading !)",
		Other:       "files-glob files glob filter (negate-able with leading !)",
	}
}

// FilesRegExParamUsageTemplData
// 🧊
type FilesRegExParamUsageTemplData struct {
	mambaTemplData
}

func (td FilesRegExParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "files-regex-filter.param-usage",
		Description: "files regex filter (negate-able with leading !)",
		Other:       "files-regex folder regular expression filter (negate-able with leading !)",
	}
}

// FilesExGlobParamUsageTemplData
// 🧊
type FilesExGlobParamUsageTemplData struct {
	mambaTemplData
}

func (td FilesExGlobParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "files-ex-glob-filter.param-usage",
		Description: "files extended glob filter (negate-able with leading !)",
		Other:       "files extended glob filter: <glob>|<suffixes csv> (negate-able with leading !)",
	}
}

// DirectoriesExGlobParamUsageTemplData
// 🧊
type DirectoriesExGlobParamUsageTemplData struct {
	mambaTemplData
}

func (td DirectoriesExGlobParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "directories-ex-glob-filter.param-usage",
		Description: "directories extended glob filter (negate-able with leading !)",
		Other:       "directories extended glob filter: <glob> (negate-able with leading !)",
	}
}

// DirectoriesGlobParamUsageTemplData
// 🧊
type DirectoriesGlobParamUsageTemplData struct {
	mambaTemplData
}

func (td DirectoriesGlobParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "dirs-glob-filter.param-usage",
		Description: "directories glob (negate-able with leading !)",
		Other:       "dirs-glob folder glob filter (negate-able with leading !)",
	}
}

// DirectoriesRexExParamUsageTemplData
// 🧊
type DirectoriesRexExParamUsageTemplData struct {
	mambaTemplData
}

func (td DirectoriesRexExParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "dirs-regex-filter.param-usage",
		Description: "directories regex filter (negate-able with leading !)",
		Other:       "dirs-regex folder regular expression filter (negate-able with leading !)",
	}
}

// WorkerPoolCPUParamUsageTemplData
// 🧊
type WorkerPoolCPUParamUsageTemplData struct {
	mambaTemplData
}

func (td WorkerPoolCPUParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "worker-pool-cpu.param-usage",
		Description: "run with the number of workers in pool set to number of CPUs available",
		Other:       "cpu denotes parallel execution with all available processors",
	}
}

// WorkerPoolCPUParamUsageTemplData
// 🧊
type WorkerPoolNoWParamUsageTemplData struct {
	mambaTemplData
}

func (td WorkerPoolNoWParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "worker-pool-cpu.param-usage",
		Description: "run with the number of workers in pool set to this number",
		Other:       "now denotes parallel execution with this number of workers in pool",
	}
}

// ProfileParamUsageTemplData
// 🧊
type ProfileParamUsageTemplData struct {
	mambaTemplData
}

func (td ProfileParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "profile.param-usage",
		Description: "pre-defined flag/option list in config file",
		Other:       "profile specifies which set of flags/options to load from config",
	}
}

// SchemeParamUsageTemplData
// 🧊
type SchemeParamUsageTemplData struct {
	mambaTemplData
}

func (td SchemeParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "scheme.param-usage",
		Description: "scheme is a collection of profiles, typically to repeat an operation over",
		Other:       "scheme is a collection of profiles, typically to repeat an operation over",
	}
}

// DryRunParamUsageTemplData
// 🧊
type DryRunParamUsageTemplData struct {
	mambaTemplData
}

func (td DryRunParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "dry-run.param-usage",
		Description: "allows the user to preview the effects of a command without running it",
		Other:       "dry-run allows the user to see the effects of a command without running it",
	}
}

// LanguageParamUsageTemplData
// 🧊
type LanguageParamUsageTemplData struct {
	mambaTemplData
}

func (td LanguageParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "language.param-usage",
		Description: "language allows the user to override the language the app runs in",
		Other:       "language allows the user to override the language the app runs in",
	}
}

// CascadeDepthParamUsageTemplData
// 🧊
type CascadeDepthParamUsageTemplData struct {
	mambaTemplData
}

func (td CascadeDepthParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "cascade-depth.param-usage",
		Description: "limits the number of sub directories navigated",
		Other:       "depth denotes the number of sub directories to navigate",
	}
}

// CascadeNoRecurseParamUsageTemplData
// 🧊
type CascadeNoRecurseParamUsageTemplData struct {
	mambaTemplData
}

func (td CascadeNoRecurseParamUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "cascade-no-recurse.param-usage",
		Description: "sets the navigator to not descend into sub directories",
		Other:       "no-recurse sets the navigator to not descend into sub directories",
	}
}

// SamplingSampleUsageTemplData
// 🧊
type SamplingSampleUsageTemplData struct {
	mambaTemplData
}

func (td SamplingSampleUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "sampling-sample.param-usage",
		Description: "sampling sample usage; activates sampling",
		Other:       "sample is a flag that activates sampling",
	}
}

// SamplingNoFilesUsageTemplData
// 🧊
type SamplingNoFilesUsageTemplData struct {
	mambaTemplData
}

func (td SamplingNoFilesUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "sampling-no-files.param-usage",
		Description: "sampling files usage; no of files in sample set",
		Other:       "no-files specifies the number of files to sample",
	}
}

// SamplingNoDirectoriesUsageTemplData
// 🧊
type SamplingNoDirectoriesUsageTemplData struct {
	mambaTemplData
}

func (td SamplingNoDirectoriesUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "sampling-no-dirs.param-usage",
		Description: "sampling directories usage; no of directories in sample set",
		Other:       "no-dirs specifies the number of directories to sample",
	}
}

// SamplingLastUsageTemplData
// 🧊
type SamplingLastUsageTemplData struct {
	mambaTemplData
}

func (td SamplingLastUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "sampling-last.param-usage",
		Description: "sampling last usage; indicates which n items are to be sampled",
		Other:       "last is a flag that indicates last n items are to be sampled instead of the first",
	}
}

// TextualInteractionIsNoTUIUsageTemplData
// 🧊
type TextualInteractionIsNoTUIUsageTemplData struct {
	mambaTemplData
}

func (td TextualInteractionIsNoTUIUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "textual-interaction-is-no-tui.param-usage",
		Description: "textual interaction is no-tui usage; deactivates tui mode",
		Other:       "no-tui is a flag that turns off tui mode",
	}
}

// CliInteractionIsTUIUsageTemplData
// 🧊
type CliInteractionIsTUIUsageTemplData struct {
	mambaTemplData
}

func (td CliInteractionIsTUIUsageTemplData) Message() *i18n.Message {
	return &i18n.Message{
		ID:          "cli-interaction-is-tui.param-usage",
		Description: "tui interaction is tui usage; activates tui mode",
		Other:       "tui is a flag that enables tui mode",
	}
}
