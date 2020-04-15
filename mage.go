// +build ignore

package main

import (
	"os"

	"github.com/magefile/mage/mage"
)

// Zero install option.
// Usage example:
//   go run mage.go test
func main() { os.Exit(mage.Main()) }
