package main

import (
	"log"
	"math/rand"
	"sort"

	"github.com/gorilla/websocket"
)

type ImposterDistribution struct {
	One   int
	Two   int
	Three int
}

type GameManager struct {
	players       map[string]*Player
	playerManager *PlayerManager
}

func NewGameManager(pm *PlayerManager) *GameManager {
	return &GameManager{
		players:       make(map[string]*Player),
		playerManager: pm,
	}
}

func (gm *GameManager) AddPlayer(player *Player) {
	gm.players[player.ID] = player
}

func (gm *GameManager) RemovePlayer(player *Player) {
	delete(gm.players, player.ID)
}

func (gm *GameManager) RemovePlayerByID(id string) {
	player := gm.players[id]
	if player == nil {
		return
	}

	if player.Conn != nil {
		err := player.Conn.WriteJSON(map[string]interface{}{
			"type": "removedPlayer",
		})

		if err != nil {
			log.Printf("Error sending removed player: %v", err)
		}
	}

	delete(gm.players, id)

	gm.BroadcastPlayerList()
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

// func (gm *GameManager) GetPlayerList() []*PlayerOutput {
// 	var playerList []*PlayerOutput
// 	for _, player := range gm.players {
// 		playerList = append(playerList, player.GetOutput(gm.playerManager))
// 	}
// 	return playerList
// }

func (gm *GameManager) GetPlayerList() []*PlayerOutput {
	var playerList []*PlayerOutput
	for _, player := range gm.players {
		playerList = append(playerList, player.GetOutput(gm.playerManager))
	}

	// Sort the player list by points.
	sort.Slice(playerList, func(i, j int) bool {
		return playerList[i].Points > playerList[j].Points
	})

	return playerList
}

func (gm *GameManager) ChangePlayerName(player *Player, newName string) {
	if _, exists := gm.players[player.ID]; exists {
		gm.players[player.ID].Name = newName
		gm.playerManager.AddUser(newName)
	}
}

func (gm *GameManager) StartGame(dist ImposterDistribution, category string, difficulty string) {
	word, err := generateWordFromOpenAI(category, difficulty)

	if err != nil {
		log.Printf("Error fetching word from OpenAI: %v", err)
		return
	}

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

func (gm *GameManager) UpdatePoints(impostorsWin bool) {
	for _, player := range gm.players {
		points := gm.playerManager.GetUserPoints(player.Name)
		log.Printf("Player %s has %d points", player.Name, points)
		if player.IsImpostor { // impostor
			if impostorsWin {
				points += 20
			} else {
				points += 5
			}
		} else { // crewmate
			if impostorsWin {
				points += 5
			} else {
				points += 10
			}
		}
		log.Printf("Player %s now has %d points", player.Name, points)
		gm.playerManager.UpdatePoints(player.Name, points)
	}
}

func (gm *GameManager) ResetPoints() {
	gm.playerManager.ResetPoints()
	for _, player := range gm.players {
		player.Points = 0
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
