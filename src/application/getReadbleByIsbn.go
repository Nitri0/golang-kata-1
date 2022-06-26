package application

import "github.com/echocat/golang-kata-1/v1/src/domain"

type GetReadbleByIsbn struct {
	bookRepository domain.IBookRepository
}

func (i GetReadbleByIsbn) Invoke(isbn string) (domain.Readble, error) {

	results, err := i.bookRepository.FindByIsbn(isbn)
	if err != nil {
		return nil, err
	}
	return results, nil
}
