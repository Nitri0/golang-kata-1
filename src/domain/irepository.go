package domain

type IRepository interface {
	FindReadbleByIsbn(isbn string) (Readble, error)
	FindReadbleByAuthor(email ...string) ([]Readble, error)
	FindAllReadbleSortByTitle() ([]Readble, error)
}
