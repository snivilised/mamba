package lab

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// Path joins the relative path.
func Path(parent, relative string) string {
	segments := strings.Split(relative, "/")
	return filepath.Join(append([]string{parent}, segments...)...)
}

// Normalise converts a path to use the correct separators for the native
// file system.
func Normalise(p string) string {
	return strings.ReplaceAll(p, "/", string(filepath.Separator))
}

// Reason creates a standardised reason message for error messages.
func Reason(name string) string {
	return fmt.Sprintf("❌ for item named: '%v'", name)
}

// JoinCwd
func JoinCwd(segments ...string) string {
	if current, err := os.Getwd(); err == nil {
		parent, _ := filepath.Split(current)
		grand := filepath.Dir(parent)
		great := filepath.Dir(grand)
		all := append([]string{great}, segments...)

		return filepath.Join(all...)
	}

	panic("could not get root path")
}

// Root
func Root() string {
	if current, err := os.Getwd(); err == nil {
		return current
	}

	panic("could not get root path")
}

// Combine creates a path from the parent combined with the relative path. The relative
// path is a file system path so should only contain forward slashes, not the standard
// file path separator as denoted by filepath.Separator, typically used when interacting
// with the local file system. Do not use trailing "/".
func Combine(parent, relative string) string {
	if relative == "" {
		return parent
	}

	return parent + "/" + relative
}

// Repo gets the root path of the repo using git and joins the relative path
func Repo(relative string) string {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, _ := cmd.Output()

	repo := strings.TrimSpace(string(output))

	return Combine(repo, relative)
}

// Log get the log path
func Log() string {
	repo := Repo("")

	return filepath.Join(repo, "test", "test.log")
}
