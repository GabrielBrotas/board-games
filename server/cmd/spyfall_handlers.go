package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GabrielBrotas/board-games/internal/games/spyfall"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	spyfallUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Consider validating against a list of approved origins
		},
	}
	spyfallManager = spyfall.NewGameManager(spyfall.NewPlayerRepository())
)

func handleGetSpyfallPlayerList(w http.ResponseWriter, r *http.Request) {
	players := spyfallManager.GetPlayerList(true)
	playerList := make([]spyfall.PlayerOut, 0, len(players))

	for _, player := range players {
		playerList = append(playerList, *player.ToOut())
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{"players": playerList})
}

func handleSpyfallConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := spyfallUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading WebSocket connection: %v", err)
		return
	}
	defer conn.Close()
	defer handleCloseSpyfallConnection(conn)

	spyfallMessageLoop(conn)
}

func spyfallMessageLoop(conn *websocket.Conn) {
	log.Printf("New Spyfall connection: %v", conn.RemoteAddr())
	for {
		msg := make(map[string]interface{})
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("Error reading JSON from Spyfall client %v: %v", conn.RemoteAddr(), err)
			break
		}
		log.Printf("Received message at Spyfall: %v", msg)
		if err := processSpyfallMessage(conn, msg); err != nil {
			log.Printf("Error processing Spyfall message from %v: %v", conn.RemoteAddr(), err)
			break
		}
	}
}

func handleCloseSpyfallConnection(conn *websocket.Conn) {
	log.Printf("Closing Spyfall connection: %v", conn.RemoteAddr())
	spyfallManager.RemoveConnection(conn)
	spyfallManager.BroadcastPlayerList()
}

func processSpyfallMessage(conn *websocket.Conn, msg map[string]interface{}) error {
	switch msg["type"].(string) {
	case "connected":
		return handleSpyfallConnected(conn, msg)
	case "startGame":
		return handleSpyfallStartGame(msg)
	case "resetGame":
		spyfallManager.ResetGame()
	case "removePlayer":
		return handleSpyfallRemovePlayer(msg) // v
	case "decideWinner":
		return handleSpyfallDecideWinner(msg)
	case "resetPoints":
		spyfallManager.ResetPoints()
		spyfallManager.BroadcastPlayerList()
		return nil
	case "showSpiesNumber":
		spyfallManager.BroadcastSpiesNumber()
		return nil
	default:
		return fmt.Errorf("unknown message type: %v", msg["type"])
	}
	return nil
}

func handleSpyfallConnected(conn *websocket.Conn, msg map[string]interface{}) error {
	userID, err := parseUUID(msg, "id")
	if err != nil {
		return err
	}

	user, err := usersRepository.GetUser(userID)
	if err != nil {
		return err
	}

	err = spyfallManager.RegisterPlayer(conn, user)
	if err != nil {
		return err
	}

	spyfallManager.BroadcastPlayerList()
	return nil
}

// handleStartGame initializes a new game with specified settings.
func handleSpyfallStartGame(msg map[string]interface{}) error {
	dist := parseSpyfallDistribution(msg)
	spyfallManager.StartGame(dist)
	return nil
}

// handleSpyfallRemovePlayer removes a player from the game.
func handleSpyfallRemovePlayer(msg map[string]interface{}) error {
	userID, err := parseUUID(msg, "id")
	if err != nil {
		return err
	}
	spyfallManager.RemovePlayerByID(userID)
	spyfallManager.BroadcastPlayerList()
	return nil
}

// handleSpyfallDecideWinner updates the game state based on the winner
func handleSpyfallDecideWinner(msg map[string]interface{}) error {
	spyWon, _ := msg["spyWon"].(bool)
	spyfallManager.FinishGame(spyWon)
	spyfallManager.BroadcastPlayerList()
	return nil
}

// handleSpyfallGetPlayerGameStatus returns the current game status for a user.
func handleSpyfallGetPlayerGameStatus(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("u")
	userID, err := uuid.Parse(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := usersRepository.GetUser(userID)

	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	gameInfo := spyfallManager.GetPlayerGameStatus(user.ID)

	respondWithJSON(w, http.StatusOK, gameInfo)
}

func parseSpyfallDistribution(msg map[string]interface{}) spyfall.SpiesDistribution {
	return spyfall.SpiesDistribution{
		One:   int(msg["one"].(float64)),
		Two:   int(msg["two"].(float64)),
		Three: int(msg["three"].(float64)),
	}
}
