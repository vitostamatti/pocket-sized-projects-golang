package handlers

import (
	"net/http"

	"github.com/vitostamatti/httpgordle/internal/api"
	"github.com/vitostamatti/httpgordle/internal/handlers/getstatus"
	"github.com/vitostamatti/httpgordle/internal/handlers/guess"
	"github.com/vitostamatti/httpgordle/internal/handlers/newgame"
	"github.com/vitostamatti/httpgordle/internal/repository"
)

func NewRouter(db *repository.GameRepository) *http.ServeMux {
	r := http.NewServeMux()
	r.HandleFunc(http.MethodPost+" "+api.NewGameRoute, newgame.Handle(db))
	r.HandleFunc(http.MethodGet+" "+api.GetStatusRoute, getstatus.Handle(db))
	r.HandleFunc(http.MethodPut+" "+api.GessRoute, guess.Handle(db))
	return r
}
