package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var mod = goyek.Define(goyek.Task{
	Name:  "mod",
	Usage: "go mod tidy",
	Action: func(tf *goyek.TF) {
		cmd.Exec(tf, "go mod tidy", cmd.Dir(dirRoot))
		cmd.Exec(tf, "go mod tidy", cmd.Dir(dirBuild))
	},
})
