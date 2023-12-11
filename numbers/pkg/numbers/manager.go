package numbers

import (
	"context"
	"fmt"
	"github.com/shopspring/decimal"
	"sync"
)

type GameManager interface {
	IncrementGame(ctx context.Context, id string, decimal string) (*Game, error)
	GetGame(ctx context.Context, id string) (*Game, error)
}

type InMemoryManager struct {
	sync.RWMutex
	games map[string]*Game
}

func NewInMemoryStore() *InMemoryManager {
	return &InMemoryManager{
		games: make(map[string]*Game),
	}
}

func (s *InMemoryManager) GetGame(_ context.Context, id string) (*Game, error) {
	s.Lock()
	defer s.Unlock()
	game, ok := s.games[id]
	if !ok {
		game = &Game{
			ID:            id,
			Total:         "0",
			LastIncrement: "0",
		}
	}
	s.games[id] = game
	return game, nil
}

func (s *InMemoryManager) IncrementGame(_ context.Context, id string, dec string) (*Game, error) {
	s.Lock()
	defer s.Unlock()
	game, ok := s.games[id]
	if !ok {
		game = &Game{
			ID:            id,
			Total:         "0",
			LastIncrement: "0",
		}
	}
	val, err := decimal.NewFromString(game.Total)
	if err != nil {
		return nil, fmt.Errorf("failed to parse total: %w", err)
	}
	valInc, err := decimal.NewFromString(dec)
	if err != nil {
		return nil, fmt.Errorf("failed to parse increment: %w", err)
	}
	game.Total = val.Add(valInc).String()
	game.LastIncrement = valInc.String()
	return game, nil
}
