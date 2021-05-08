package main

import (
	"log"
	"os"

	"github.com/goyek/goyek"
)

func main() {
	if err := os.Chdir(".."); err != nil {
		log.Fatalln(err)
	}

	flow := &goyek.Taskflow{}

	test := flow.Register(taskTest())
	lint := flow.Register(taskLint())
	misspell := flow.Register(taskMisspell())
	all := flow.Register(taskAll(goyek.Deps{
		test, lint, misspell,
	}))

	flow.DefaultTask = all
	flow.Main()
}

const buildDir = "build"

func taskTest() goyek.Task {
	return goyek.Task{
		Name:    "test",
		Usage:   "go test with code covarage",
		Command: goyek.Exec("go", "test", "-covermode=atomic", "-coverprofile=coverage.out", "./..."),
	}
}

func taskLint() goyek.Task {
	return goyek.Task{
		Name:  "lint",
		Usage: "golangci-lint",
		Command: func(tf *goyek.TF) {
			installCmd := tf.Cmd("go", "install", "github.com/golangci/golangci-lint/cmd/golangci-lint")
			installCmd.Dir = buildDir
			if err := installCmd.Run(); err != nil {
				tf.Errorf("go install golangci-lint: %v", err)
			}
			cmd := tf.Cmd("golangci-lint", "run")
			if err := cmd.Run(); err != nil {
				tf.Errorf("golangci-lint run: %v", err)
			}
		},
	}
}

func taskMisspell() goyek.Task {
	return goyek.Task{
		Name:  "misspell",
		Usage: "misspell",
		Command: func(tf *goyek.TF) {
			installCmd := tf.Cmd("go", "install", "github.com/client9/misspell/cmd/misspell")
			installCmd.Dir = buildDir
			if err := installCmd.Run(); err != nil {
				tf.Errorf("go install misspell: %v", err)
			}
			cmd := tf.Cmd("misspell", "-error", "-locale=US", "*.md")
			if err := cmd.Run(); err != nil {
				tf.Errorf("misspell: %v", err)
			}
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
