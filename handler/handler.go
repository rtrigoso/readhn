package handler

import (
	"errors"
	"fmt"
	"os"

	"github.com/PaulRosset/go-hacknews"
)

const (
	TOP_STORIES  = "topstories"
	NEW_STORIES  = "newstories"
	BEST_STORIES = "beststories"
)

var (
	err error
)

type Hn struct {
	container hacknews.Initializer
	posts     []hacknews.Post
	codes     []int
}

func (h *Hn) getCodes() {
	defer handleError()
	codes, err := h.container.GetCodesStory()
	if err != nil {
		return
	}

	h.codes = codes
}

func (h *Hn) getPosts() {
	defer handleError()
	posts, err := h.container.GetPostStory(h.codes)
	if err != nil {
		return
	}

	h.posts = posts
}

func (h *Hn) New(st string, n int) error {
	if st != TOP_STORIES && st != NEW_STORIES && st != BEST_STORIES {
		return errors.New(fmt.Sprintf("%s is not a valid story type", st))
	}

	h.container = hacknews.Initializer{Story: st, NbPosts: n}

	if len(h.codes) <= 0 {
		h.getCodes()
	}

	if len(h.posts) <= 0 {
		h.getPosts()
	}

	return nil
}

func (h *Hn) ListArticles() (map[string]string, error) {
	posts := make(map[string]string)
	for _, p := range h.posts {
		posts[p.Url] = p.Title
	}

	return posts, nil
}

// func (h Hn) GetArticlesBody() string {
// 	post := h.posts[0]
// }

func handleError() {
	if err != nil {
		fmt.Printf("error running tmt: %v", err)
		os.Exit(1)
	}
}
