package infraestructure

import (
	"encoding/csv"
	"github.com/echocat/golang-kata-1/v1/src/domain"
	"os"
	"strings"
)

const CSV_URL_PATH = "./../../resources/books.csv"

type booksRaw struct {
	title       string
	isbn        string
	authors     string
	description string
}

type MemoryBookRepository struct {
	data []booksRaw
}

func (r MemoryBookRepository) FindByIsbn(isbn string) (domain.Book, error) {
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

func (r MemoryBookRepository) FindByAuthor(emails ...string) ([]domain.Book, error) {
	acumulator := []domain.Book{}
	for _, rawBook := range r.data {
	L:
		for _, email := range emails {
			if strings.Contains(rawBook.authors, email) {
				book, err := domain.NewBook(
					rawBook.title,
					rawBook.isbn,
					rawBook.description,
					rawBook.authors,
				)
				if err != nil {
					break L
				}
				acumulator = append(acumulator, book)
				break L
			}
		}
	}
	if len(acumulator) == 0 {
		return nil, domain.NotFoundReadbleErr
	}
	return acumulator, nil
}

func (r MemoryBookRepository) FindAllSortByTitle() ([]domain.Book, error) {
	return nil, nil
}

func (r MemoryBookRepository) GetAll() ([]domain.Book, error) {
	acumulator := []domain.Book{}
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

func NewMemoryBookRepository() (MemoryBookRepository, error) {
	f, err := os.Open(CSV_URL_PATH)
	if err != nil {
		return MemoryBookRepository{}, domain.InvalidDataOriginErr
	}
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	allData, err := csvReader.ReadAll()
	if err != nil {
		return MemoryBookRepository{}, domain.InvalidDataOriginErr
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

	return MemoryBookRepository{books}, nil
}
