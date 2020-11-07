package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

const functionURL = "https://%s.cloudfunctions.net/minagine/%s"

func Post(workType, project, token string) error {

	if workType != "start" && workType != "end" {
		return fmt.Errorf("you specified type %s is invalid. \"start\" or \"end\" are valid type", workType)
	}

	if token == "" {
		return errors.New("token must not be null")
	}

	req, err := http.NewRequest(
		"POST",
		fmt.Sprintf(functionURL, project, workType),
		strings.NewReader(url.Values{}.Encode()),
	)

	if err != nil {
		return err
	}

	req.Header.Set("X-TOKEN", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return err
}
