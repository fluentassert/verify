// Build is the build pipeline for this repository.
package main

import (
	"fmt"
	"os"

	"github.com/goyek/goyek/v2"
)

var flow = &goyek.Flow{Verbose: true}

func main() {
	if err := os.Chdir(".."); err != nil {
		fmt.Println(err)
		os.Exit(goyek.CodeInvalidArgs)
	}
	configure()
	flow.Main(os.Args[1:])
}
