package plugin

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {
	var testConfig = &Config{
		Message:  "the message",
		Token:    "the token",
		Title:    "the title",
		Click:    "https://the.click.url",
		Icon:     "https://the.icon.url",
		Priority: "alert",
		Actions:  "the actions",
		Tags:     "the,notification,tags",
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/topic")
		assert.Equal(t, req.Header.Get("Authorization"), "Bearer "+testConfig.Token)
		assert.Equal(t, req.Header.Get("Title"), testConfig.Title)
		assert.Equal(t, req.Header.Get("Click"), testConfig.Click)
		assert.Equal(t, req.Header.Get("Icon"), testConfig.Icon)
		assert.Equal(t, req.Header.Get("Priority"), testConfig.Priority)
		assert.Equal(t, req.Header.Get("Actions"), testConfig.Actions)
		assert.Equal(t, req.Header.Get("Tags"), testConfig.Tags)

		rw.Write([]byte(`OK`))
	}))

	testConfig.URL = server.URL + "/topic"

	err := New(testConfig).Run()
	assert.Nil(t, err)
}

func TestRunWithDefaults(t *testing.T) {
	var testConfig = &Config{}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), "/topic")
		assert.Equal(t, req.Header.Get("Authorization"), testConfig.Token)
		assert.Equal(t, req.Header.Get("Title"), testConfig.Title)
		assert.Equal(t, req.Header.Get("Click"), testConfig.Click)
		assert.Equal(t, req.Header.Get("Icon"), testConfig.Icon)
		assert.Equal(t, req.Header.Get("Priority"), testConfig.Priority)
		assert.Equal(t, req.Header.Get("Actions"), testConfig.Actions)
		assert.Equal(t, req.Header.Get("Tags"), testConfig.Tags)

		rw.Write([]byte(`OK`))
	}))

	testConfig.URL = server.URL + "/topic"

	assert.Nil(t, New(testConfig).Run())
}

func TestRunWithServerError(t *testing.T) {
	var testConfig = &Config{}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(403)
		rw.Write([]byte(`error`))
	}))

	testConfig.URL = server.URL + "/topic"

	assert.Error(t, New(testConfig).Run(), "error")
}
