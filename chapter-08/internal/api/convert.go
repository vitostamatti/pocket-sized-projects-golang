package api

import (
	"github.com/vitostamatti/httpgordle/internal/session"
)

func ToGameResponse(g session.Game) GameResponse {
	apiGame := GameResponse{
		ID:           string(g.ID),
		AttemptsLeft: g.AttemptsLeft,
		Guesses:      make([]Guess, len(g.Guesses)),
		// TODO WordLength
	}
	for index := 0; index < len(g.Guesses); index++ {
		apiGame.Guesses[index].Word = g.Guesses[index].Word
		apiGame.Guesses[index].Feedback = g.Guesses[index].Feedback
	}
	if g.AttemptsLeft == 0 {
		apiGame.Solution = "" // TODO solution
	}
	return apiGame
}
