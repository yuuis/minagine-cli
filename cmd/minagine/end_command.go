package main

import (
	"fmt"

	"github.com/mitchellh/cli"
)

type EndCommand struct {
	Token   string
	Project string
}

func EndCommandFactory(token, project string) func() (cli.Command, error) {
	return func() (cli.Command, error) {
		return &EndCommand{Token: token, Project: project}, nil
	}
}

func (*EndCommand) Synopsis() string {
	return "finish working"
}
func (*EndCommand) Help() string {
	return "usage: minagine end [option]"
}

func (c *EndCommand) Run(args []string) int {
	fmt.Println("finish working...")
	defer fmt.Println("done finish working.")

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

	if err := Post("end", project, token); err != nil {
		fmt.Println("failed finish working")
	}

	return 0
}
