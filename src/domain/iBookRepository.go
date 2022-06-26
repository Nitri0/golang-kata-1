package domain

type IBookRepository interface {
	FindByIsbn(isbn string) (Book, error)
	FindByAuthor(email ...string) ([]Book, error)
	FindAllSortByTitle() ([]Book, error)
}
