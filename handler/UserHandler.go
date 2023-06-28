package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	m "github.com/geeky-robot/golang-gin/models"
	s "github.com/geeky-robot/golang-gin/service"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type UserHandler struct {
	UserService *s.UserService
}

func NewUserHandler(userService *s.UserService) UserHandler {
	return UserHandler{userService}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user m.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Something went wrong while decoding request with error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user = h.UserService.CreateUser(user)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println("Something went wrong while encoding response with error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) CreateUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var users []m.User
	err := json.NewDecoder(r.Body).Decode(&users)
	if err != nil {
		fmt.Println("Something went wrong while decoding request with error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	users = h.UserService.CreateUsers(users)
	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		fmt.Println("Something went wrong while encoding response with error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user m.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println("Something went wrong while decoding request with error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	user = h.UserService.CreateUser(user)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println("Something went wrong while encoding response with error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}

func (h *UserHandler) GetUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user m.User
	id := mux.Vars(r)["id"]
	user = h.UserService.GetUserById(uuid.MustParse(id))
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println("Something went wrong while encoding response with error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := h.UserService.GetUsers()
	err := json.NewEncoder(w).Encode(user)
	if err != nil {
		fmt.Println("Something went wrong while encoding response with error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := mux.Vars(r)["id"]
	isDeleted := h.UserService.DeleteUser(uuid.MustParse(id))
	if !isDeleted {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	err := json.NewEncoder(w).Encode("Success")
	if err != nil {
		fmt.Println("Something went wrong while encoding response with error", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
