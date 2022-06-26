package application

import (
	"github.com/echocat/golang-kata-1/v1/src/infraestructure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAllReadble_Invoke(t *testing.T) {

	testTable := []struct {
		testName    string
		expectError error
	}{
		{
			testName:    "happy path, shouldn't be error",
			expectError: nil,
		},
	}

	bookRepository, err := infraestructure.NewMemoryBookRepository()
	assert.NoError(t, err)

	magazineRepository, err := infraestructure.NewMemoryMagazineRepository()
	assert.NoError(t, err)

	useCase := GetAllReadble{
		bookRepository,
		magazineRepository,
	}

	for _, test := range testTable {
		t.Run(test.testName, func(t *testing.T) {
			result, err := useCase.Invoke()
			assert.NoError(t, err)
			assert.NotZero(t, len(result))
		})
	}
}
