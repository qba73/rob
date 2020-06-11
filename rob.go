package main

import (
	"fmt"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing/object"
	"github.com/go-git/go-git/v5/storage/memory"
	"time"
)

func main() {
	rurl := "https://github.com/qba73/fbom"

	r, err := git.Clone(memory.NewStorage(), nil, &git.CloneOptions{
		URL: rurl,
	})
	if err != nil {
		fmt.Println(err)
	}

	// get head
	ref, err := r.Head()
	if err != nil {
		fmt.Println(err)
	}

	since := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
	until := time.Now()

	cIter, err := r.Log(&git.LogOptions{From: ref.Hash(), Since: &since, Until: &until})
	if err != nil {
		fmt.Println(err)
	}

	// iterate over the commits
	err = cIter.ForEach(func(c *object.Commit) error {
		fmt.Println(c)
		return nil
	})
	if err != nil {
		fmt.Println()
	}
}
