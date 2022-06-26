package domain

import (
	"strings"
	"time"
)

type Magazine struct {
	title       string
	isbn        string
	authors     []string
	publishedAt time.Time
}

func NewMagazine(title string, isbn string, authors string, publishedAt string) (Magazine, error) {
	authorsSplited := strings.Split(authors, ",")
	publishedAtFormated := time.Now() // RREFACTOR

	return Magazine{
		title:       title,
		isbn:        isbn,
		authors:     authorsSplited,
		publishedAt: publishedAtFormated,
	}, nil
}
