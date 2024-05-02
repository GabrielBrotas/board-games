package main

import (
	"sync"
)

type User struct {
	Username string
	Points   int
}

type PlayerManager struct {
	users map[string]*User
	lock  sync.RWMutex
}

func NewPlayerManager() *PlayerManager {
	return &PlayerManager{
		users: make(map[string]*User),
	}
}

// AddUser adds a new user or updates an existing one.
func (manager *PlayerManager) AddUser(username string) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	if _, exists := manager.users[username]; exists {
		return
	}
	user := &User{Username: username, Points: 0}
	manager.users[username] = user
}

// UpdatePoints updates the points for a given user.
func (manager *PlayerManager) UpdatePoints(username string, points int) {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	if user, exists := manager.users[username]; exists {
		user.Points = points
	}
}

// GetUserPoints retrieves the points for a user.
func (manager *PlayerManager) GetUserPoints(username string) int {
	manager.lock.RLock()
	defer manager.lock.RUnlock()

	if user, exists := manager.users[username]; exists {
		return user.Points
	}
	return 0 // return 0 if the user does not exist
}

func (manager *PlayerManager) ResetPoints() {
	manager.lock.Lock()
	defer manager.lock.Unlock()

	for _, user := range manager.users {
		user.Points = 0
	}
}
