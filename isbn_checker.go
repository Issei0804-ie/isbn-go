package isbngo

func IsValidIsbn(maybeIsbn string) bool {
	IsValidIsbn10, _ := IsValidIsbn10(maybeIsbn)
	IsValidIsbn13, _ := IsValidIsbn13(maybeIsbn)
	return IsValidIsbn10 || IsValidIsbn13
}
