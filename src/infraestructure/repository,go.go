package infraestructure

import "github.com/echocat/golang-kata-1/v1/src/domain"

type Repository struct {
}

func (r Repository) FindReadbleByIsbn(isbn string) (domain.Readble, error) {
	return nil, nil
}
func (r Repository) FindReadbleByAuthor(email ...string) ([]domain.Readble, error) {
	return nil, nil
}
func (r Repository) FindAllReadbleSortByTitle() ([]domain.Readble, error) {
	return nil, nil
}
