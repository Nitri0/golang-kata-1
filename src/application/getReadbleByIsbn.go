package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
)

type GetReadbleByIsbn struct {
	BookRepository     domain.IBookRepository
	MagazineRepository domain.IMagazineRepository
}

func (i GetReadbleByIsbn) Invoke(isbn string) (domain.IReadble, error) {

	resultsBooks, err1 := i.BookRepository.FindByIsbn(isbn)
	if err1 == nil {
		return resultsBooks, nil
	}

	resultsMagazine, err2 := i.MagazineRepository.FindByIsbn(isbn)

	if err2 == nil {
		return resultsMagazine, nil
	}

	return nil, err1

}
