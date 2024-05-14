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
	log.Println("[handleGetSpyfallPlayerList] Getting Spyfall player list")
	players := spyfallManager.GetPlayerList(true)
	playerList := make([]spyfall.PlayerOut, 0, len(players))

	for _, player := range players {
		playerList = append(playerList, *player.ToOut())
	}

	log.Printf("[handleGetSpyfallPlayerList] Returning player list: %v", playerList)
	respondWithJSON(w, http.StatusOK, map[string]interface{}{"players": playerList})
}

func handleSpyfallConnections(w http.ResponseWriter, r *http.Request) {
	log.Println("[handleSpyfallConnections] Handling new Spyfall connection")
	conn, err := spyfallUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("[handleSpyfallConnections] Error upgrading WebSocket connection: %v", err)
		return
	}
	defer conn.Close()
	defer handleCloseSpyfallConnection(conn)

	spyfallMessageLoop(conn)
}

func spyfallMessageLoop(conn *websocket.Conn) {
	log.Printf("[spyfallMessageLoop] New Spyfall connection: %v", conn.RemoteAddr())
	for {
		msg := make(map[string]interface{})
		if err := conn.ReadJSON(&msg); err != nil {
			log.Printf("[spyfallMessageLoop] Error reading JSON from Spyfall client %v: %v", conn.RemoteAddr(), err)
			break
		}
		if err := processSpyfallMessage(conn, msg); err != nil {
			log.Printf("[spyfallMessageLoop] Error processing Spyfall message from %v: %v", conn.RemoteAddr(), err)
			break
		}
	}
}

func handleCloseSpyfallConnection(conn *websocket.Conn) {
	log.Printf("[handleCloseSpyfallConnection] Closing Spyfall connection: %v", conn.RemoteAddr())
	spyfallManager.RemoveConnection(conn)
	spyfallManager.BroadcastPlayerList()
}

func processSpyfallMessage(conn *websocket.Conn, msg map[string]interface{}) error {
	log.Printf("[processSpyfallMessage] Processing message: %v", msg)
	switch msg["type"].(string) {
	case "connected":
		return handleSpyfallConnected(conn, msg)
	case "startGame":
		return handleSpyfallStartGame(msg)
	case "resetGame":
		log.Println("[processSpyfallMessage] Resetting game")
		spyfallManager.ResetGame()
	case "removePlayer":
		return handleSpyfallRemovePlayer(msg) // v
	case "decideWinner":
		return handleSpyfallDecideWinner(msg)
	case "resetPoints":
		log.Println("[processSpyfallMessage] Resetting points")
		spyfallManager.ResetPoints()
		spyfallManager.BroadcastPlayerList()
		return nil
	case "showSpiesNumber":
		log.Println("[processSpyfallMessage] Showing spies number")
		spyfallManager.BroadcastSpiesNumber()
		return nil
	default:
		return fmt.Errorf("unknown message type: %v", msg["type"])
	}
	return nil
}

func handleSpyfallConnected(conn *websocket.Conn, msg map[string]interface{}) error {
	log.Printf("[handleSpyfallConnected] Handling connected message: %v", msg)
	userID, err := parseUUID(msg, "id")
	if err != nil {
		log.Printf("[handleSpyfallConnected] Error parsing user ID: %v", err)
		return err
	}

	user, err := usersRepository.GetUser(userID)
	if err != nil {
		log.Printf("[handleSpyfallConnected] Error getting user: %v", err)
		return err
	}

	err = spyfallManager.RegisterPlayer(conn, user)
	if err != nil {
		log.Printf("[handleSpyfallConnected] Error registering player: %v", err)
		return err
	}

	spyfallManager.BroadcastPlayerList()
	return nil
}

// handleStartGame initializes a new game with specified settings.
func handleSpyfallStartGame(msg map[string]interface{}) error {
	log.Println("[handleSpyfallStartGame] Starting game")
	dist := parseSpyfallDistribution(msg)
	spyfallManager.StartGame(dist)
	return nil
}

// handleSpyfallRemovePlayer removes a player from the game.
func handleSpyfallRemovePlayer(msg map[string]interface{}) error {
	log.Println("[handleSpyfallRemovePlayer] Removing player")
	userID, err := parseUUID(msg, "id")
	if err != nil {
		log.Printf("[handleSpyfallRemovePlayer] Error parsing user ID: %v", err)
		return err
	}
	spyfallManager.RemovePlayerByID(userID)
	spyfallManager.BroadcastPlayerList()
	return nil
}

// handleSpyfallDecideWinner updates the game state based on the winner
func handleSpyfallDecideWinner(msg map[string]interface{}) error {
	log.Println("[handleSpyfallDecideWinner] Deciding winner")
	spyWon, _ := msg["spyWon"].(bool)
	spyfallManager.FinishGame(spyWon)
	spyfallManager.BroadcastPlayerList()
	return nil
}

// handleSpyfallGetPlayerGameStatus returns the current game status for a user.
func handleSpyfallGetPlayerGameStatus(w http.ResponseWriter, r *http.Request) {
	log.Println("[handleSpyfallGetPlayerGameStatus] Getting player game status")
	id := r.URL.Query().Get("u")
	userID, err := uuid.Parse(id)

	if err != nil {
		log.Printf("[handleSpyfallGetPlayerGameStatus] Error parsing user ID: %v", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := usersRepository.GetUser(userID)

	if err != nil {
		log.Printf("[handleSpyfallGetPlayerGameStatus] Error getting user: %v", err)
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	gameInfo := spyfallManager.GetPlayerGameStatus(user.ID)

	respondWithJSON(w, http.StatusOK, gameInfo)
}

func parseSpyfallDistribution(msg map[string]interface{}) spyfall.SpiesDistribution {
	log.Printf("[parseSpyfallDistribution] Parsing spies distribution: %v", msg)
	return spyfall.SpiesDistribution{
		One:   int(msg["one"].(float64)),
		Two:   int(msg["two"].(float64)),
		Three: int(msg["three"].(float64)),
	}
}
