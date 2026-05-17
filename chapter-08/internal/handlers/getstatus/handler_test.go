package getstatus

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/vitostamatti/httpgordle/internal/api"
	"github.com/vitostamatti/httpgordle/internal/repository"
	"github.com/vitostamatti/httpgordle/internal/session"
)

func TestHandle(t *testing.T) {
	req, err := http.NewRequest(
		http.MethodGet, "/games", nil,
	)
	require.NoError(t, err)
	gameID := "123456"
	req.SetPathValue(api.GameID, gameID)

	recorder := httptest.NewRecorder()

	db := repository.New()
	db.Add(session.Game{
		ID: session.GameID(gameID),
	})
	Handle(db)(recorder, req)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "application/json", recorder.Header().Get("Content-Type"))
	assert.JSONEq(t, `{"id":"123456","attempts_left":0,"guesses":[],"word_length":0,"status":""}`, recorder.Body.String())

}
