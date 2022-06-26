package domain

import "errors"

var (
	InvalidDataOriginErr = errors.New("Can't conect source data")

	InvalidFormatIsbnErr = errors.New("Isbn is invalid format, the format is xxxx-xxxx-xxxx")
	NotFoundReadbleErr   = errors.New("Not found Magazine either Book")
)
