package main

import (
	"fmt"
	"log"
	"os"

	"github.com/slack-go/slack"
	"github.com/urfave/cli/v2"
)

var (
	myVar       string
	version     string
	slackClient *slack.Client
)

func init() {
	myVar = "hi"
	version = "1.0.0"
	slackClient = slack.New(os.Getenv("SLACK_TOKEN"))
}

func main() {
	app := &cli.App{
		Name:    "vormir",
		Usage:   "Runs some stuff",
		Version: version,
		Commands: []*cli.Command{
			{
				Name:    "doctor",
				Aliases: []string{"dc"},
				Usage:   "Simple command to test installation + configuration",
				Action: func(c *cli.Context) error {
					fmt.Printf("Version: %q\nmyVar: %q\n", version, myVar)
					return nil
				},
			},
			{
				Name:  "users",
				Usage: "Updates the list of available users",
				Action: func(c *cli.Context) error {
					users := findUsersThatAreQuitting()

					for _, user := range users {
						fmt.Printf("%s is quitting!", user.DisplayName)
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
