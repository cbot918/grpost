package main

import (
	"database/sql"
	"fmt"

	mapset "github.com/deckarep/golang-set"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DB_URL = "postgres://postgres:12345@localhost:5432/grpost?sslmode=disable"
)

var log = fmt.Println
var logf = fmt.Printf

func InsertUsers(u *User) sql.Result {
	db := setupDb(DB_URL)
	err := db.Ping()
	if err != nil {
		fmt.Println("ping failed")
	}

	var GetMulInsStmt = func() string {
		var temp string

		return temp
	}
	GetMulInsStmt()

	// stage1
	// "INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", "Jason", "Moiron", "jmoiron@jmoiron.net"
	stmt := "INSERT INTO users (email, name, password, pic) VALUES($1, $2, $3, $4)"
	res, err := db.Exec(stmt, u.Email, u.Name, u.Password, u.Pic)
	if err != nil {
		log("insert exec failed")
		panic(err)
	}
	log("insert stage1 success")

	// // stage2
	// if HaveFollower(u) {
	// 	followers := ReduceDup(u)
	// 	log(followers)
	// 	stmt = "INSERT INTO follow (from_user,to_user) VALUES ($1,$2)"
	// 	_, err = db.Exec(stmt, u.ID, followers)
	// 	if err != nil {
	// 		log("insert follow table failed")
	// 	}
	// }

	// PrintJson(u)
	return res
}

func ReduceDup(t *User) string {
	set := mapset.NewSet()
	for _, item := range t.Followers {
		set.Add(item.Oid)
	}
	g := set.ToSlice()
	return g[0].(string)
}

func HaveFollower(target *User) bool {
	return len(target.Followers) > 0
}

func setupDb(dburl string) *sqlx.DB {
	db, err := sqlx.Connect("postgres", dburl)
	if err != nil {
		log("sqlx connect failed")
		panic(err)
	}
	return db
}
