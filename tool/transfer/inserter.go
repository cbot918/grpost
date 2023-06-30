package main

import (
	"database/sql"
	"fmt"
	"os"

	mapset "github.com/deckarep/golang-set"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DB_URL = "postgres://postgres:12345@localhost:5432/grpost?sslmode=disable"
)

var log = fmt.Println
var logf = fmt.Printf

func InsertUsers(users []User) {
	if len(users) == 0 {
		log("no user data, quit program")
		os.Exit(1)
	}

	stmt := GetCatStmt(users)
	log(stmt)
	insertUser(stmt)

}
func GetCatStmt(users []User) string {
	str := "INSERT INTO users (email, name, password, pic) VALUES"
	for _, item := range users {
		str += fmt.Sprintf("('%s','%s','%s','%s'),", item.Email, item.Name, item.Password, item.Pic)
	}
	str = str[:len(str)-1]
	str += ";"
	return str
}

func insertUser(stmt string) sql.Result {
	db := setupDb(DB_URL)
	err := db.Ping()
	if err != nil {
		fmt.Println("ping failed")
	}

	// stage1
	res, err := db.Exec(stmt)
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
