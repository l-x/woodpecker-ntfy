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
var defaultIconFn = func() string { return "https://woodpecker-ci.org/img/logo.svg" }

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
	{name: "Click", valueFn: env("PLUGIN_CLICK"), fallbackFn: env("CI_BUILD_LINK")},
	{name: "Icon", valueFn: env("PLUGIN_ICON"), fallbackFn: defaultIconFn},
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

func getAuth() string {
	if token := getEnv("PLUGIN_TOKEN", ""); token != "" {
		return fmt.Sprintf("Bearer %s", token)
	}

	return ""
}

func checkErr(err error) {
	if err == nil {
		return
	}

	if getEnv("PLUGIN_FAIL_ON_ERROR", "") == "" {
		log.Println(err)
	} else {
		log.Fatalln(err)
	}
}

func main() {
	req, err := http.NewRequest(
		"POST",
		getEnv("PLUGIN_URL", defaultNtfyUrl),
		strings.NewReader(getEnv("PLUGIN_MESSAGE", "")),
	)

	checkErr(err)

	for _, h := range headers {
		h.add(&req.Header)
	}

	res, err := http.DefaultClient.Do(req)
	checkErr(err)

	defer res.Body.Close()

	b, err := io.ReadAll(res.Body)
	checkErr(err)

	if res.StatusCode != http.StatusOK {
		checkErr(fmt.Errorf("%s %s", res.Status, b))
	}

	log.Printf("%s %s", res.Status, b)
}
