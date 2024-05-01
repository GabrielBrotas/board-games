package main

import (
	"log"
	"math/rand"

	"github.com/gorilla/websocket"
)

type ImposterDistribution struct {
	One   int
	Two   int
	Three int
}

type GameManager struct {
	players map[string]*Player
}

func NewGameManager() *GameManager {
	return &GameManager{
		players: make(map[string]*Player),
	}
}

func (gm *GameManager) AddPlayer(player *Player) {
	gm.players[player.ID] = player
}

func (gm *GameManager) RemovePlayer(player *Player) {
	delete(gm.players, player.ID)
}

func (gm *GameManager) SendPlayerList(conn *websocket.Conn) {
	playerList := gm.GetPlayerList()
	err := conn.WriteJSON(map[string]interface{}{
		"type":    "playerList",
		"players": playerList,
	})
	if err != nil {
		log.Printf("Error sending player list: %v", err)
	}
}

func (gm *GameManager) BroadcastPlayerList() {
	playerList := gm.GetPlayerList()
	for _, player := range gm.players {
		err := player.Conn.WriteJSON(map[string]interface{}{
			"type":    "playerList",
			"players": playerList,
		})
		if err != nil {
			log.Printf("Error broadcasting player list: %v", err)
		}
	}
}

func (gm *GameManager) ResetGame() {
	for _, player := range gm.players {
		player.IsImpostor = false
		err := player.Conn.WriteJSON(map[string]interface{}{
			"type": "resetGame",
		})
		if err != nil {
			log.Printf("Error sending reset: %v", err)
		}
	}
}

func (gm *GameManager) GetPlayerList() []string {
	var playerList []string
	for _, player := range gm.players {
		playerList = append(playerList, player.Name)
	}
	return playerList
}

func (gm *GameManager) ChangePlayerName(player *Player, newName string) {
	if _, exists := gm.players[player.ID]; exists {
		gm.players[player.ID].Name = newName
	}
}

func (gm *GameManager) StartGame(dist ImposterDistribution) {
	word := words[rand.Intn(len(words))]
	numImposters := chooseImposters(dist)

	// Create a slice of player IDs to properly select random players.
	playerIDs := make([]string, 0, len(gm.players))
	for id := range gm.players {
		playerIDs = append(playerIDs, id)
	}

	chosen := make(map[string]bool)

	for i := 0; i < numImposters; i++ {
		for {
			// Randomly select a player index.
			idx := rand.Intn(len(playerIDs))
			playerID := playerIDs[idx]
			if !chosen[playerID] {
				gm.players[playerID].IsImpostor = true
				chosen[playerID] = true
				break
			}
		}
	}

	// Send the role information to all players.
	for _, player := range gm.players {
		message := word
		if player.IsImpostor {
			message = "You are the impostor"
		}
		err := player.Conn.WriteJSON(map[string]interface{}{
			"type":       "role",
			"wordOrRole": message,
		})
		if err != nil {
			log.Printf("Error sending role: %v", err)
		}
	}
}

func chooseImposters(dist ImposterDistribution) int {
	roll := rand.Intn(100)
	if roll < dist.One {
		return 1
	} else if roll < dist.One+dist.Two {
		return 2
	}
	return 3
}
