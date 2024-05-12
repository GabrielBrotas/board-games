package spyfall

import (
	"log"
	"math/rand"
	"sort"

	"github.com/GabrielBrotas/board-games/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type SpiesDistribution struct {
	One   int
	Two   int
	Three int
}

type GameManager struct {
	playerRepository *PlayerRepository
	gameStarted      bool
}

func NewGameManager(playerRepository *PlayerRepository) *GameManager {
	return &GameManager{
		playerRepository: playerRepository,
		gameStarted:      false,
	}
}

// RegisterPlayer registers a player in the game
func (gm *GameManager) RegisterPlayer(conn *websocket.Conn, user *models.User) error {
	player := gm.playerRepository.GetPlayerByID(user.ID)

	if player == nil {
		newPlayer := NewPlayer(user)
		gm.playerRepository.AddPlayer(newPlayer)
		player = newPlayer
	}

	player.UpdateConnection(conn)
	return nil
}

// RemovePlayerByID removes a player from the game by its ID
func (gm *GameManager) RemovePlayerByID(id uuid.UUID) {
	player := gm.playerRepository.GetPlayerByID(id)
	gm.playerRepository.RemovePlayerByID(id)

	if player.User != nil && player.Conn != nil {
		log.Printf("Removing player %s", player.User.Name)
		err := player.Conn.WriteJSON(map[string]interface{}{
			"type": "removedPlayer",
		})

		if err != nil {
			log.Printf("Error sending removed player: %v", err)
		}
	}
}

// RemoveConnection removes a connection from a player
func (gm *GameManager) RemoveConnection(conn *websocket.Conn) {
	player := gm.playerRepository.GetPlayerByConnection(conn)
	if player == nil {
		return
	}

	player.UpdateConnection(nil)
}

// BroadcastPlayerList broadcasts the player list to all players
func (gm *GameManager) BroadcastPlayerList() {
	players := gm.playerRepository.GetPlayerList(false)
	playersOut := make([]*PlayerOut, 0, len(players))
	for _, player := range players {
		playersOut = append(playersOut, player.ToOut())
	}

	// sort players by points
	sort.Slice(playersOut, func(i, j int) bool {
		return playersOut[i].Points > playersOut[j].Points
	})

	for _, player := range players {
		if player == nil || player.User == nil {
			continue
		}

		if player.User != nil && player.Conn != nil {
			err := player.Conn.WriteJSON(map[string]interface{}{
				"type":    "playerList",
				"players": playersOut,
			})

			if err != nil {
				log.Printf("Error broadcasting player list: %v", err)
			}
		}
	}
}

// StartGame starts the game
func (gm *GameManager) StartGame(dist SpiesDistribution) {
	// 1 - generate a random location and a role for the amount of players
	data := generateLocationAndRoles()

	// 2 - select the amount of spies
	spiesNumber := chooseSpiesNumber(dist)

	// 3 - Create a slice of player IDs to properly select random players.
	playerIDs := make([]uuid.UUID, 0, gm.playerRepository.GetActiveUsersCount())
	playerList := gm.playerRepository.GetPlayerList(false)
	for _, p := range playerList {
		playerIDs = append(playerIDs, p.User.ID)
	}

	chosen := make(map[uuid.UUID]bool)

	for i := 0; i < spiesNumber; i++ {
		for {
			// Randomly select a player index.
			idx := rand.Intn(len(playerIDs))
			playerID := playerIDs[idx]
			if !chosen[playerID] {
				player := gm.playerRepository.GetPlayerByID(playerID)
				player.SetSpy()
				chosen[playerID] = true
				break
			}
		}
	}

	// 4 - broadcast the location and role to each player
	for i, player := range playerList {
		player.SetInPlay()
		if player.IsSpy() {
			player.SetLocation("")
		} else {
			player.SetLocation(data.Location)
			if i < len(data.Roles) {
				player.SetRole(data.Roles[i])
			} else {
				player.SetRole(getRandomRole(data.Roles))
			}
		}

		if player.Conn == nil {
			continue
		}

		err := player.Conn.WriteJSON(map[string]interface{}{
			"type":     "role",
			"location": player.Location,
			"role":     player.Role,
		})

		if err != nil {
			log.Printf("Error sending role: %v", err)
		}
	}

	gm.gameStarted = true
}

