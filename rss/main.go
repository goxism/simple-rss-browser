package main

import (
	"fmt"
	"github.com/SlyMarbo/rss"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("USAGE: %s <feed-url>\n", os.Args[0])
		os.Exit(1)
	}

	feed, err := rss.Fetch(os.Args[1])
	if err != nil {
		panic(err)
	}

	fmt.Println("Title:      ", feed.Title)
	fmt.Println("Episodes:   ", len(feed.Items))
	fmt.Println("Link:       ", feed.Link)
	fmt.Println("Description:", feed.Description)
}
