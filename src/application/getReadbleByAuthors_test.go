package application

import (
	"github.com/echocat/golang-kata-1/v1/src/domain"
	"github.com/echocat/golang-kata-1/v1/src/infraestructure"
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
		}, {
			"pass a incomplete email, should be return 6 result",
			[]string{"walter"},
			6,
			nil,
		},
		{
			"pass a invalid email, should be return 0 result",
			[]string{"test"},
			0,
			domain.NotFoundReadbleErr,
		}, {
			"pass a incomplete email with a valid email, should be return 6 result",
			[]string{"test", "walter"},
			6,
			nil,
		},
	}
	bookRepository, err := infraestructure.NewMemoryBookRepository()
	assert.NoError(t, err)

	magazineRepository, err := infraestructure.NewMemoryMagazineRepository()
	assert.NoError(t, err)

	useCase := GetReadbleByAuthors{
		bookRepository,
		magazineRepository,
	}

	for _, test := range testTable {
		t.Run(test.testName, func(t *testing.T) {
			result, err := useCase.Invoke(test.paramAuthor...)
			assert.Equal(t, test.expectError, err)
			assert.Equal(t, test.expectLenResult, len(result))
		})

	}
}
