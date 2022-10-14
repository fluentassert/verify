package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/goyek/goyek/v2"
)

const (
	rootDir  = "."
	buildDir = "build"
	toolsDir = "tools"
)

func configure() {
	flow.SetDefault(flow.Define(goyek.Task{
		Name:  "all",
		Usage: "build pipeline",
		Deps: goyek.Deps{
			flow.Define(goyek.Task{
				Name:  "mod",
				Usage: "go mod tidy",
				Action: func(tf *goyek.TF) {
					Exec(tf, rootDir, "go mod tidy")
					Exec(tf, buildDir, "go mod tidy")
					Exec(tf, toolsDir, "go mod tidy")
				},
			}),
			flow.Define(goyek.Task{
				Name:  "install",
				Usage: "go install tools",
				Action: func(tf *goyek.TF) {
					tools := &strings.Builder{}
					toolsCmd := tf.Cmd("go", "list", `-f={{ join .Imports " " }}`, "-tags=tools")
					toolsCmd.Dir = toolsDir
					toolsCmd.Stdout = tools
					if err := toolsCmd.Run(); err != nil {
						tf.Fatal(err)
					}
					Exec(tf, toolsDir, "go install "+strings.TrimSpace(tools.String()))
				},
			}),
			flow.Define(goyek.Task{
				Name:  "golint",
				Usage: "golangci-lint",
				Action: func(tf *goyek.TF) {
					Exec(tf, rootDir, "golangci-lint run --fix")
					Exec(tf, buildDir, "golangci-lint run --fix")
				},
			}),
			flow.Define(goyek.Task{
				Name:  "spell",
				Usage: "misspell",
				Action: func(tf *goyek.TF) {
					Exec(tf, rootDir, "misspell -error -locale=US -i=importas -w .")
				},
			}),
			flow.Define(goyek.Task{
				Name:  "mdlint",
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
					Exec(tf, rootDir, fmt.Sprintf("docker build -t %s -f %s/markdownlint-cli.dockerfile .", toolsDir, dockerTag))
					Exec(tf, rootDir, fmt.Sprintf("docker run --rm -v '%s:/workdir' %s *.md", curDir, dockerTag))
				},
			}),
			flow.Define(goyek.Task{
				Name:  "test",
				Usage: "go test with code covarage",
				Action: func(tf *goyek.TF) {
					Exec(tf, rootDir, "go test -covermode=atomic -coverprofile=coverage.out ./...")
				},
			}),
		},
	}))
}
