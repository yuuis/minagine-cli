package main

import (
	"log"
	"os"

	"github.com/mitchellh/cli"
)

func main() {

	c := cli.NewCLI("minagine", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"start": StartCommandFactory,
		"end":   EndCommandFactory,
	}

	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}

	os.Exit(exitStatus)
}
