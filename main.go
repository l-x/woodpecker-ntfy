package main

import (
	"fmt"
	"io"
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

func createRequest() (*http.Request, error) {
	req, err := http.NewRequest(
		"PUT",
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

func notify() error {
	req, err := createRequest()
	if err != nil {
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("%s %s", res.Status, b)
	}

	log.Printf("%s %s", res.Status, b)

	return nil
}

func main() {
	if err := notify(); err != nil {
		log.Fatal(err)
	}
}
