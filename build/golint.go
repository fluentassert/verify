package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var golint = goyek.Define(goyek.Task{
	Name:  "golint",
	Usage: "golangci-lint run --fix",
	Action: func(tf *goyek.TF) {
		if !cmd.Exec(tf, "go install github.com/golangci/golangci-lint/cmd/golangci-lint", cmd.Dir(dirBuild)) {
			return
		}
		cmd.Exec(tf, "golangci-lint run --fix", cmd.Dir(dirRoot))
		cmd.Exec(tf, "golangci-lint run --fix", cmd.Dir(dirBuild))
	},
})
