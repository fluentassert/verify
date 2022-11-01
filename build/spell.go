package main

import (
	"strings"

	"github.com/goyek/goyek/v2"
	"github.com/goyek/x/cmd"
)

var spell = goyek.Define(goyek.Task{
	Name:  "spell",
	Usage: "misspell",
	Action: func(tf *goyek.TF) {
		if !cmd.Exec(tf, "go install github.com/client9/misspell/cmd/misspell", cmd.Dir(dirBuild)) {
			return
		}
		mdFiles := find(tf, ".md")
		if len(mdFiles) == 0 {
			tf.Skip("no .md files")
		}
		cmd.Exec(tf, "misspell -error -locale=US -w "+strings.Join(mdFiles, " "))
	},
})
