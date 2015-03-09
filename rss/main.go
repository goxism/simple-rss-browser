package main

import (
	"fmt"
	"github.com/SlyMarbo/rss"
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	const episodesCommandName = "episodes"
	cli.AppHelpTemplate = apphelptmpl
	cli.CommandHelpTemplate = cmdhelptmpl
	app := &cli.App{
		Author: "Craig Weber",
		Commands: []cli.Command{
			cli.Command{
				Name:        "episodes",
				ShortName:   "ep",
				Description: "Print episode details",
				Action: func(c *cli.Context) {
					if len(c.Args()) < 1 {
						fmt.Println(c.Command.Usage)
						os.Exit(1)
					}

					feed, err := rss.Fetch(c.Args()[0])
					if err != nil {
						fmt.Fprintln(os.Stderr, err)
						os.Exit(1)
					}

					for i, episode := range feed.Items {
						fmt.Println("Ref:       ", len(feed.Items)-i-1)
						fmt.Println("Episode id:", episode.ID)
						fmt.Println("Title:     ", episode.Title)
						fmt.Println("Date:      ", episode.Date)
						if episode.Summary != "" {
							fmt.Println("Summary:   ", episode.Summary)
						}
						fmt.Println("Link:      ", episode.Link)
						fmt.Println()
					}
				},
				Usage: os.Args[0] + " " + episodesCommandName + " <feed-url>",
				// BashComplete
				// Before
				// Flags
				// SkipFlagParsing
				// Subcommands
			},
		},
		Email: "weberc2@gmail.com",
		Name:  "rss - a simple command-line rss browser",
		Action: func(c *cli.Context) {
			if len(c.Args()) < 1 {
				fmt.Println(c.App.Usage)
				os.Exit(1)
			}

			feed, err := rss.Fetch(c.Args()[0])
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(1)
			}

			fmt.Println("Title:      ", feed.Title)
			fmt.Println("NickName:   ", feed.Nickname)
			fmt.Println("Link:       ", feed.Link)
			fmt.Println("Episodes:   ", len(feed.Items))
			fmt.Println("Description:", feed.Description)
		},
		Usage: "rss <feed-url>",
		// BashComplete
		// Before
		// CommandNotFound
		// Compiled
		// EnableBashCompletion
		// Flags
		// Version
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
