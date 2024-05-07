package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID   uuid.UUID
	Name string
}

func NewUser(name string) (*User, error) {
	return &User{
		ID:   uuid.New(),
		Name: name,
	}, nil
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
