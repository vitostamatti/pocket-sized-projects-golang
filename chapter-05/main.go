package main

import (
	"github.com/vitostamatti/gordle/gordle"
)

const maxAttempts = 6

func main() {
	corpus, err := gordle.ReadCorpus("./corpus/english.txt")
	if err != nil {
		panic(err)
	}
	g, err := gordle.New(corpus, gordle.WithMaxAttempts(maxAttempts))
	if err != nil {
		panic(err)
	}
	g.Play()
}
