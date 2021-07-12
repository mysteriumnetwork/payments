// +build mage

package main

import (
	"strings"

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

func GenerateMatic() error {
	command := `run bindings/abi/abigen.go --localdir=./bindings/matic/abi --contracts=ChildContract.json,RootChainManager.json,RootChain.json --out=bindings/matic --pkg=matic`
	return sh.RunV("go", strings.Split(command, " ")...)
}

func Test() error {
	return sh.RunV("go", "test", "--short", "-race", "-cover", "./...")
}
