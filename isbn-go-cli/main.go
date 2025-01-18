package main

import (
	"fmt"

	isbngo "github.com/Issei0804-ie/isbn-go"
)

func main() {
	hoge := isbngo.IsValidIsbn("0471958697")
	fmt.Println(hoge)
}
