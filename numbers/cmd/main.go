package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
	"os/signal"
	"time"
	gamev1 "yuridev.com/googdemo/domain/proto/game/v1"
	"yuridev.com/googdemo/numbers/pkg"
	"yuridev.com/googdemo/numbers/pkg/numbers"
)

var gameHost string
var addr string

func parseFlags() {
	flag.StringVar(&gameHost, "game-host", "", "host of the game service")
	flag.StringVar(&addr, "addr", ":8080", "address to listen on")
	flag.Parse()
	if gameHost == "" {
		panic("game-host flag is required")
	}
}

func main() {
	parseFlags()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	conn, err := grpc.DialContext(ctx, gameHost,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		panic(err)
	}
	client := gamev1.NewGameServiceClient(conn)
	manager := numbers.NewInMemoryStore()
	ctrl := numbers.NewController(manager, client)
	server := pkg.NewServer(ctrl, addr)
	server.Start(ctx)

	<-ctx.Done()
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	server.Stop(shutdownCtx)
	fmt.Println("server stopped gracefully")
}
