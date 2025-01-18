package isbngo_test

import (
	"testing"

	isbngo "github.com/Issei0804-ie/isbn-go"
	"github.com/stretchr/testify/assert"
)

func TestIsValidIsbn13(t *testing.T) {
	actual, err := isbngo.IsValidIsbn13("9780470059029")

	assert.Equal(t, true, actual)
	assert.Equal(t, nil, err)

	// 最後のチェックディジットが9であるべきだが、0になっているためfalseが返ってくる
	actual, err = isbngo.IsValidIsbn13("9780470059020")

	assert.Equal(t, false, actual)
	assert.Equal(t, nil, err)
}
