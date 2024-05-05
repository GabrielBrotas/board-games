package repository

import (
	"fmt"
	"sync"

	"github.com/GabrielBrotas/board-games/internal/models"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// UserRepository is an in-memory storage for User objects.
type UserRepository struct {
	users map[uuid.UUID]*models.User
	lock  sync.RWMutex // To handle concurrent access to the map.
}

// NewUserRepository creates a new instance of UserRepository.
func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make(map[uuid.UUID]*models.User),
	}
}

// AddUser adds a new user to the repository.
func (r *UserRepository) AddUser(user *models.User) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, exists := r.users[user.ID]; exists {
		return fmt.Errorf("user %s already exists", user.ID)
	}

	r.users[user.ID] = user
	return nil
}

// GetUser retrieves a user by ID from the repository.
func (r *UserRepository) GetUser(id uuid.UUID) (*models.User, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	user, exists := r.users[id]
	if !exists {
		return nil, fmt.Errorf("user %s not found", id)
	}

	return user, nil
}

// GetUserByName retrieves a user by name from the repository.
func (r *UserRepository) GetUserByName(name string) *models.User {
	r.lock.RLock()
	defer r.lock.RUnlock()

	for _, user := range r.users {
		if user.Name == name {
			return user
		}
	}

	return nil
}

// UpdateUser updates an existing user's details in the repository.
func (r *UserRepository) UpdateUser(user *models.User) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, exists := r.users[user.ID]; !exists {
		return fmt.Errorf("user %s not found", user.ID)
	}

	r.users[user.ID] = user
	return nil
}

// RemoveUser removes a user from the repository.
func (r *UserRepository) RemoveUser(id uuid.UUID) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if _, exists := r.users[id]; !exists {
		return fmt.Errorf("user %s not found", id)
	}

	delete(r.users, id)
	return nil
}

// GetUsers returns a list of all users in the repository.
func (r *UserRepository) GetUsers() []*models.UserOut {
	r.lock.RLock()
	defer r.lock.RUnlock()

	users := make([]*models.UserOut, 0, len(r.users))
	for _, user := range r.users {
		users = append(users, user.ToOut())
	}

	return users
}

// RemoveUserConn removes the connection from a user.
func (r *UserRepository) RemoveUserConn(Conn *websocket.Conn) {
	r.lock.Lock()
	defer r.lock.Unlock()

	for _, user := range r.users {
		if user.Conn == Conn {
			user.Conn = nil
		}
	}
}
