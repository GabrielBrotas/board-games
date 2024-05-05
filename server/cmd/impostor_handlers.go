package main

import (
	"fmt"
	"log"
	"net/http"

	game "github.com/GabrielBrotas/who-is-the-imposter/internal/games/imposter"
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

func handleGetPlayerList(w http.ResponseWriter, r *http.Request) {
	players := gameManager.GetPlayerList(true)
	playerList := make([]game.PlayerOut, 0, len(players))

	for _, player := range players {
		playerList = append(playerList, *player.ToOut())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"players": %v}`, playerList)))
}

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
		return handleRemovePlayer(msg)
	case "decideWinner":
		return handleDecideWinner(msg)
	case "resetPoints":
		gameManager.ResetPoints()
		gameManager.BroadcastPlayerList()
	default:
		return fmt.Errorf("unknown message type: %v", msg["type"])
	}
	return nil
}

func handleConnected(conn *websocket.Conn, msg map[string]interface{}) error {
	id, ok := msg["id"].(string)
	if !ok {
		return fmt.Errorf("invalid id value")
	}

	userID, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	user, err := usersRepository.GetUser(userID)

	if err != nil {
		return err
	}

	player := gameManager.GetPlayerByID(user.ID)

	if player == nil {
		newPlayer, err := game.NewPlayer(user)
		if err != nil {
			return err
		}
		gameManager.AddPlayer(newPlayer)
		player = newPlayer
	}

	player.User.UpdateConnection(conn)
	gameManager.BroadcastPlayerList()

	return nil
}

func handleStartGame(msg map[string]interface{}) error {
	dist := game.ImposterDistribution{
		One:   int(msg["one"].(float64)),
		Two:   int(msg["two"].(float64)),
		Three: int(msg["three"].(float64)),
	}
	category, _ := msg["category"].(string)
	difficulty, _ := msg["difficulty"].(string)
	gameManager.StartGame(dist, category, difficulty)
	return nil
}

func handleRemovePlayer(msg map[string]interface{}) error {
	log.Printf("Received removePlayer message: %v", msg)
	id, ok := msg["id"].(string)
	if !ok {
		return fmt.Errorf("invalid id value")
	}

	userID, err := uuid.Parse(id)

	if err != nil {
		return err
	}

	log.Printf("Removing player: %s", userID)
	gameManager.RemovePlayerByID(userID)
	gameManager.BroadcastPlayerList()
	return nil
}

func handleDecideWinner(msg map[string]interface{}) error {
	log.Printf("Received decideWinner message: %v", msg)
	impostorWon, ok := msg["impostorWon"].(bool)
	if !ok {
		return fmt.Errorf("invalid impostorWon value")
	}
	gameManager.UpdatePoints(impostorWon)
	gameManager.BroadcastPlayerList()
	gameManager.ResetGame()
	gameManager.BroadcastWinner(impostorWon)
	return nil
}

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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"gameStarted": %v, "word": "%s", "inGame": %v}`, gameInfo.GameStarted, gameInfo.WordOrRole, gameInfo.InGame)))
}
