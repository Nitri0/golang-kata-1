package domain

type Repository interface {
	FindReadbleByIsbn(isbn string) (Readble, error)
	FindReadbleByAuthor(email ...string) ([]Readble, error)
	FindAllReadbleSortByTitle() ([]Readble, error)
}
