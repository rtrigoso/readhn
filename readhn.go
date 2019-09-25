package main

import (
	"flag"
	"fmt"
	"os"
)

const (
	hnbaseurl = "https://hacker-news.firebaseio.com/v0/"
)

var (
	u = flag.Bool("u", false, "")
	l = flag.Bool("l", false, "")
	s = flag.String("s", "p", "")
	n = flag.Int("n", 10, "")
)

var usage = `Usage: readhn [options...]

Options:
	-n	Count of articles to show. Defaults to 10.
	-s	Sort by p: popular (Default), dd: date descending, da: date ascending 

	-u	Print the top article from the top stories.
	-l	Print the latest article submitted to Hacker News.
`

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	flag.Parse()

}
