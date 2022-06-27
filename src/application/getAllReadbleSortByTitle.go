package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
	"sort"
)

type GetAllReadbleSortByTitle struct {
	BookRepository     domain.IBookRepository
	MagazineRepository domain.IMagazineRepository
}

func (a GetAllReadbleSortByTitle) Invoke() ([]domain.IReadble, error) {
	books, err1 := a.BookRepository.GetAll()
	mgs, err2 := a.MagazineRepository.GetAll()
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
	orderByTitle(result)
	return result, nil
}

func orderByTitle(readbles []domain.IReadble) {
	sort.SliceStable(readbles, func(i, j int) bool {
		return readbles[i].GetTitle() < readbles[j].GetTitle()
	})
}
