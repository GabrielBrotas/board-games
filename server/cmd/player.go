package main

import (
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Player struct {
	ID         string
	Name       string
	Conn       *websocket.Conn
	IsImpostor bool
	Points     int
	IsDeleted  bool
}

func NewPlayer(conn *websocket.Conn) *Player {
	return &Player{
		ID:         generatePlayerID(),
		Name:       generatePlayerName(),
		Conn:       conn,
		IsImpostor: false,
		Points:     0,
		IsDeleted:  false,
	}
}

func generatePlayerID() string {
	return "Player-" + uuid.New().String()
}

func generatePlayerName() string {
	return "Player-" + fmt.Sprint(rand.Intn(100))
}

type PlayerOutput struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
}

func (p *Player) GetOutput(playerManager *PlayerManager) *PlayerOutput {
	points := playerManager.GetUserPoints(p.Name)
	return &PlayerOutput{
		ID:     p.ID,
		Name:   p.Name,
		Points: points,
	}
}
