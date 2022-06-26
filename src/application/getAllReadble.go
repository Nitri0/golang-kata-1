package application

import "github.com/echocat/golang-kata-1/v1/src/domain"

type GetAllReadble struct {
}

func (a GetAllReadble) Invoke() ([]domain.Readble, error) {
	return nil, nil
}
