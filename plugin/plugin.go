package plugin

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	client httpClient  = http.DefaultClient
	debug  *log.Logger = log.New(io.Discard, "DEBUG ", log.Ldate+log.Ltime)
)

type httpClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type Config struct {
	URL,
	Message,
	Token,
	Title,
	Click,
	Icon,
	Priority,
	Actions,
	Email,
	Tags string
	Debug bool
}

type plugin struct {
	config *Config
	client httpClient
}

func (p *plugin) Run() error {
	if p.config.Debug {
		debug.SetOutput(os.Stderr)
		for _, v := range os.Environ() {
			debug.Print(v)
		}
	}

	debug.Printf("Request: %+v", p.config)

	req, err := createRequest(p.config)
	if err != nil {
		return err
	}

	res, err := p.client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if err != nil {
		return err
	}

	debug.Printf("Response: %s %s", res.Status, body)

	if res.StatusCode != http.StatusOK {
		return errors.New(string(body))
	}

	log.Printf("Notification successfully sent to %s", p.config.URL)

	return nil
}

func New(config *Config) *plugin {
	return &plugin{
		config: config,
		client: client,
	}
}

func setHeader(header *http.Header, kv map[string]string) {
	for k, v := range kv {
		if v == "" {
			continue
		}
		header.Set(k, v)
	}
}

func createRequest(c *Config) (*http.Request, error) {
	req, err := http.NewRequest(
		"PUT",
		c.URL,
		strings.NewReader(c.Message),
	)
	if err != nil {
		return req, err
	}

	if c.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.Token))
	}

	setHeader(&req.Header, map[string]string{
		"Title":    c.Title,
		"Click":    c.Click,
		"Icon":     c.Icon,
		"Priority": c.Priority,
		"Actions":  c.Actions,
		"Tags":     c.Tags,
		"Email":    c.Email,
	})

	return req, nil
}
