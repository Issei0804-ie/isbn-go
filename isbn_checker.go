package isbngo

import (
	"github.com/Issei0804-ie/isbn-go/isbn10"
	"github.com/Issei0804-ie/isbn-go/isbn13"
)

func IsValidIsbn(maybeIsbn string) bool {
	return isbn10.IsValid() || isbn13.IsValid()
}
