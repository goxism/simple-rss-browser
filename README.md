README
======

This repository defines a simple command (`rss`) for browsing podcasts. It's powered entirely by https://github.com/SlyMarbo/rss, and it is written in Go.

##USAGE
* Get feed info: `rss info <feed-url>`
* List feed's episodes: `rss episodes <feed-url>`
* Help: `rss help <feed-url>`

```
$ rss info http://blog.stackoverflow.com/feed/
Title:       Blog - Stack Exchange
Episodes:    40
Link:        http://blog.stackoverflow.com
Description: free, community powered Q&A
```

##INSTALLATION

1. Install [Go](https://golang.org)
2. `go get https://github.com/weberc2/simple-rss-browser/rss`
