package repository

import (
	"database/sql"
	"fmt"

	m "github.com/geeky-robot/golang-gin/models"
	"github.com/google/uuid"
)

type UserDbRepo struct {
	Db *sql.DB
}

func NewUserDbRepo() UserDbRepo {
	return UserDbRepo{}
}

func (u *UserDbRepo) CreateUser(user m.User) m.User {
	var aid uuid.UUID
	var id uuid.UUID
	sqlStatement := `INSERT INTO "USER"."ADDRESS"("CITY", "STATE", "COUNTRY") VALUES ($1, $2, $3) RETURNING "ID"`
	err := u.Db.QueryRow(sqlStatement, user.Address.City, user.Address.Country, user.Address.Country).Scan(&aid)
	if err != nil {
		fmt.Println("Something went wrong while inserting into address", err)
	}
	sqlStatement = `INSERT INTO "USER"."USER"("NAME", "DOB", "AGE", "AID") VALUES ($1, $2, $3, $4) RETURNING "ID"`
	err = u.Db.QueryRow(sqlStatement, user.Name, user.DoB, user.Age, aid).Scan(&id)
	if err != nil {
		fmt.Println("Something went wrong while inserting into user", err)
	}
	user.Id = id
	user.Address.Id = aid
	return user
}

func (u *UserDbRepo) CreateUsers(users []m.User) []m.User {
	for i := range users {
		var aid uuid.UUID
		var id uuid.UUID
		sqlStatement := `INSERT INTO "USER"."ADDRESS"("CITY", "STATE", "COUNTRY") VALUES ($1, $2, $3) RETURNING "ID"`
		err := u.Db.QueryRow(sqlStatement, users[i].Address.City, users[i].Address.Country, users[i].Address.Country).Scan(&aid)
		if err != nil {
			fmt.Println("Something went wrong while inserting into address", err)
		}
		sqlStatement = `INSERT INTO "USER"."USER"("NAME", "DOB", "AGE", "AID") VALUES ($1, $2, $3, $4) RETURNING "ID"`
		err = u.Db.QueryRow(sqlStatement, users[i].Name, users[i].DoB, users[i].Age, aid).Scan(&id)
		if err != nil {
			fmt.Println("Something went wrong while inserting into user", err)
		}
		users[i].Id = id
		users[i].Address.Id = aid
		users = append(users, users[i])
	}
	return users
}

func (u *UserDbRepo) UpdateUser(user m.User) m.User {
	var aid uuid.UUID
	sqlStatement := `UPDATE "USER"."USER" SET "NAME" = $1, "DOB" = $2, "AGE" = $3 where "ID" = $4 RETURNING "AID"`
	err := u.Db.QueryRow(sqlStatement, user.Name, user.DoB, user.Age, user.Id).Scan(&aid)
	if err != nil {
		fmt.Println("Something went wrong while updating in user table", err)
		return m.User{}
	}
	user.Address.Id = aid
	sqlStatement = `UPDATE "USER"."ADDRESS" SET "CITY" = $1, "STATE" = $2, "COUNTRY" = $3 where "ID" = $4`
	u.Db.QueryRow(sqlStatement, user.Address.City, user.Address.State, user.Address.Country, aid)
	return user
}

func (u *UserDbRepo) GetUserById(id uuid.UUID) m.User {
	var user m.User
	sqlStatement := `SELECT "ID", "NAME", "DOB", "AGE", "AID" FROM "USER"."USER" WHERE "ID" = $1`
	u.Db.QueryRow(sqlStatement, id).Scan(&user.Id, &user.Name, &user.DoB, &user.Age, &user.Address.Id)
	sqlStatement = `SELECT "CITY", "STATE", "COUNTRY" FROM "USER"."ADDRESS" WHERE "ID" = $1`
	u.Db.QueryRow(sqlStatement, user.Address.Id).Scan(&user.Address.City, &user.Address.State, &user.Address.Country)
	return user
}

func (u *UserDbRepo) GetUsers() []m.User {
	var users []m.User
	sqlStatement := `SELECT A."ID" AS "ID", "NAME", "DOB", "AGE", B."ID" AS "AID", "CITY", "STATE", "COUNTRY" FROM "USER"."USER" A 
	LEFT JOIN "USER"."ADDRESS" B ON A."AID" = B."ID"`
	rows, err := u.Db.Query(sqlStatement)
	if err != nil {
		return []m.User{}
	}
	for i := 0; rows.Next(); i++ {
		var user m.User
		rows.Scan(&user.Id, &user.Name, &user.DoB, &user.Age, &user.Address.Id, &user.Address.City, &user.Address.State, &user.Address.Country)
		users = append(users, user)
	}
	return users
}

func (u *UserDbRepo) DeleteUser(id uuid.UUID) bool {
	var dbId uuid.UUID
	sqlStatement := `DELETE FROM "USER"."USER" WHERE "ID" = $1 RETURNING "ID"`
	u.Db.QueryRow(sqlStatement, id).Scan(&dbId)
	return dbId != uuid.Nil
}
