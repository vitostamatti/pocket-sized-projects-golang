package session

import "errors"

type GameID string
type Status string

const (
	StatusPlayin = "Playing"
	StatusWon    = "Won"
	StatusLost   = "Lost"
)

type Guess struct {
	Word     string
	Feedback string
}
type Game struct {
	ID           GameID
	AttemptsLeft byte
	Guesses      []Guess
	Status       Status
}

var ErrGameOver = errors.New("game over")
