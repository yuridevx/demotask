package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"os/signal"
	gamev1 "yuridev.com/googdemo/domain/proto/game/v1"
	"yuridev.com/googdemo/game/pkg/game"
	"yuridev.com/googdemo/game/pkg/random"
)

// TODO I would add CLI and configuration manager here, for now it's just a hardcode
// https://github.com/spf13/viper
// https://github.com/spf13/cobra
func main() {
	appCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	lcnf := &net.ListenConfig{}
	lis, err := lcnf.Listen(appCtx, "tcp", ":3000")
	if err != nil {
		panic(err)
	}

	// use 2 point of precision
	gen := random.NewDefaultGenerator(2)
	srv := game.NewGameServer(gen)

	grpcServer := grpc.NewServer()

	go func() {
		<-appCtx.Done()
		grpcServer.GracefulStop()
	}()

	gamev1.RegisterGameServiceServer(grpcServer, srv)
	err = grpcServer.Serve(lis)
	if err != nil {
		panic(err)
	}
	log.Println("server exited gracefully")
}
