package repository

import (
	m "github.com/geeky-robot/golang-gin/models"
	"github.com/google/uuid"
)

type Operations interface {
	CreateUser(user m.User) m.User
	CreateUsers(users []m.User) []m.User
	UpdateUser(user m.User) m.User
	GetUserById(id uuid.UUID) m.User
	GetUsers() []m.User
	DeleteUser(id uuid.UUID) bool
}
