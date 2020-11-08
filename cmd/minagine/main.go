package main

import (
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("minagine", "1.0.0")
	c.Args = os.Args[1:]

	envToken, envProject := GetTokenAndProjectFromEnvVariables()

	c.Commands = map[string]cli.CommandFactory{
		"start": StartCommandFactory(envToken, envProject),
		"end":   EndCommandFactory(envToken, envProject),
	}

	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(exitStatus)
}
