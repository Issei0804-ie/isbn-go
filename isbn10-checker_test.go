package isbngo_test

import (
	"testing"

	isbngo "github.com/Issei0804-ie/isbn-go"
	"github.com/stretchr/testify/assert"
)

func TestIsValidIsbn10(t *testing.T) {
	actual, err := isbngo.IsValidIsbn10("0471958697")

	assert.Equal(t, true, actual)
	assert.Equal(t, nil, err)

	// 最後のチェックディジットが7であるべきだが、8になっているためfalseが返ってくる
	actual, err = isbngo.IsValidIsbn10("0471958698")

	assert.Equal(t, false, actual)
	assert.Equal(t, nil, err)
}
