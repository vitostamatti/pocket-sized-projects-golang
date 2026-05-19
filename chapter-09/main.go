package main

import (
	"fmt"
	"log"
	"os"

	"github.com/vitostamatti/maze-solver/internal/solver"
)

func main() {
	if len(os.Args) != 3 {
		usage()
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	_, err := solver.New(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	log.Printf("Solving maze %q and saving it as %q", inputFile, outputFile)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: maze_solver input.png output.png\n")
	os.Exit(1)
}
