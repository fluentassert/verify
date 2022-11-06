package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var mod = goyek.Define(goyek.Task{
	Name:  "mod",
	Usage: "go mod tidy",
	Action: func(a *goyek.A) {
		cmd.Exec(a, "go mod tidy", cmd.Dir(dirRoot))
		cmd.Exec(a, "go mod tidy", cmd.Dir(dirBuild))
	},
})
