package domain

import (
	"fmt"
	"strings"
)

type Book struct {
	title       string
	isbn        string
	description string
	authorEmail []string
}

func (b Book) GetTitle() string {
	return b.title
}

func (b Book) GetPrintData() (result []string) {
	title := fmt.Sprintf("Tituto: %v", b.title)
	typeReadble := fmt.Sprintf("Tipo de texto: Libro")
	isbn := fmt.Sprintf("ISBN: %v", b.isbn)
	description := fmt.Sprintf("Description: %v", b.description)
	authors := fmt.Sprintf("Authors emails: %v", strings.Join(b.authorEmail, ", "))

	result = append(result, title, typeReadble, isbn, description, authors)
	return
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
