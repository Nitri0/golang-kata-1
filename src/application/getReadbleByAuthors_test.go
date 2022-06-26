package application

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetReadbleByAuthors_Invoke(t *testing.T) {
	testTable := []struct {
		testName        string
		paramAuthor     []string
		expectLenResult int
		expectError     error
	}{
		{
			"pass a valid email, should be return two result (a book and a magazine)",
			[]string{"null-gustafsson@echocat.org"},
			2,
			nil,
		},
	}

	useCase := GetReadbleByAuthors{}

	for _, test := range testTable {
		t.Run(test.testName, func(t *testing.T) {
			result, err := useCase.Invoke(test.paramAuthor...)
			assert.Equal(t, test.expectError, err)
			assert.Equal(t, test.expectLenResult, len(result))
		})

	}
}
