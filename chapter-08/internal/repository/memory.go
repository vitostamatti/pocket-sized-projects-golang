package repository

import (
	"fmt"
	"sync"

	"github.com/vitostamatti/httpgordle/internal/session"
)

type GameRepository struct {
	mutex   sync.Mutex
	storage map[session.GameID]session.Game
}

func New() *GameRepository {
	return &GameRepository{
		storage: make(map[session.GameID]session.Game),
	}
}

func (gr *GameRepository) Add(game session.Game) error {
	gr.mutex.Lock()
	defer gr.mutex.Unlock()
	_, ok := gr.storage[game.ID]
	if ok {
		return fmt.Errorf("game with ID %s already exists", game.ID)
	}
	gr.storage[game.ID] = game
	return nil
}

func (gr *GameRepository) Find(id session.GameID) (session.Game, error) {
	gr.mutex.Lock()
	defer gr.mutex.Unlock()

	game, ok := gr.storage[id]
	if !ok {
		return session.Game{}, fmt.Errorf("game with ID %s does not exists", id)
	}
	return game, nil
}

func (gr *GameRepository) Update(game session.Game) error {

	gr.mutex.Lock()
	defer gr.mutex.Unlock()

	_, ok := gr.storage[game.ID]
	if !ok {
		return fmt.Errorf("game with ID %s does not exists", game.ID)
	}
	gr.storage[game.ID] = game
	return nil
}
