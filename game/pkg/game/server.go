package game

import (
	"context"
	corev1 "yuridev.com/googdemo/domain/proto/core/v1"
	gamev1 "yuridev.com/googdemo/domain/proto/game/v1"
	"yuridev.com/googdemo/game/pkg/random"
)

type GameServer struct {
	gamev1.UnimplementedGameServiceServer
	generator random.Generator
}

func (g *GameServer) RandomNumber(ctx context.Context, request *gamev1.RandomNumberRequest) (*gamev1.RandomNumberResponse, error) {
	number, err := g.generator.GetRandom()
	if err != nil {
		return nil, err
	}
	return &gamev1.RandomNumberResponse{
		Random: &corev1.RandomNumber{
			Number: number,
		},
	}, nil
}

func NewGameServer(
	generator random.Generator,
) *GameServer {
	return &GameServer{
		generator: generator,
	}
}

var (
	_ gamev1.GameServiceServer = (*GameServer)(nil)
)
