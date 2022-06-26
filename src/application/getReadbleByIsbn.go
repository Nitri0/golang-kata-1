package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
)

type GetReadbleByIsbn struct {
	bookRepository     domain.IBookRepository
	magazineRepository domain.IMagazineRepository
}

func (i GetReadbleByIsbn) Invoke(isbn string) (domain.Readble, error) {

	resultsBooks, err1 := i.bookRepository.FindByIsbn(isbn)
	if err1 == nil {
		return resultsBooks, nil
	}

	resultsMagazine, err2 := i.magazineRepository.FindByIsbn(isbn)

	if err2 == nil {
		return resultsMagazine, nil
	}

	return nil, err1

}
