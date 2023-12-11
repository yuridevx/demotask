package integration

import (
	"github.com/go-resty/resty/v2"
	"math/rand"
	"strconv"
	"testing"
)

type Game struct {
	ID            string `json:"id"`
	Total         string `json:"total"`
	LastIncrement string `json:"lastIncrement"`
}

func TestIntegration(t *testing.T) {
	// TODO Use testify, but Github copilot works just fine to automate tedious assertions
	r := resty.New().SetBaseURL("http://localhost:8080")
	gameId := randomString(10)
	gameJson := &Game{}
	getGameResponse, err := r.R().SetResult(gameJson).Get("/api/game/" + gameId)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if getGameResponse.StatusCode() != 200 {
		t.Errorf("Expected 200, got %d", getGameResponse.StatusCode())
	}
	if gameJson.ID != gameId {
		t.Errorf("Expected %s, got %s", gameId, gameJson.ID)
	}
	if gameJson.Total != "0" {
		t.Errorf("Expected 0, got %s", gameJson.Total)
	}
	if gameJson.LastIncrement != "0" {
		t.Errorf("Expected 0, got %s", gameJson.LastIncrement)
	}

	incrementResponse, err := r.R().SetResult(gameJson).Post("/api/game/" + gameId)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if incrementResponse.StatusCode() != 200 {
		t.Errorf("Expected 200, got %d", incrementResponse.StatusCode())
	}
	if gameJson.ID != gameId {
		t.Errorf("Expected %s, got %s", gameId, gameJson.ID)
	}
	total, err := strconv.ParseFloat(gameJson.Total, 64)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	lastIncrement, err := strconv.ParseFloat(gameJson.LastIncrement, 64)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if total != lastIncrement {
		t.Errorf("Expected %f, got %f", total, lastIncrement)
	}
	if total == 0 {
		t.Errorf("Expected non-zero, got %f", total)
	}
}

func randomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	return string(b)
}
