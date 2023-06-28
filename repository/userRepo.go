package repository

import (
	m "github.com/geeky-robot/golang-gin/models"
	"github.com/google/uuid"
)

type UserRepo struct {
}

var Users []m.User

func NewUserRepo() UserRepo {
	return UserRepo{}
}

func (u *UserRepo) CreateUser(user m.User) m.User {
	user.Id = uuid.New()
	user.Address.Id = uuid.New()
	Users = append(Users, user)
	return user
}

func (u *UserRepo) CreateUsers(users []m.User) []m.User {
	for i := range users {
		users[i].Id = uuid.New()
		users[i].Address.Id = uuid.New()
		Users = append(Users, users[i])
	}
	return Users
}

func (u *UserRepo) UpdateUser(user m.User) m.User {
	for i := range Users {
		if Users[i].Id == user.Id {
			Users[i] = user
			return user
		}
	}
	return m.User{}
}

func (u *UserRepo) GetUserById(id uuid.UUID) m.User {
	for i := range Users {
		if Users[i].Id == id {
			return Users[i]
		}
	}
	return m.User{}
}

func (u *UserRepo) GetUsers() []m.User {
	return Users
}

func (u *UserRepo) DeleteUser(id uuid.UUID) bool {
	for i := range Users {
		if Users[i].Id == id {
			Users = append(Users[:i], Users[i+1:]...)
			return true
		}
	}
	return false
}
