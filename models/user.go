package models

import (
	"github.com/google/uuid"
)

type User struct {
	Id      uuid.UUID
	Name    string
	DoB     string
	Age     int64
	Address Address
}
