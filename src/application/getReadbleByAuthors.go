package application

import "github.com/echocat/golang-kata-1/v1/src/domain"

type GetReadbleByAuthors struct {
}

func (a GetReadbleByAuthors) Invoke(authors ...string) ([]domain.Readble, error) {
	return nil, nil
}
