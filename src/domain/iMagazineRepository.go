package domain

type IMagazineRepository interface {
	FindByIsbn(isbn string) (Magazine, error)
	FindByAuthor(email ...string) ([]Magazine, error)
	FindAllSortByTitle() ([]Magazine, error)
}
