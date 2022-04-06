package main

import (
	"flag"
	"os"
)

const (
	tokenEnvVariable   = "MINAGINE_TOKEN"
	projectEnvVariable = "MINAGINE_PROJECT"
)

func GetTokenAndProjectFromEnvVariables() (string, string) {
	return os.Getenv(tokenEnvVariable), os.Getenv(projectEnvVariable)
}

func GetTokenAndProjectFromFlags(args []string, defaultToken, defaultProject string) (string, string, error) {
	var token, project string

	flag.StringVar(&token, "token", defaultToken, "token for authorization")
	flag.StringVar(&project, "project", defaultProject, "project target function URL")

	if err := flag.CommandLine.Parse(args); err != nil {
		return "", "", err
	}

	return token, project, nil
}
