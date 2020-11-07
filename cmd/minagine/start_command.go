package main

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type StartCommand struct{}

func StartCommandFactory() (cli.Command, error) {
	return &StartCommand{}, nil
}

func (*StartCommand) Synopsis() string {
	return "start working"
}
func (*StartCommand) Help() string {
	return "usage: minagine start [option]"
}

func (*StartCommand) Run(args []string) int {
	fmt.Println("start working...")
	defer fmt.Println("done start working!")

	envToken, envProject := GetTokenAndProjectFromEnvVariables()

	token, project, err := GetTokenAndProjectFromFlags(args, envToken, envProject)
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
