package main

import (
	"github.com/gorilla/websocket"
	"github.com/google/uuid"
	"math/rand"
	"fmt"
)

type Player struct {
	ID         string
	Name       string
	Conn       *websocket.Conn
	IsImpostor bool
}

func NewPlayer(conn *websocket.Conn) *Player {
	return &Player{
		ID:         generatePlayerID(),
		Name:       generatePlayerName(),
		Conn:       conn,
		IsImpostor: false,
	}
}

func generatePlayerID() string {
	return "Player-" + uuid.New().String()
}

func generatePlayerName() string {
	return "Player-" + fmt.Sprint(rand.Intn(100))
}
