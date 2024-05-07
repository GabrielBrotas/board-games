package impostor

import (
	"github.com/GabrielBrotas/board-games/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Player struct {
	User       *models.User
	IsImpostor bool
	Points     int
	InPlay     bool
	Conn       *websocket.Conn
}

func NewPlayer(user *models.User) (*Player, error) {
	return &Player{
		User:       user,
		IsImpostor: false,
		Points:     0,
		InPlay:     false,
	}, nil
}

func (p *Player) SetPoints(points int) {
	p.Points = points
}

func (p *Player) ResetPoints() {
	p.Points = 0
}

func (p *Player) SetImpostor() {
	p.IsImpostor = true
}

func (p *Player) UnsetImpostor() {
	p.IsImpostor = false
}

func (p *Player) SetInPlay() {
	p.InPlay = true
}

func (p *Player) UnsetInPlay() {
	p.InPlay = false
}

func (p *Player) UpdateConnection(conn *websocket.Conn) {
	p.Conn = conn
}

type PlayerOut struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Points int       `json:"points"`
}

func (p *Player) ToOut() *PlayerOut {
	return &PlayerOut{
		ID:     p.User.ID,
		Name:   p.User.Name,
		Points: p.Points,
	}
}
