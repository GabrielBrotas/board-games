package spyfall

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// PlayerRepository is an in-memory storage for Player objects.
type PlayerRepository struct {
	players map[uuid.UUID]*Player
}

func NewPlayerRepository() *PlayerRepository {
	return &PlayerRepository{
		players: make(map[uuid.UUID]*Player),
	}
}

func (r *PlayerRepository) GetPlayerByID(id uuid.UUID) *Player {
	player, ok := r.players[id]

	if !ok {
		return nil
	}

	return player
}

func (r *PlayerRepository) GetPlayerByConnection(conn *websocket.Conn) *Player {
	for _, player := range r.players {
		if player.Conn == conn {
			return player
		}
	}
	return nil
}

func (r *PlayerRepository) AddPlayer(player *Player) {
	exists := r.GetPlayerByID(player.User.ID)

	if exists != nil {
		return
	}

	r.players[player.User.ID] = player
}

func (r *PlayerRepository) RemovePlayerByID(id uuid.UUID) {
	player := r.players[id]
	if player == nil {
		return
	}

	delete(r.players, id)
}

func (r *PlayerRepository) UpdatePlayerPoints(id uuid.UUID, points int) {
	player := r.GetPlayerByID(id)
	if player == nil {
		return
	}

	player.SetPoints(points)
}

func (r *PlayerRepository) GetPlayerList(all bool) []*Player {
	players := make([]*Player, 0, len(r.players))
	for _, player := range r.players {
		// all, or in the game, or connected to the room
		if all || player.InPlay || player.Conn != nil {
			players = append(players, player)
		}
	}
	return players
}

func (r *PlayerRepository) GetActiveUsersCount() int {
	activeUsers := 0
	for _, player := range r.players {
		if player.Conn != nil {
			activeUsers++
		}
	}
	return activeUsers
}
