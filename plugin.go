package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const defaultNtfyUrl = "https://ntfy.sh/woodpecker-ntfy"

var (
	server = setting{
		valueFn:    env("PLUGIN_URL"),
		fallbackFn: func() string { return defaultNtfyUrl }}
	message = setting{valueFn: env("PLUGIN_MESSAGE")}
	headers = map[string]setting{
		"Authorization": {valueFn: getAuth},
		"Title":         {valueFn: env("PLUGIN_TITLE")},
		"Priority":      {valueFn: env("PLUGIN_PRIORITY")},
		"Tags":          {valueFn: env("PLUGIN_TAGS")},
		"Actions":       {valueFn: env("PLUGIN_ACTIONS")},
		"Click":         {valueFn: env("PLUGIN_CLICK"), fallbackFn: env("CI_BUILD_LINK")},
		"Icon":          {valueFn: env("PLUGIN_ICON"), fallbackFn: env("CI_COMMIT_AUTHOR_AVATAR")},
	}
	failOnError = setting{valueFn: env("PLUGIN_FAIL_ON_ERROR")}
)

type setting struct {
	valueFn    func() string
	fallbackFn func() string
}

func (h *setting) getValue() string {
	value := h.valueFn()

	if value == "" && h.fallbackFn != nil {
		return h.fallbackFn()
	}

	return value
}

func env(key string) func() string {
	return func() string {
		return os.Getenv(key)
	}
}

func getAuth() string {
	if token := os.Getenv("PLUGIN_TOKEN"); token != "" {
		return fmt.Sprintf("Bearer %s", token)
	}

	return ""
}

func checkErr(err error) {
	if err == nil {
		return
	}

	if failOnError.getValue() == "" {
		log.Println(err)
	} else {
		log.Fatalln(err)
	}
}

func createRequest() (*http.Request, error) {
	req, err := http.NewRequest(
		"POST",
		server.getValue(),
		strings.NewReader(message.getValue()),
	)
	if err != nil {
		return req, err
	}

	for k, v := range headers {
		req.Header.Set(k, v.getValue())
	}

	return req, nil
}
