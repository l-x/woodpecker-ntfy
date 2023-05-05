package main

import (
	"fmt"
	"os"
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestSettings(t *testing.T) {
	testCases := []struct{
		setting setting
		envName,
		envValue,
		expectedValue string
	}{
		{server, "PLUGIN_URL", "", defaultNtfyUrl},
		{server, "PLUGIN_URL", "https://plugin.url/topic", "https://plugin.url/topic"},
		
		{message, "PLUGIN_MESSAGE", "", ""},
		{message, "PLUGIN_MESSAGE", "plugin message", "plugin message"},
		
		{headers["Authorization"], "PLUGIN_TOKEN", "", ""},
		{headers["Authorization"], "PLUGIN_TOKEN", "token value", "Bearer token value"},
		
		{headers["Title"], "PLUGIN_TITLE", "", ""},
		{headers["Title"], "PLUGIN_TITLE", "plugin title", "plugin title"},

		{headers["Priority"], "PLUGIN_PRIORITY", "", ""},
		{headers["Priority"], "PLUGIN_PRIORITY", "plugin priority", "plugin priority"},

		{headers["Tags"], "PLUGIN_TAGS", "", ""},
		{headers["Tags"], "PLUGIN_TAGS", "plugin,tags", "plugin,tags"},

		{headers["Actions"], "PLUGIN_ACTIONS", "", ""},
		{headers["Actions"], "PLUGIN_ACTIONS", "plugin actions", "plugin actions"},

		{headers["Click"], "PLUGIN_CLICK", "", os.Getenv("CI_BUILD_LINK")},
		{headers["Click"], "PLUGIN_CLICK", "https://plugin.click", "https://plugin.click"},
		
		{headers["Icon"], "PLUGIN_ICON", "", os.Getenv("CI_COMMIT_AUTHOR_AVATAR")},
		{headers["Icon"], "PLUGIN_ICON", "https://plugin.icon", "https://plugin.icon"},
		
		{failOnError, "PLUGIN_FAIL_ON_ERROR", "", ""},
		{failOnError, "PLUGIN_FAIL_ON_ERROR", "fail!", "fail!"},
	}
	
	for _, tc := range testCases {
		t.Run(
			fmt.Sprintf("%s=\"%s\"", tc.envName, tc.envValue),
			func(t *testing.T) {
				os.Setenv(tc.envName, tc.envValue)
				assert.Equal(t, tc.expectedValue, tc.setting.getValue())
			},
		)
	}
}
