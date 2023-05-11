package main

import (
	"log"
	"os"
	"woodpecker-ntfy/plugin"

	"github.com/urfave/cli/v2"
)

const defaultNtfyUrl = "https://ntfy.sh/woodpecker-ntfy"

func run(c *cli.Context) error {
	config := &plugin.Config{
		URL:      c.String("url"),
		Token:    c.String("token"),
		Title:    c.String("title"),
		Priority: c.String("priority"),
		Actions:  c.String("actions"),
		Click:    c.String("click"),
		Tags:     c.String("tags"),
		Message:  c.String("message"),
	}

	return plugin.New(config).Run()
}

func createApp() *cli.App {
	app := cli.NewApp()

	app.Name = "woodpecker-ntfy"
	app.Usage = "Woodpecker plugin to send notifications to a ntfy.sh instance"
	app.Action = run
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "url",
			Usage:   "ntfy instance url (including topic)",
			EnvVars: []string{"PLUGIN_URL"},
			Value:   defaultNtfyUrl,
		},
		&cli.StringFlag{
			Name:    "token",
			Usage:   "access token for sending notifications to write-protected topics",
			EnvVars: []string{"PLUGIN_TOKEN"},
		},
		&cli.StringFlag{
			Name:    "title",
			Usage:   "notification title",
			EnvVars: []string{"PLUGIN_TITLE"},
		},
		&cli.StringFlag{
			Name:    "priority",
			Usage:   "notification priority",
			EnvVars: []string{"PLUGIN_PRIORITY"},
		},
		&cli.StringFlag{
			Name:    "actions",
			Usage:   "notification actions",
			EnvVars: []string{"PLUGIN_ACTIONS"},
		},
		&cli.StringFlag{
			Name:    "click",
			Usage:   "notification click url",
			EnvVars: []string{"PLUGIN_PRIORITY", "CI_BUILD_LINK"},
		},
		&cli.StringFlag{
			Name:    "icon",
			Usage:   "notification icon url",
			EnvVars: []string{"PLUGIN_ICON", "CI_COMMIT_AUTHOR_AVATAR"},
		},
		&cli.StringFlag{
			Name:    "tags",
			Usage:   "notification tags",
			EnvVars: []string{"PLUGIN_TAGS"},
		},
		&cli.StringFlag{
			Name:    "message",
			Usage:   "notification message body",
			EnvVars: []string{"PLUGIN_MESSAGE"},
		},
	}

	return app
}

func main() {
	if err := createApp().Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
