package infraestructure

import (
	"encoding/csv"
	"github.com/echocat/golang-kata-1/v1/src/domain"
	"os"
	"strings"
)

const CSV_URL_PATH_MAGAZINE = "./../../resources/magazines.csv"

type magazineRaw struct {
	title       string
	isbn        string
	authors     string
	publishedAt string
}

type MemoryMagazineRepository struct {
	data []magazineRaw
}

func (r MemoryMagazineRepository) FindByIsbn(isbn string) (domain.Magazine, error) {
	for _, magazineRaw := range r.data {
		if magazineRaw.isbn == isbn {
			bookFound, err := domain.NewMagazine(
				magazineRaw.title,
				magazineRaw.isbn,
				magazineRaw.authors,
				magazineRaw.publishedAt,
			)
			if err != nil {
				return domain.Magazine{}, err
			}
			return bookFound, nil
		}
	}
	return domain.Magazine{}, domain.NotFoundReadbleErr
}

func (r MemoryMagazineRepository) FindByAuthor(emails ...string) ([]domain.Magazine, error) {
	acumulator := []domain.Magazine{}
	for _, rawMg := range r.data {
	L:
		for _, email := range emails {
			if strings.Contains(rawMg.authors, email) {
				mg, err := domain.NewMagazine(
					rawMg.title,
					rawMg.isbn,
					rawMg.authors,
					rawMg.publishedAt,
				)
				// TODO: take the error case
				if err != nil {
					break L
				}
				acumulator = append(acumulator, mg)
				break L
			}
		}
	}
	if len(acumulator) == 0 {
		return nil, domain.NotFoundReadbleErr
	}

	return acumulator, nil
}
func (r MemoryMagazineRepository) FindAllSortByTitle() ([]domain.Magazine, error) {
	return nil, nil
}

func (r MemoryMagazineRepository) GetAll() ([]domain.Magazine, error) {
	acumulator := []domain.Magazine{}
	for _, rawMg := range r.data {
		mg, err := domain.NewMagazine(
			rawMg.title,
			rawMg.isbn,
			rawMg.authors,
			rawMg.publishedAt,
		)
		if err != nil {
			return nil, err
		}
		acumulator = append(acumulator, mg)
	}
	return acumulator, nil
}

func NewMemoryMagazineRepository() (MemoryMagazineRepository, error) {
	f, err := os.Open(CSV_URL_PATH_MAGAZINE)
	if err != nil {
		return MemoryMagazineRepository{}, domain.InvalidDataOriginErr
	}
	csvReader := csv.NewReader(f)
	csvReader.Comma = ';'
	allData, err := csvReader.ReadAll()
	if err != nil {
		return MemoryMagazineRepository{}, domain.InvalidDataOriginErr
	}
	magazines := []magazineRaw{}
	for index, data := range allData {
		if index != 0 {
			magazine := magazineRaw{
				data[0],
				data[1],
				data[2],
				data[3],
			}
			magazines = append(magazines, magazine)
		}
	}

	return MemoryMagazineRepository{magazines}, nil
}
