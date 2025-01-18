package main

import (
	"flag"
	"fmt"

	isbngo "github.com/Issei0804-ie/isbn-go"
)

func main() {
	flag.Parse()
	maybeIsbns := flag.Args()

	for _, maybeIsbn := range maybeIsbns {
		result := isbngo.IsValidIsbn(maybeIsbn)
		fmt.Println(result)
	}
}
