// +AceptationTest
package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
	"github.com/echocat/golang-kata-1/v1/src/infraestructure"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetReadbleByIsbn(t *testing.T) {
	tableCases := []struct {
		nameTest    string
		isbn        string
		expectError error
	}{
		{
			"Valid ISBN should be without error",
			"5554-5545-4518",
			nil,
		}, {
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

	repository := infraestructure.Repository{}
	useCase := GetReadbleByIsbn{repository}

	for _, tc := range tableCases {
		t.Run(tc.nameTest, func(t *testing.T) {
			_, error := useCase.Invoke(tc.isbn)
			assert.Equal(t, tc.expectError, error)
		})
	}

}
