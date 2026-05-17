package newgame

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/vitostamatti/httpgordle/internal/api"
	"github.com/vitostamatti/httpgordle/internal/repository"
	"github.com/vitostamatti/httpgordle/internal/session"
)

func Handle(db *repository.GameRepository) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		game, err := createGame(db)
		if err != nil {
			log.Printf("unable to create a new game: %s", err)
			http.Error(w, "failed to create a new game", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)

		apiGame := response(game)
		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	}
}

func createGame(_ *repository.GameRepository) (session.Game, error) {
	return session.Game{}, nil
}

func response(game session.Game) api.GameResponse {
	return api.GameResponse{}
}
