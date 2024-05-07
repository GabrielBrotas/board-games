package spyfall

import (
	"github.com/GabrielBrotas/board-games/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Player struct {
	User     *models.User
	InPlay   bool
	Role     string // "spy" or a profession
	Location string
	Points   int
	Conn     *websocket.Conn
}

func NewPlayer(user *models.User) *Player {
	return &Player{
		User:     user,
		InPlay:   false,
		Role:     "",
		Location: "",
	}
}

func (p *Player) SetPoints(points int) {
	p.Points = points
}

func (p *Player) ResetPoints() {
	p.Points = 0
}

func (p *Player) SetInPlay() {
	p.InPlay = true
}

func (p *Player) UnsetInPlay() {
	p.InPlay = false
}

func (p *Player) SetRole(role string) {
	p.Role = role
}

func (p *Player) SetSpy() {
	p.Role = "spy"
}

func (p *Player) SetLocation(location string) {
	p.Location = location
}

func (p *Player) UpdateConnection(conn *websocket.Conn) {
	p.Conn = conn
}

func (p *Player) IsSpy() bool {
	return p.Role == "spy"
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
