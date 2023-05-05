package main

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"os"
	"strings"
	"testing"
)

var envVars = []string{
	"PLUGIN_URL",
	"PLUGIN_MESSAGE",
	"PLUGIN_TOKEN",
	"PLUGIN_TITLE",
	"PLUGIN_PRIORITY",
	"PLUGIN_TAGS",
	"PLUGIN_ACTIONS",
	"PLUGIN_CLICK",
	"PLUGIN_ICON",
}

func unsetEnvVars() {
	for _, v := range envVars {
		os.Unsetenv(v)
	}
}

func TestCreateRequestFromDefaults(t *testing.T) {
	unsetEnvVars()
	req, err := createRequest()

	assert.NoError(t, err)
	assert.NotNil(t, req)

	assert.Equal(t, defaultNtfyUrl, req.URL.String())
	assert.Equal(t, "PUT", req.Method)
	assert.Equal(t, http.NoBody, req.Body)
	assert.Equal(t, "", req.Header.Get("Authorization"))
	assert.Equal(t, "", req.Header.Get("Title"))
	assert.Equal(t, "", req.Header.Get("Priority"))
	assert.Equal(t, "", req.Header.Get("Tags"))
	assert.Equal(t, "", req.Header.Get("Actions"))
	assert.Equal(t, os.Getenv("CI_BUILD_LINK"), req.Header.Get("Click"))
	assert.Equal(t, os.Getenv("CI_COMMIT_AUTHOR_AVATAR"), req.Header.Get("Icon"))
}

func TestCreateRequestFromSetting(t *testing.T) {
	defer unsetEnvVars()

	for _, v := range envVars {
		os.Setenv(v, v)
	}

	req, err := createRequest()

	assert.NoError(t, err)
	assert.NotNil(t, req)

	body := new(strings.Builder)
	io.Copy(body, req.Body)

	assert.Equal(t, "PLUGIN_URL", req.URL.String())
	assert.Equal(t, "PUT", req.Method)
	assert.Equal(t, "PLUGIN_MESSAGE", body.String())
	assert.Equal(t, "Bearer PLUGIN_TOKEN", req.Header.Get("Authorization"))
	assert.Equal(t, "PLUGIN_TITLE", req.Header.Get("Title"))
	assert.Equal(t, "PLUGIN_PRIORITY", req.Header.Get("Priority"))
	assert.Equal(t, "PLUGIN_TAGS", req.Header.Get("Tags"))
	assert.Equal(t, "PLUGIN_ACTIONS", req.Header.Get("Actions"))
	assert.Equal(t, "PLUGIN_CLICK", req.Header.Get("Click"))
	assert.Equal(t, "PLUGIN_ICON", req.Header.Get("Icon"))
}