func (gm *GameManager) BroadcastSpiesNumber() {
	players := gm.playerRepository.GetPlayerList(false)
	spiesNumber := 0
	for _, player := range players {
		if player.InPlay && player.IsSpy() {
			spiesNumber++
		}
	}

	for _, player := range players {
		if player.User == nil || player.Conn == nil {
			continue
		}

		err := player.Conn.WriteJSON(map[string]interface{}{
			"type":        "spiesNumber",
			"spiesNumber": spiesNumber,
		})
		if err != nil {
			log.Printf("Error sending spies number: %v", err)
		}
	}
}

// FinishGame finishes the game
func (gm *GameManager) FinishGame(spiesWon bool) {
	gm.updatePoints(spiesWon)
	gm.ResetGame()
	gm.broadcastWinner(spiesWon)
}

func (gm *GameManager) ResetGame() {
	players := gm.playerRepository.GetPlayerList(false)

	for _, player := range players {
		player.SetRole("")
		player.SetLocation("")
		player.UnsetInPlay()
		if player.User != nil && player.Conn != nil {
			err := player.Conn.WriteJSON(map[string]interface{}{
				"type": "resetGame",
			})
			if err != nil {
				log.Printf("Error sending reset: %v", err)
			}
		}
	}

	gm.gameStarted = false
}

func (gm *GameManager) updatePoints(spiesWon bool) {
	players := gm.playerRepository.GetPlayerList(false)
	for _, player := range players {
		if player.InPlay {
			points := player.Points

			if player.IsSpy() {
				if spiesWon {
					points += 20
				} else {
					points += 5
				}
			} else {
				if spiesWon {
					points += 5
				} else {
					points += 10
				}
			}
			gm.playerRepository.UpdatePlayerPoints(player.User.ID, points)
		}
	}
}

func (gm *GameManager) broadcastWinner(spiesWon bool) {
	players := gm.playerRepository.GetPlayerList(false)
	for _, player := range players {
		if player.User == nil || player.Conn == nil {
			continue
		}

		err := player.Conn.WriteJSON(map[string]interface{}{
			"type":   "winner",
			"spyWon": spiesWon,
		})
		if err != nil {
			log.Printf("Error sending winner: %v", err)
		}
	}
}

func (gm *GameManager) ResetPoints() {
	players := gm.playerRepository.GetPlayerList(false)
	for _, player := range players {
		gm.playerRepository.UpdatePlayerPoints(player.User.ID, 0)
	}
}

func (gm *GameManager) GetPlayerList(all bool) []*Player {
	return gm.playerRepository.GetPlayerList(all)
}

type playerGameStatus struct {
	GameStarted bool   `json:"gameStarted"`
	InGame      bool   `json:"inGame"`
	Location    string `json:"location"`
	Role        string `json:"role"`
}

func (gm *GameManager) GetPlayerGameStatus(playerID uuid.UUID) *playerGameStatus {
	players := gm.playerRepository.GetPlayerList(false)

	for _, player := range players {
		if player.User.ID == playerID {
			if player.IsSpy() {
				return &playerGameStatus{
					GameStarted: gm.gameStarted,
					InGame:      true,
					Location:    "",
					Role:        "spy",
				}
			}
			return &playerGameStatus{
				GameStarted: gm.gameStarted,
				InGame:      true,
				Location:    player.Location,
				Role:        player.Role,
			}
		}
	}

	return &playerGameStatus{
		GameStarted: gm.gameStarted,
		InGame:      false,
		Location:    "",
		Role:        "",
	}
}

func chooseSpiesNumber(dist SpiesDistribution) int {
	roll := rand.Intn(100) // select a number between 0 and 100
	if roll < dist.One {
		return 1
	} else if roll < dist.One+dist.Two {
		return 2
	}
	return 3
}
