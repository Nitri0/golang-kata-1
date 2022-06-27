package domain

import (
	"fmt"
	"strings"
	"time"
)

type Magazine struct {
	title       string
	isbn        string
	authors     []string
	publishedAt time.Time
}

func (b Magazine) GetTitle() string {
	return b.title
}

func (b Magazine) GetPrintData() (result []string) {
	title := fmt.Sprintf("Tituto: %v", b.title)
	typeReadble := fmt.Sprintf("Tipo de texto: Revista")
	isbn := fmt.Sprintf("ISBN: %v", b.isbn)
	publishedAt := fmt.Sprintf("Fecha de publicación: %v", b.publishedAt.Format("2006-01-02"))
	authors := fmt.Sprintf("Authors emails: %v", strings.Join(b.authors, ", "))

	result = append(result, title, typeReadble, isbn, publishedAt, authors)
	return
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
