package main

import (
	"fmt"
	"log"
	"net/http"

	impostor "github.com/GabrielBrotas/board-games/internal/games/impostor"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var (
	impostorUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true // Consider validating against a list of approved origins
		},
	}
	impostorManager = impostor.NewGameManager(impostor.NewPlayerRepository())
)

// handleGetImpostorPlayerList retrieves the list of players.
func handleGetImpostorPlayerList(w http.ResponseWriter, r *http.Request) {
	log.Println("[handleGetImpostorPlayerList] Getting Impostor player list")
	players := impostorManager.GetPlayerList(true)
	playerList := make([]impostor.PlayerOut, 0, len(players))

	for _, player := range players {
		playerList = append(playerList, *player.ToOut())
	}

	log.Printf("[handleGetImpostorPlayerList] Returning player list: %v", playerList)
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"players": playerList})
}

// handleImpostorConnections upgrades HTTP to WebSocket and handles the connection.
func handleImpostorConnections(w http.ResponseWriter, r *http.Request) {
	log.Println("[handleImpostorConnections] Handling new Impostor connection")
	conn, err := impostorUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[handleImpostorConnections] Error upgrading WebSocket connection: %v", err)
		return
	}
	defer conn.Close()
	defer handleCloseImpostorConnection(conn)

	impostorMessageLoop(conn)
}

func impostorMessageLoop(conn *websocket.Conn) {
	log.Printf("[impostorMessageLoop] New Impostor connection: %v", conn.RemoteAddr())
	for {
		msg := make(map[string]interface{})
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("[impostorMessageLoop] Error reading JSON from Impostor client %v: %v", conn.RemoteAddr(), err)
			break
		}
		if err := processImpostorMessage(conn, msg); err != nil {
			log.Printf("[impostorMessageLoop] Error processing Impostor message from %v: %v", conn.RemoteAddr(), err)
			break
		}
	}
}

func handleCloseImpostorConnection(conn *websocket.Conn) {
	log.Printf("[handleCloseImpostorConnection] Closing Impostor connection: %v", conn.RemoteAddr())
	impostorManager.RemoveConnection(conn)
	impostorManager.BroadcastPlayerList()
}

func processImpostorMessage(conn *websocket.Conn, msg map[string]interface{}) error {
	log.Printf("[processImpostorMessage] Processing Impostor message: %v", msg)
	switch msg["type"].(string) {
	case "connected":
		return handleImpostorConnected(conn, msg)
	case "startGame":
		return handleImpostorStartGame(msg)
	case "resetGame":
		log.Println("[processImpostorMessage] Resetting game")
		impostorManager.ResetGame()
	case "removePlayer":
		return handleImpostorRemovePlayer(msg) // v
	case "decideWinner":
		return handleImpostorDecideWinner(msg)
	case "resetPoints":
		log.Println("[processImpostorMessage] Resetting points")
		impostorManager.ResetPoints()
		impostorManager.BroadcastPlayerList()
		return nil
	case "showImpostorsNumber":
		log.Println("[processImpostorMessage] Showing impostors number")
		impostorManager.BroadcastImpostorsNumber()
		return nil
	default:
		return fmt.Errorf("unknown message type: %v", msg["type"])
	}
	return nil
}

func handleImpostorConnected(conn *websocket.Conn, msg map[string]interface{}) error {
	log.Printf("[handleImpostorConnected] Handling connected message: %v", msg)
	userID, err := parseUUID(msg, "id")
	if err != nil {
		log.Printf("[handleImpostorConnected] Error parsing user ID: %v", err)
		return err
	}

	user, err := usersRepository.GetUser(userID)
	if err != nil {
		log.Printf("[handleImpostorConnected] Error getting user: %v", err)
		return err
	}

	err = impostorManager.RegisterPlayer(conn, user)
	if err != nil {
		log.Printf("[handleImpostorConnected] Error registering player: %v", err)
		return err
	}

	impostorManager.BroadcastPlayerList()
	return nil
}

// handleImpostorStartGame initializes a new game with specified settings.
func handleImpostorStartGame(msg map[string]interface{}) error {
	log.Println("[handleImpostorStartGame] Starting game")
	dist := parseDistribution(msg)
	category, _ := msg["category"].(string)
	difficulty, _ := msg["difficulty"].(string)
	impostorManager.StartGame(dist, category, difficulty)
	return nil
}

// handleImpostorRemovePlayer removes a player from the game.
func handleImpostorRemovePlayer(msg map[string]interface{}) error {
	log.Println("[handleImpostorRemovePlayer] Removing player")
	userID, err := parseUUID(msg, "id")
	if err != nil {
		log.Printf("[handleImpostorRemovePlayer] Error parsing user ID: %v", err)
		return err
	}
	impostorManager.RemovePlayerByID(userID)
	impostorManager.BroadcastPlayerList()
	return nil
}

// handleImpostorDecideWinner updates the game state based on the winner
func handleImpostorDecideWinner(msg map[string]interface{}) error {
	log.Println("[handleImpostorDecideWinner] Deciding winner")
	impostorWon, _ := msg["impostorWon"].(bool)
	impostorManager.UpdatePoints(impostorWon)
	impostorManager.BroadcastPlayerList()
	impostorManager.ResetGame()
	impostorManager.BroadcastWinner(impostorWon)
	return nil
}

// handleImpostorGetGameStatus returns the current game status for a user.
func handleImpostorGetGameStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("[handleImpostorGetGameStatus] Getting Impostor game status")
	id := r.URL.Query().Get("u")
	userID, err := uuid.Parse(id)

	if err != nil {
		log.Printf("[handleImpostorGetGameStatus] Error parsing user ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := usersRepository.GetUser(userID)

	if err != nil {
		log.Printf("[handleImpostorGetGameStatus] Error getting user: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	gameInfo := impostorManager.GetGameStatus(user.ID)

	respondWithJSON(w, http.StatusOK, gameInfo)
}

// parseDistribution extracts game distribution settings from a message.
func parseDistribution(msg map[string]interface{}) impostor.ImpostorDistribution {
	log.Printf("[parseDistribution] Parsing distribution: %v", msg)
	return impostor.ImpostorDistribution{
		One:   int(msg["one"].(float64)),
		Two:   int(msg["two"].(float64)),
		Three: int(msg["three"].(float64)),
	}
}
