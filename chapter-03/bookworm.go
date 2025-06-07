package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type Bookworm struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Author string `json:"author"`
	Title  string `json:"title"`
}

func loadBookworms(filePath string) ([]Bookworm, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var bookworms []Bookworm

	// err = json.NewDecoder(f).Decode(&bookworms)
	buffer := bufio.NewReaderSize(f, 1024*1024)
	decoder := json.NewDecoder(buffer)
	err = decoder.Decode(&bookworms)

	if err != nil {
		return nil, err
	}
	return bookworms, nil
}

func booksCount(bookworms []Bookworm) map[Book]uint {
	count := make(map[Book]uint)
	for _, bookworm := range bookworms {
		for _, book := range bookworm.Books {
			count[book]++
		}
	}
	return count
}

func findCommonBooks(bookworms []Bookworm) []Book {
	booksOnShelves := booksCount(bookworms)
	var commonBooks []Book
	for book, count := range booksOnShelves {
		if count > 1 {
			commonBooks = append(commonBooks, book)
		}
	}
	return sortBooks(commonBooks)
}

func sortBooks(books []Book) []Book {
	// sort.Slice(books, func(i, j int) bool {
	// 	if books[i].Author != books[j].Author {
	// 		return books[i].Author < books[j].Author
	// 	}
	// 	return books[i].Title < books[j].Title
	// })
	sort.Sort(byAuthor(books))
	return books
}

func displayBooks(books []Book) {
	for _, book := range books {
		fmt.Println("-", book.Title, "by", book.Author)
	}
}

type byAuthor []Book

func (b byAuthor) Len() int { return len(b) }
func (b byAuthor) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b byAuthor) Less(i, j int) bool {
	if b[i].Author != b[j].Author {
		return b[i].Author < b[j].Author
	}
	return b[i].Title < b[j].Title
}
