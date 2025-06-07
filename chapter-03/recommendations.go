package main

import "sort"

type bookCollection map[Book]struct{}

type bookRecommendations map[Book]bookCollection

func newCollection() bookCollection {
	return make(bookCollection)
}

func recommendOtherBooks(bookworms []Bookworm) []Bookworm {
	sb := make(bookRecommendations)
	for _, bookworm := range bookworms {
		for i, book := range bookworm.Books {
			otherBooksOnShelves := listOtherBooksOnShelves(i, bookworm.Books)
			registerBookRecommendations(sb, book, otherBooksOnShelves)
		}
	}

	recommendations := make([]Bookworm, len(bookworms))
	for i, bookworm := range bookworms {
		recommendations[i] = Bookworm{
			Name:  bookworm.Name,
			Books: recommendBooks(sb, bookworm.Books),
		}
	}

	return recommendations
}

func recommendBooks(recommendations bookRecommendations, books []Book) []Book {
	bc := make(bookCollection)
	shelf := make(map[Book]bool)
	for _, book := range books {
		shelf[book] = true
	}
	for _, book := range books {
		for recommendation := range recommendations[book] {
			if shelf[recommendation] {
				continue
			}
			bc[recommendation] = struct{}{}
		}
	}
	recommendationsForBook := bookCollectionToListOfBooks(bc)
	return recommendationsForBook

}

func bookCollectionToListOfBooks(bc bookCollection) []Book {
	books := make([]Book, len(bc))
	for book := range bc {
		books = append(books, book)
	}
	sort.Slice(books, func(i, j int) bool {
		if books[i].Author != books[j].Author {
			return books[i].Author < books[j].Author
		}
		return books[i].Title < books[j].Title
	})
	return books
}

func registerBookRecommendations(recommendations bookRecommendations, book Book, otherBooksOnShelves []Book) {
	for _, book := range otherBooksOnShelves {
		collection, ok := recommendations[book]
		if !ok {
			collection = newCollection()
			recommendations[book] = collection
		}
		collection[book] = struct{}{}
	}

}

func listOtherBooksOnShelves(bookIndexToRemove int, books []Book) []Book {
	otherBooksOnShelves := make([]Book, len(books)-1)
	copy(otherBooksOnShelves, books[:bookIndexToRemove])
	otherBooksOnShelves = append(otherBooksOnShelves, books[bookIndexToRemove+1:]...)
	return otherBooksOnShelves
}
