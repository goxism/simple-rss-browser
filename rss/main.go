package main

import (
	"fmt"
	"github.com/SlyMarbo/rss"
	"github.com/codegangsta/cli"
	"os"
)

func infoCommand(c *cli.Context) {
	if len(c.Args()) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
		os.Exit(1)
	}

	feed, err := rss.Fetch(c.Args()[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if feed.Title != "" {
		fmt.Println("Title:      ", feed.Title)
	}
	if feed.Nickname != "" {
		fmt.Println("NickName:   ", feed.Nickname)
	}
	if feed.Link != "" {
		fmt.Println("Link:       ", feed.Link)
	}
	if feed.Items != nil {
		fmt.Println("Episodes:   ", len(feed.Items))
	}
	if feed.Description != "" {
		fmt.Println("Description:", feed.Description)
	}
}

func episodesCommand(c *cli.Context) {
	if len(c.Args()) < 1 {
		cli.ShowCommandHelp(c, c.Command.Name)
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
}

func main() {
	const episodesCommandName = "episodes"
	const infoCommandName = "info"
	cli.AppHelpTemplate = apphelptmpl
	cli.CommandHelpTemplate = cmdhelptmpl
	app := &cli.App{
		Action: cli.ShowAppHelp,
		Author: "Craig Weber",
		Commands: []cli.Command{
			cli.Command{
				Name:        infoCommandName,
				ShortName:   "in",
				Description: "Fetch basic feed info",
				Action:      infoCommand,
				Usage:       os.Args[0] + " " + infoCommandName + " <feed-url>",
			},
			cli.Command{
				Name:        episodesCommandName,
				ShortName:   "ep",
				Description: "Print episode details",
				Action:      episodesCommand,
				Usage:       os.Args[0] + " " + episodesCommandName + " <feed-url>",
				// BashComplete
				// Before
				// Flags
				// SkipFlagParsing
				// Subcommands
			},
		},
		Email: "weberc2@gmail.com",
		Name:  "rss - a simple command-line rss browser",
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
