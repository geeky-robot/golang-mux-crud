package main

import (
	"database/sql"
	"flag"
	"fmt"
	"net/http"

	h "github.com/geeky-robot/golang-gin/handler"
	r "github.com/geeky-robot/golang-gin/repository"
	s "github.com/geeky-robot/golang-gin/service"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

var (
	dbName = flag.String("dbname", "postgres", "Database name")
	dbHost = flag.String("dbhost", "127.0.0.1", "Database host")
	dbPort = flag.String("dbport", "3000", "Port for the database")
	dbUser = flag.String("dbuser", "postgres", "Database user")
	dbPass = flag.String("dbpass", "admin", "Database password")
)

func main() {

	router := mux.NewRouter()
	var db *sql.DB = InitDb(*dbHost, *dbPort, *dbUser, *dbPass, *dbName)
	repo := r.UserDbRepo{Db: db}
	service := &s.UserService{Repo: &(repo)}
	handler := h.NewUserHandler(service)
	router.HandleFunc("/user", handler.CreateUser).Methods("POST")
	router.HandleFunc("/users", handler.CreateUsers).Methods("POST")
	router.HandleFunc("/user/{id}", handler.GetUserById).Methods("GET")
	router.HandleFunc("/user", handler.GetUsers).Methods("GET")
	router.HandleFunc("/user", handler.UpdateUser).Methods("PUT")
	router.HandleFunc("/user/{id}", handler.DeleteUser).Methods("DELETE")
	http.ListenAndServe(":5000", router)

}

func InitDb(dbHost, dbPort, dbUser, dbPass, dbName string) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPass, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Database connection error", err.Error())
	}
	return db
}
