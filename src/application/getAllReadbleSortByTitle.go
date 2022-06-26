package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
)

type GetAllReadbleSortByTitle struct {
	bookRepository     domain.IBookRepository
	magazineRepository domain.IMagazineRepository
}

func (a GetAllReadbleSortByTitle) Invoke() ([]domain.IReadble, error) {
	books, err1 := a.bookRepository.GetAll()
	mgs, err2 := a.magazineRepository.GetAll()
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
