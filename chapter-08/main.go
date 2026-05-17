package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/vitostamatti/httpgordle/internal/handlers"
	"github.com/vitostamatti/httpgordle/internal/repository"
)

func main() {
	db := repository.New()

	err := http.ListenAndServe(":8080", handlers.NewRouter(db))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
