package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
)

type GetAllReadble struct {
	bookRepository     domain.IBookRepository
	magazineRepository domain.IMagazineRepository
}

func (a GetAllReadble) Invoke() ([]domain.Readble, error) {
	books, err := a.bookRepository.GetAll()
	if err != nil {
		return nil, err
	}

	var result []domain.Readble

	for _, book := range books {
		result = append(result, book)
	}

	return result, nil
}
