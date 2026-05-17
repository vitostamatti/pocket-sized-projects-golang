package guess

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
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		g := api.GuessRequest{}
		err := json.NewDecoder(req.Body).Decode(&g)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id := req.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing game id", http.StatusBadRequest)
			return
		}

		game, err := guess(id, g, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		apiGame := api.ToGameResponse(game)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err = json.NewEncoder(w).Encode(apiGame)
		if err != nil {
			log.Printf("Failed to write response: %s", err)
		}
	}
}

func guess(id string, _ api.GuessRequest, _ *repository.GameRepository) (session.Game, error) {
	return session.Game{
		ID: session.GameID(id),
	}, nil
}
