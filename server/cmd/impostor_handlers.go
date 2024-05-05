package main

import (
	"fmt"
	"log"
	"net/http"

	game "github.com/GabrielBrotas/board-games/internal/games/imposter"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Consider validating against a list of approved origins
		},
	}
	gameManager = game.NewGameManager(game.NewPlayerRepository())
)

// handleGetPlayerList retrieves the list of players.
func handleGetPlayerList(w http.ResponseWriter, r *http.Request) {
	players := gameManager.GetPlayerList(true)
	playerList := make([]game.PlayerOut, 0, len(players))

	for _, player := range players {
		playerList = append(playerList, *player.ToOut())
	}

	respondWithJSON(w, http.StatusOK, map[string]interface{}{"players": playerList})
}

// handleConnections upgrades HTTP to WebSocket and handles the connection.
func handleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading WebSocket connection: %v", err)
		return
	}
	defer conn.Close()
	defer handleCloseConnection(conn)

	messageLoop(conn)
}

func messageLoop(conn *websocket.Conn) {
	for {
		msg := make(map[string]interface{})
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("Error reading JSON: %v", err)
			break
		}

		if err := processMessage(conn, msg); err != nil {
			log.Printf("Error processing message: %v", err)
			break
		}
	}
}

func handleCloseConnection(conn *websocket.Conn) {
	log.Printf("Closing connection...")
	usersRepository.RemoveUserConn(conn)
	gameManager.BroadcastPlayerList()
}

func processMessage(conn *websocket.Conn, msg map[string]interface{}) error {
	switch msg["type"].(string) {
	case "connected":
		return handleConnected(conn, msg)
	case "startGame":
		return handleStartGame(msg)
	case "resetGame":
		gameManager.ResetGame()
	case "removePlayer":
		return handleRemovePlayer(msg) // v
	case "decideWinner":
		return handleDecideWinner(msg)
	case "resetPoints":
		gameManager.ResetPoints()
		gameManager.BroadcastPlayerList()
		return nil
	default:
		return fmt.Errorf("unknown message type: %v", msg["type"])
	}
	return nil
}

func handleConnected(conn *websocket.Conn, msg map[string]interface{}) error {
	userID, err := parseUUID(msg, "id")
	if err != nil {
		return err
	}

	user, err := usersRepository.GetUser(userID)
	if err != nil {
		return err
	}

	err = gameManager.RegisterPlayer(conn, user)
	if err != nil {
		return err
	}

	gameManager.BroadcastPlayerList()
	return nil
}

// handleStartGame initializes a new game with specified settings.
func handleStartGame(msg map[string]interface{}) error {
	dist := parseDistribution(msg)
	category, _ := msg["category"].(string)
	difficulty, _ := msg["difficulty"].(string)
	gameManager.StartGame(dist, category, difficulty)
	return nil
}

// handleRemovePlayer removes a player from the game.
func handleRemovePlayer(msg map[string]interface{}) error {
	userID, err := parseUUID(msg, "id")
	if err != nil {
		return err
	}
	gameManager.RemovePlayerByID(userID)
	gameManager.BroadcastPlayerList()
	return nil
}

// handleDecideWinner updates the game state based on the winner
func handleDecideWinner(msg map[string]interface{}) error {
	impostorWon, _ := msg["impostorWon"].(bool)
	gameManager.UpdatePoints(impostorWon)
	gameManager.BroadcastPlayerList()
	gameManager.ResetGame()
	gameManager.BroadcastWinner(impostorWon)
	return nil
}

// handleGetGameStatus returns the current game status for a user.
func handleGetGameStatus(w http.ResponseWriter, r *http.Request) {
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

	gameInfo := gameManager.GetGameStatus(user.ID)

	respondWithJSON(w, http.StatusOK, gameInfo)
}

// parseDistribution extracts game distribution settings from a message.
func parseDistribution(msg map[string]interface{}) game.ImposterDistribution {
	return game.ImposterDistribution{
		One:   int(msg["one"].(float64)),
		Two:   int(msg["two"].(float64)),
		Three: int(msg["three"].(float64)),
	}
}
