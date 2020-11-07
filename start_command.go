package minaginecli

import (
	"flag"
	"fmt"
	"os"

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
		fmt.Println("failed start working")
	}

	return 0
}
