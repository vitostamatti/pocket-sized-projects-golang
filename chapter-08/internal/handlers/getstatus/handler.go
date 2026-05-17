package getstatus

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

		id := req.PathValue(api.GameID)
		if id == "" {
			http.Error(w, "missing game id", http.StatusBadRequest)
			return
		}
		game, err := db.Find(session.GameID(id))
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
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
