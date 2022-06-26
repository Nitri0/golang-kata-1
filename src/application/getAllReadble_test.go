package application

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllReadble_Invoke(t *testing.T) {

	useCase := GetAllReadble{}
	result, err := useCase.Invoke()

	assert.NoError(t, err)
	assert.NotZero(t, len(result))
}
