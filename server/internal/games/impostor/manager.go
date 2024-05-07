package impostor

import (
	"log"
	"math/rand"
	"sort"

	"github.com/GabrielBrotas/board-games/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type ImpostorDistribution struct {
	One   int
	Two   int
	Three int
}

type GameManager struct {
	playerRepository *PlayerRepository
	gameStarted      bool
	word             string
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
		newPlayer, err := NewPlayer(user)
		if err != nil {
			return err
		}
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

	if player.Conn != nil {
		log.Printf("Removing player %s", player.User.Name)
		err := player.Conn.WriteJSON(map[string]interface{}{
			"type": "removedPlayer",
		})

		if err != nil {
			log.Printf("Error sending removed player: %v", err)
		}
	}
}

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
		if player == nil || player.Conn == nil {
			continue
		}

		err := player.Conn.WriteJSON(map[string]interface{}{
			"type":    "playerList",
			"players": playersOut,
		})

		if err != nil {
			log.Printf("Error broadcasting player list: %v", err)
		}
	}
}

// StartGame starts the game
func (gm *GameManager) StartGame(dist ImpostorDistribution, category string, difficulty string) {
	word, err := generateWordFromOpenAI(category, difficulty)

	if err != nil {
		log.Printf("Error fetching word from OpenAI: %v", err)
		return
	}

	numImpostors := chooseImpostorsNumber(dist)

	// Create a slice of player IDs to properly select random players.
	playerIDs := make([]uuid.UUID, 0, gm.playerRepository.GetActiveUsersCount())
	playerList := gm.playerRepository.GetPlayerList(false)
	for _, p := range playerList {
		playerIDs = append(playerIDs, p.User.ID)
	}

	chosen := make(map[uuid.UUID]bool)

	for i := 0; i < numImpostors; i++ {
		for {
			// Randomly select a player index.
			idx := rand.Intn(len(playerIDs))
			playerID := playerIDs[idx]
			if !chosen[playerID] {
				player := gm.playerRepository.GetPlayerByID(playerID)
				player.SetImpostor()
				chosen[playerID] = true
				break
			}
		}
	}

	// Send the role information to all players.
	for _, player := range playerList {
		// set user in play
		player.SetInPlay()
		message := word
		if player.IsImpostor {
			message = "Você é o impostor!"
		}
		if player.Conn == nil {
			continue
		}
		err := player.Conn.WriteJSON(map[string]interface{}{
			"type":       "role",
			"wordOrRole": message,
		})
		if err != nil {
			log.Printf("Error sending role: %v", err)
		}
	}

	gm.gameStarted = true
	gm.word = word
}

func (gm *GameManager) BroadcastImpostorsNumber() {
	players := gm.playerRepository.GetPlayerList(false)
	impostorsNumber := 0
	for _, player := range players {
		if player.InPlay && player.IsImpostor {
			impostorsNumber++
		}
	}

	for _, player := range players {
		if player.User == nil || player.Conn == nil {
			continue
		}

		err := player.Conn.WriteJSON(map[string]interface{}{
			"type":            "impostorsNumber",
			"impostorsNumber": impostorsNumber,
		})
		if err != nil {
			log.Printf("Error sending spies number: %v", err)
		}
	}
}

func (gm *GameManager) ResetGame() {
	players := gm.playerRepository.GetPlayerList(false)

	for _, player := range players {
		player.UnsetImpostor()
		player.UnsetInPlay()
		if player.Conn != nil {
			err := player.Conn.WriteJSON(map[string]interface{}{
				"type": "resetGame",
			})
			if err != nil {
				log.Printf("Error sending reset: %v", err)
			}
		}
	}

	gm.gameStarted = false
	gm.word = ""
}

func (gm *GameManager) UpdatePoints(impostorsWin bool) {
	log.Printf("Updating points")
	players := gm.playerRepository.GetPlayerList(false)
	for _, player := range players {
		if player.InPlay {
			log.Printf("Updating points to player %s", player.User.Name)
			points := player.Points
			log.Printf("Player %s has %d points", player.User.Name, points)

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

			log.Printf("Player %s now has %d points", player.User.Name, points)
			gm.playerRepository.UpdatePlayerPoints(player.User.ID, points)
		}
	}
}

func (gm *GameManager) BroadcastWinner(impostorsWin bool) {
	log.Printf("Broadcasting winner")
	players := gm.playerRepository.GetPlayerList(false)
	for _, player := range players {
		if player.Conn == nil {
			continue
		}

		log.Printf("Broadcasting winner to %s", player.User.Name)
		err := player.Conn.WriteJSON(map[string]interface{}{
			"type":         "winner",
			"impostorsWon": impostorsWin,
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

type GameStatus struct {
	GameStarted bool   `json:"gameStarted"`
	WordOrRole  string `json:"word"`
	InGame      bool   `json:"inGame"`
}

func (gm *GameManager) GetGameStatus(playerID uuid.UUID) *GameStatus {
	players := gm.playerRepository.GetPlayerList(false)
	for _, player := range players {
		if player.User.ID == playerID {
			if player.IsImpostor {
				return &GameStatus{
					GameStarted: gm.gameStarted,
					WordOrRole:  "Você é o impostor!",
					InGame:      true,
				}
			}
			return &GameStatus{
				GameStarted: gm.gameStarted,
				WordOrRole:  gm.word,
				InGame:      true,
			}
		}
	}

	return &GameStatus{
		GameStarted: gm.gameStarted,
		WordOrRole:  "",
		InGame:      false,
	}
}

func chooseImpostorsNumber(dist ImpostorDistribution) int {
	roll := rand.Intn(100)
	if roll < dist.One {
		return 1
	} else if roll < dist.One+dist.Two {
		return 2
	}
	return 3
}
