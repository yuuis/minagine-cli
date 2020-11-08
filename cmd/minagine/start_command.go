package main

import (
	"fmt"
	"github.com/mitchellh/cli"
)

type StartCommand struct {
	Token   string
	Project string
}

func StartCommandFactory(token, project string) func() (cli.Command, error) {
	return func() (cli.Command, error) {
		return &StartCommand{Token: token, Project: project}, nil
	}
}

func (*StartCommand) Synopsis() string {
	return "start working"
}
func (*StartCommand) Help() string {
	return "usage: minagine start [option]"
}

func (c *StartCommand) Run(args []string) int {
	fmt.Println("start working...")
	defer fmt.Println("done start working.")

	token, project, err := GetTokenAndProjectFromFlags(args, c.Token, c.Project)
	if err != nil {
		fmt.Printf("fail: parse argument, %s", err.Error())
		return 1
	}

	if token == "" {
		fmt.Println("token must be specified by argument or environment variable")
		return 1
	}

	if project == "" {
		fmt.Println("project must be specified by argument or environment variable")
		return 1
	}

	if err := Post("start", project, token); err != nil {
		fmt.Println("failed start working")
	}

	return 0
}
