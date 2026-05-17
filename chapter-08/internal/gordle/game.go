package gordle

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

var errInvalidWordLength = fmt.Errorf("invalid guess, word doesn't have the same number of characters as the solution")

type Game struct {
	reader      *bufio.Reader
	solution    []rune
	maxAttempts int
}

func New(corpus []string, cfs ...ConfigFunc) (*Game, error) {
	if len(corpus) == 0 {
		return nil, ErrCorpusIsEmpty
	}
	g := &Game{
		reader:      bufio.NewReader(os.Stdin),
		solution:    splitToUppercaseCharacters(pickWord(corpus)),
		maxAttempts: -1,
	}
	for _, cf := range cfs {
		err := cf(g)
		if err != nil {
			return nil, fmt.Errorf("unable to apply config func: %w", err)
		}
	}
	return g, nil
}

func (g *Game) Play() {
	fmt.Println("Welcome to Gordle!")

	for currentAttempt := 1; currentAttempt <= g.maxAttempts; currentAttempt++ {
		guess := g.ask()
		fb := computeFeedback(guess, g.solution)
		fmt.Println(fb.String())
		if slices.Equal(guess, g.solution) {
			fmt.Printf("🎉 You won! You found it in %d guess(es)! The word was: %s.\n", currentAttempt, string(g.solution))
			return
		}
	}

	fmt.Printf("😔 You've lost! The solution was: %s\n", string(g.solution))
}

func (g *Game) ask() []rune {
	fmt.Printf("Enter  %d-character guess:\n", len(g.solution))
	for {
		playerInput, _, err := g.reader.ReadLine()
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Gordle failed to read your guess: %s\n", err.Error())
			continue
		}
		guess := splitToUppercaseCharacters(string(playerInput))
		err = g.validateGuess(guess)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Your attempt is invalid with Gordle's solution: %s. \n", err.Error())
		} else {
			return guess
		}
	}
}

func (g *Game) validateGuess(guess []rune) error {
	if len(guess) != len(g.solution) {
		return fmt.Errorf("expected %d, got %d, %w", len(g.solution), len(guess), errInvalidWordLength)
	}
	return nil
}

func splitToUppercaseCharacters(input string) []rune {
	return []rune(strings.ToUpper(input))
}

// computeFeedback verifies every character in the guess against the solution.
func computeFeedback(guess, solution []rune) feedback {
	result := make(feedback, len(guess))
	used := make([]bool, len(solution))
	if len(guess) != len(solution) {
		_, _ = fmt.Fprintf(os.Stderr, "Internal error. Guess and solution have different lengths: %d vs %d\n", len(guess), len(solution))
	}
	for posInGuess, character := range guess {
		if character == solution[posInGuess] {
			result[posInGuess] = correctPosition
			used[posInGuess] = true
		}
	}

	for posInGuess, character := range guess {
		if result[posInGuess] != absentCharacter {
			continue
		}
		for posInSolution, target := range solution {
			if used[posInSolution] {
				continue
			}

			if character == target {
				result[posInGuess] = wrongPosition
				used[posInSolution] = true
				break
			}
		}

	}
	return result
}
