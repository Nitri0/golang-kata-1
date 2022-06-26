package domain

import "strings"

type Book struct {
	title       string
	isbn        string
	description string
	authorEmail []string
}

func (b Book) GetTitle() string {
	return b.title
}

func NewBook(title string, isbn string, description string, authorEmails string) (Book, error) {
	authors := strings.Split(authorEmails, ",")
	return Book{
		title:       title,
		isbn:        isbn,
		description: description,
		authorEmail: authors,
	}, nil
}
