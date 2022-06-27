package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
)

type GetReadbleByAuthors struct {
	BookRepository     domain.IBookRepository
	MagazineRepository domain.IMagazineRepository
}

func (a GetReadbleByAuthors) Invoke(authors ...string) ([]domain.IReadble, error) {
	var books []domain.Book
	var mgs []domain.Magazine

	books, err1 := a.BookRepository.FindByAuthor(authors...)
	mgs, err2 := a.MagazineRepository.FindByAuthor(authors...)

	if err1 != nil && err2 != nil {
		return nil, err1
	}

	var result []domain.IReadble
	for _, book := range books {
		result = append(result, book)
	}
	for _, mg := range mgs {
		result = append(result, mg)
	}
	return result, nil
}
