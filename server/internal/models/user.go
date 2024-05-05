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
