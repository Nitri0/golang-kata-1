package domain

import "errors"

var (
	InvalidFormatIsbnErr = errors.New("Isbn is invalid format, the format is xxxx-xxxx-xxxx")
	NotFoundReadbleErr   = errors.New("Not found Magazine either Book")
)
