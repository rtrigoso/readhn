package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rtrigoso/readhn/handler"
)

var (
	u   = flag.Bool("u", false, "")
	l   = flag.Bool("l", false, "")
	s   = flag.String("s", "p", "")
	n   = flag.Int("n", 10, "")
	err error
)

var usage = `Usage: readhn [options...]

Options:
	-n	Count of articles to show. Defaults to 10.
	-s	Sort by p: popular (Default), dd: date descending, da: date ascending 

	-u	Print the top article from the top stories.
	-l	Print the latest article submitted to Hacker News.
`

func main() {
	defer func() {
		if err != nil {
			fmt.Printf("error running tmt: %v", err)
			os.Exit(1)
		}
	}()

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	flag.Parse()
	hn := handler.Hn{}
	hn.New(handler.BEST_STORIES, 10)

	a, err := hn.ListArticles()
	if err != nil {
		return
	}

	fmt.Printf("%v", a)
}
