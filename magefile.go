// +build mage

package main

import (
	"github.com/magefile/mage/sh"
	"github.com/mysteriumnetwork/go-ci/commands"
)

// CheckCopyright checks for copyright headers in files.
func CheckCopyright() error {
	return commands.CopyrightD(".", "pb")
}

// Generate generates the bindings.
func Generate() error {
	return sh.RunV("go", "generate")
}


