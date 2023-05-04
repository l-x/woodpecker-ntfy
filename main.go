package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const defaultNtfyServer = "https://ntfy.sh"

type header struct {
	name         string
	valueFn      func() string
	fallbackFn   func() string
	includeEmpty bool
}

func (h *header) add(r *http.Header) {
	value := h.valueFn()
	if value == "" && h.fallbackFn != nil {
		value = h.fallbackFn()
	}

	if value == "" && !h.includeEmpty {
		return
	}

	r.Add(h.name, value)
}

var headers = []header{
	{name: "Authorization", valueFn: getAuth},
	{name: "Title", valueFn: env("PLUGIN_TITLE")},
	{name: "Priority", valueFn: env("PLUGIN_PRIORITY")},
	{name: "Tags", valueFn: env("PLUGIN_TAGS")},
	{name: "Actions", valueFn: env("PLUGIN_ACTIONS")},
	{name: "Click", valueFn: env("PLUGIN_CLICK")},
	{name: "Icon", valueFn: env("PLUGIN_ICON")},
}

func env(key string) func() string {
	return func() string {
		return os.Getenv(key)
	}
}

func getEnv(key, def string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}

	return def
}

func getServerUrl(defaultServer string) (*url.URL, error) {
	u, err := url.Parse(getEnv("PLUGIN_URL", defaultServer))
	if err != nil {
		return u, err
	}

	if topic := getEnv("PLUGIN_TOPIC", ""); topic != "" {
		return u.JoinPath(topic), nil
	}

	return u, fmt.Errorf("no topic configured")
}

func getAuth() string {
	if token := getEnv("PLUGIN_TOKEN", ""); token != "" {
		return fmt.Sprintf("Bearer %s", token)
	}

	return ""
}

func main() {
	serverUrl, err := getServerUrl(defaultNtfyServer)
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest(
		"POST",
		serverUrl.String(),
		strings.NewReader(getEnv("PLUGIN_MESSAGE", "")),
	)

	if err != nil {
		log.Fatalln(err)
	}

	for _, h := range headers {
		h.add(&req.Header)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	
	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}
	
	log.Printf(" <= %s", b)
}
