// Build is the build pipeline for this repository.
package main

import (
	"os"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/goyek/v2/middleware"
)

// Directories used in repository.
const (
	dirRoot  = "."
	dirBuild = "build"
)

func main() {
	goyek.Use(middleware.ReportStatus)
	goyek.SetDefault(all)
	goyek.Main(os.Args[1:])
}
