package pkg

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"yuridev.com/googdemo/numbers/pkg/numbers"
)

type Server struct {
	server *http.Server
	g      *gin.Engine
	ctrl   *numbers.Controller
}

func (s *Server) GetGame(c *gin.Context) {
	id := c.Param("id")
	game, err := s.ctrl.GetGame(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, game)
}

func (s *Server) PlayGame(c *gin.Context) {
	id := c.Param("id")
	game, err := s.ctrl.Increment(c, id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, game)
}

func (s *Server) Start(ctx context.Context) {
	go func() {
		fmt.Println("starting server on", s.server.Addr)
		err := s.server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
}

func (s *Server) Stop(ctx context.Context) {
	err := s.server.Shutdown(ctx)
	if err != nil {
		panic(err)
	}
}

func NewServer(
	ctrl *numbers.Controller,
	addr string,
) *Server {
	s := &Server{
		g:    gin.New(),
		ctrl: ctrl,
	}
	s.server = &http.Server{
		Addr:    addr,
		Handler: s.g,
	}
	s.g.GET("/api/game/:id", s.GetGame)
	s.g.POST("/api/game/:id", s.PlayGame)
	return s
}
