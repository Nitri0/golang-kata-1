// +AceptationTest
package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
	"github.com/echocat/golang-kata-1/v1/src/infraestructure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetReadbleByIsbn_Invoke(t *testing.T) {
	tableCases := []struct {
		nameTest    string
		isbn        string
		expectError error
	}{
		{
			"Valid ISBN Book should be without error",
			"5554-5545-4518",
			nil,
		},
		{
			"Valid ISBN Magazine should be without error",
			"5454-5587-3210",
			nil,
		},
		{
			"valid ISBN not exist should be NotFoundReadbleErr",
			"5554-5545-4512",
			domain.NotFoundReadbleErr,
		},
		/*
			{
				"invalid ISBN should be InvalidFormatIsbnErr",
				"00000",
				domain.InvalidFormatIsbnErr,
			},
		*/

	}

	bookRepository, err := infraestructure.NewBookRepository()
	assert.Nil(t, err)

	magRepository, err := infraestructure.NewMagazineRepository()
	assert.Nil(t, err)

	useCase := GetReadbleByIsbn{
		bookRepository,
		magRepository}

	for _, tc := range tableCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			_, error := useCase.Invoke(tc.isbn)
			assert.Equal(t, tc.expectError, error)
		})
	}

}
