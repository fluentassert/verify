package main

import (
	"github.com/goyek/goyek/v2"
	"github.com/mattn/go-shellwords"
)

// Exec runs the provided command line.
// It fails the task in case of any problems.
func Exec(tf *goyek.TF, workDir string, cmdLine string) {
	tf.Logf("Run %q in %s", cmdLine, workDir)
	args, err := shellwords.Parse(cmdLine)
	if err != nil {
		tf.Fatalf("parse command line: %v", err)
	}
	cmd := tf.Cmd(args[0], args[1:]...)
	cmd.Dir = workDir
	if err := cmd.Run(); err != nil {
		tf.Error(err)
	}
}
