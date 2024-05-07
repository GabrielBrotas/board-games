package impostor

import (
	"log"

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
	log.Printf("Getting player %s", id)
	player, ok := r.players[id]

	if !ok {
		log.Printf("Player not found: %s", id)
		return nil
	}

	return player
}

// GetPlayerByConnection retrieves a player by its connection.
func (r *PlayerRepository) GetPlayerByConnection(conn *websocket.Conn) *Player {
	for _, player := range r.players {
		if player.Conn == conn {
			return player
		}
	}
	return nil
}

func (r *PlayerRepository) AddPlayer(player *Player) {
	log.Printf("Adding player %s", player.User.Name)
	exists := r.GetPlayerByID(player.User.ID)

	if exists != nil {
		log.Printf("Player %s already exists", player.User.Name)
		return
	}

	r.players[player.User.ID] = player
	log.Printf("Player %s added", player.User.Name)
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
