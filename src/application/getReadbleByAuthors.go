package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
)

type GetReadbleByAuthors struct {
	bookRepository     domain.IBookRepository
	magazineRepository domain.IMagazineRepository
}

func (a GetReadbleByAuthors) Invoke(authors ...string) ([]domain.Readble, error) {
	var books []domain.Book
	var mgs []domain.Magazine

	books, err1 := a.bookRepository.FindByAuthor(authors...)
	mgs, err2 := a.magazineRepository.FindByAuthor(authors...)

	if err1 != nil && err2 != nil {
		return nil, err1
	}

	var result []domain.Readble
	for _, book := range books {
		result = append(result, book)
	}
	for _, mg := range mgs {
		result = append(result, mg)
	}
	return result, nil
}
