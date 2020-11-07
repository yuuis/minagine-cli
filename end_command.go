package minaginecli

import (
	"flag"
	"fmt"
	"os"

	"github.com/mitchellh/cli"
)

type EndCommand struct{}

func EndCommandFactory() (cli.Command, error) {
	return &EndCommand{}, nil
}

func (*EndCommand) Synopsis() string {
	return "end working"
}
func (*EndCommand) Help() string {
	return "usage: minagine end [option]"
}

func (*EndCommand) Run(args []string) int {
	fmt.Println("finish running!")
	defer fmt.Println("done finish working!")

	envToken := os.Getenv("MINAGINE_TOKEN")
	envProject := os.Getenv("MINAGINE_PROJECT")

	var token string
	var project string
	flag.StringVar(&token, "token", envToken, "token for authorization")
	flag.StringVar(&project, "project", envProject, "project target function URL")

	if token == "" {
		fmt.Println("token must be specified by argument or environment variable")
		return 1
	}

	if project == "" {
		fmt.Println("project must be specified by argument or environment variable")
		return 1
	}

	if err := Post("start", project, token); err != nil {
		fmt.Println("failed finish working")
	}

	return 0
}
