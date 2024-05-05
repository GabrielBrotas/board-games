package models

import (
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type User struct {
	ID   uuid.UUID
	Name string
	Conn *websocket.Conn
}

func NewUser(name string) (*User, error) {
	return &User{
		ID:   uuid.New(),
		Name: name,
		Conn: nil,
	}, nil
}

func (u *User) UpdateConnection(conn *websocket.Conn) {
	u.Conn = conn
}

func (u *User) UpdateName(name string) {
	u.Name = name
}

type UserOut struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func (u *User) ToOut() *UserOut {
	return &UserOut{
		ID:   u.ID,
		Name: u.Name,
	}
}
