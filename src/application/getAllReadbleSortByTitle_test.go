package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
	"github.com/echocat/golang-kata-1/v1/src/infraestructure"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func TestGetAllReadble_Invoke(t *testing.T) {

	testTable := []struct {
		testName    string
		expectError error
	}{
		{
			testName:    "Get all readble shoud be orderer",
			expectError: nil,
		},
	}

	bookRepository, err := infraestructure.NewMemoryBookRepository()
	assert.NoError(t, err)

	magazineRepository, err := infraestructure.NewMemoryMagazineRepository()
	assert.NoError(t, err)

	useCase := GetAllReadbleSortByTitle{
		bookRepository,
		magazineRepository,
	}

	for _, test := range testTable {
		t.Run(test.testName, func(t *testing.T) {
			result, err := useCase.Invoke()
			assert.NoError(t, err)
			assert.NotZero(t, len(result))
			assert.True(t, sort.StringsAreSorted(getArrayTitleStrings(result)), "Should be orderer")
		})
	}
}

func getArrayTitleStrings(readbles []domain.IReadble) []string {
	acumulator := []string{}
	for _, readble := range readbles {
		acumulator = append(acumulator, readble.GetTitle())
	}
	return acumulator
}
