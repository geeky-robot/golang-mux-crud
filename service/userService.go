package service

import (
	"fmt"
	"time"

	m "github.com/geeky-robot/golang-gin/models"
	r "github.com/geeky-robot/golang-gin/repository"
	"github.com/google/uuid"
)

type UserService struct {
	Repo r.Operations
}

func NewUserRepo(repo r.Operations) UserService {
	return UserService{Repo: repo}
}

func (s *UserService) CreateUser(user m.User) m.User {
	if user.DoB != "" {
		dob, err := time.Parse("2006-01-02", user.DoB)
		if err != nil {
			fmt.Println(err)
			return m.User{}
		}
		user.Age = int64((time.Since(dob).Hours()) / (24 * 365))
	}
	return (s.Repo).CreateUser(user)
}

func (s *UserService) CreateUsers(users []m.User) []m.User {
	for i := range users {
		if users[i].DoB != "" {
			dob, err := time.Parse("2006-01-02", users[i].DoB)
			if err != nil {
				fmt.Println(err)
			} else {
				users[i].Age = int64((time.Since(dob).Hours()) / (24 * 365))
			}
		}
	}
	return (s.Repo).CreateUsers(users)
}

func (s *UserService) UpdateUser(user m.User) m.User {
	if user.DoB != "" {
		dob, err := time.Parse("2006-01-02", user.DoB)
		if err != nil {
			fmt.Println(err)
			return m.User{}
		}
		user.Age = int64((time.Since(dob).Hours()) / (24 * 365))
	}
	return (s.Repo).UpdateUser(user)
}

func (s *UserService) GetUserById(id uuid.UUID) m.User {
	return (s.Repo).GetUserById(id)
}

func (s *UserService) GetUsers() []m.User {
	return (s.Repo).GetUsers()
}

func (s *UserService) DeleteUser(id uuid.UUID) bool {
	return (s.Repo).DeleteUser(id)
}
