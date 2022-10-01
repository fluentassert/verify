// Build is the build pipeline for this repository.
package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/goyek/goyek"
	"github.com/mattn/go-shellwords"
)

const buildDir = "build"

func main() {
	flow().Main()
}

func flow() *goyek.Flow {
	flow := &goyek.Flow{}

	test := flow.Register(taskTest())
	lint := flow.Register(taskLint())
	misspell := flow.Register(taskMisspell())
	markdownlint := flow.Register(taskMarkdownLint())
	all := flow.Register(taskAll(goyek.Deps{
		test, lint, misspell, markdownlint,
	}))

	flow.DefaultTask = all
	return flow
}

func taskTest() goyek.Task {
	return goyek.Task{
		Name:  "test",
		Usage: "go test with code covarage",
		Action: func(tf *goyek.TF) {
			Exec(tf, buildDir, "go test")
			Exec(tf, "", "go test -covermode=atomic -coverprofile=coverage.out ./...")
		},
	}
}

func taskLint() goyek.Task {
	return goyek.Task{
		Name:  "lint",
		Usage: "golangci-lint",
		Action: func(tf *goyek.TF) {
			Exec(tf, buildDir, "go install github.com/golangci/golangci-lint/cmd/golangci-lint")
			Exec(tf, buildDir, "golangci-lint run")
			Exec(tf, "", "golangci-lint run")
		},
	}
}

func taskMisspell() goyek.Task {
	return goyek.Task{
		Name:  "misspell",
		Usage: "misspell",
		Action: func(tf *goyek.TF) {
			Exec(tf, buildDir, "go install github.com/client9/misspell/cmd/misspell")
			Exec(tf, "", "misspell -error -locale=US *.md")
		},
	}
}

func taskMarkdownLint() goyek.Task {
	return goyek.Task{
		Name:  "markdownlint",
		Usage: "markdownlint-cli (uses docker)",
		Action: func(tf *goyek.TF) {
			if _, err := exec.LookPath("docker"); err != nil {
				tf.Skip(err)
			}

			curDir, err := os.Getwd()
			if err != nil {
				tf.Fatal(err)
			}
			dockerTag := "markdownlint-cli"
			buildCmd := fmt.Sprintf("docker build -t %s -f build/markdownlint-cli.dockerfile .", dockerTag)
			Exec(tf, "", buildCmd)
			runCmd := fmt.Sprintf("docker run --rm -v %s:/workdir %s *.md", curDir, dockerTag)
			Exec(tf, "", runCmd)
		},
	}
}

func taskAll(deps goyek.Deps) goyek.Task {
	return goyek.Task{
		Name:  "all",
		Usage: "build pipeline",
		Deps:  deps,
	}
}

// Exec runs the provided command line.
// It fails the task in case of any problems.
func Exec(tf *goyek.TF, workDir string, cmdLine string) {
	args, err := shellwords.Parse(cmdLine)
	if err != nil {
		tf.Fatalf("parse command line: %v", err)
	}
	cmd := tf.Cmd(args[0], args[1:]...)
	cmd.Dir = workDir
	if err := cmd.Run(); err != nil {
		tf.Fatalf("run command: %v", err)
	}
}
