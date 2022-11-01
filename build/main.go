// Build is the build pipeline for this repository.
package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/boot"
)

// Directories used in repository.
const (
	dirRoot  = "."
	dirBuild = "build"
)

func main() {
	goyek.SetDefault(all)
	boot.Main()
}
