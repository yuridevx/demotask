package numbers

import (
	"context"
	"fmt"
	gamev1 "yuridev.com/googdemo/domain/proto/game/v1"
)

type Controller struct {
	store  GameManager
	remote gamev1.GameServiceClient
}

func (c *Controller) GetGame(ctx context.Context, id string) (*Game, error) {
	game, err := c.store.GetGame(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get game: %w", err)
	}
	return game, nil
}

func (c *Controller) Increment(ctx context.Context, id string) (*Game, error) {
	rnd, err := c.remote.RandomNumber(ctx, &gamev1.RandomNumberRequest{})
	if err != nil {
		return nil, fmt.Errorf("failed to get random number: %w", err)
	}
	game, err := c.store.IncrementGame(ctx, id, rnd.Random.Number)
	if err != nil {
		return nil, fmt.Errorf("failed to increment game: %w", err)
	}
	return game, nil
}

func NewController(
	store GameManager,
	remote gamev1.GameServiceClient,
) *Controller {
	return &Controller{
		store:  store,
		remote: remote,
	}
}
