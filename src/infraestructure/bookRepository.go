package infraestructure

import (
	"encoding/csv"
	"github.com/echocat/golang-kata-1/v1/src/domain"
	"os"
)

const CSV_URL_PATH = "./../../resources/books.csv"

type booksRaw struct {
	title       string
	isbn        string
	authors     string
	description string
}

type BookRepository struct {
	data []booksRaw
}

func (r BookRepository) FindByIsbn(isbn string) (domain.Book, error) {
L:
	for _, bookRaw := range r.data {
		if bookRaw.isbn == isbn {
			bookFound, err := domain.NewBook(
				bookRaw.title,
				bookRaw.isbn,
				bookRaw.description,
				bookRaw.authors,
			)
			if err != nil {
				break L
			}
			return bookFound, nil
		}
	}
	return domain.Book{}, domain.NotFoundReadbleErr
}

func (r BookRepository) FindByAuthor(email ...string) ([]domain.Book, error) {
	return nil, nil
}
func (r BookRepository) FindAllSortByTitle() ([]domain.Book, error) {
	return nil, nil
}

func (r BookRepository) GetAll() ([]domain.Book, error) {
	var acumulator []domain.Book
	for _, rawBook := range r.data {
		book, err := domain.NewBook(
			rawBook.title,
			rawBook.isbn,
			rawBook.description,
			rawBook.authors,
		)
		if err != nil {
			return nil, err
		}
		acumulator = append(acumulator, book)
	}
	return acumulator, nil
}

func NewBookRepository() (BookRepository, error) {
	f, err := os.Open(CSV_URL_PATH)
	if err != nil {
		return BookRepository{}, domain.InvalidDataOriginErr
	}
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	allData, err := csvReader.ReadAll()
	if err != nil {
		return BookRepository{}, domain.InvalidDataOriginErr
	}
	books := []booksRaw{}
	for _, data := range allData {
		book := booksRaw{
			data[0],
			data[1],
			data[2],
			data[3],
		}
		books = append(books, book)
	}

	return BookRepository{books}, nil
}
