README
======

This repository defines a simple command (`rss`) for browsing podcasts. It's powered entirely by https://github.com/SlyMarbo/rss, and it is written in Go.

##USAGE
`rss <feed-url>`

```
$ rss http://blog.stackoverflow.com/feed/
Title:       Blog - Stack Exchange
Episodes:    40
Link:        http://blog.stackoverflow.com
Description: free, community powered Q&A
```

##INSTALLATION

1. Install [Go](https://golang.org)
2. `go get https://github.com/weberc2/simple-rss-browser/rss`
