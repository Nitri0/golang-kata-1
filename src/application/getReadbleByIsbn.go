package application

import "github.com/echocat/golang-kata-1/v1/src/domain"

type GetReadbleByIsbn struct {
	repository domain.IRepository
}

func (i GetReadbleByIsbn) Invoke(isbn string) (domain.Readble, error) {

	results, err := i.repository.FindReadbleByIsbn(isbn)
	if err != nil {
		return nil, err
	}
	return results, nil
}
