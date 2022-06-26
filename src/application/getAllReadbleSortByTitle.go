package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
)

type GetAllReadble struct {
	bookRepository     domain.IBookRepository
	magazineRepository domain.IMagazineRepository
}

func (a GetAllReadble) Invoke() ([]domain.Readble, error) {
	books, err1 := a.bookRepository.GetAll()
	mgs, err2 := a.magazineRepository.GetAll()
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
