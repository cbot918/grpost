package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/cbot918/liby/util"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DB_URL = "postgres://postgres:12345@localhost:5432/grpost?sslmode=disable"
)

var log = fmt.Println
var logf = fmt.Printf

func InsertUserObj(users []User) {
	if len(users) == 0 {
		log("no user data, quit program")
		os.Exit(1)
	}

	var stmt string

	// stmt = GetCatStmt(users, "users")
	// log(stmt)
	// insertUser(stmt)

	stmt = GetCatStmt(users, "follow")
	log(stmt)
	insertUser(stmt)

}
func GetCatStmt(users []User, tableType string) (str string) {
	switch tableType {
	case "users":
		{
			str = "INSERT INTO users (id, email, name, password, pic) VALUES"
			for _, user := range users {
				str += fmt.Sprintf("('%s','%s','%s','%s','%s'),", util.GetUuidFill(user.ID.Oid, 32), user.Email, user.Name, user.Password, user.Pic)
			}
			str = str[:len(str)-1]
			str += ";"
		}
	case "follow":
		{
			str = "INSERT INTO follow (from_user,to_user) VALUES"
			for _, user := range users {
				if len(user.Followers) > 0 {
					for _, follower := range user.Followers {
						str += fmt.Sprintf("('%s','%s'),", util.GetUuidFill(user.ID.Oid, 32), util.GetUuidFill(follower.Oid, 32))
					}
				}
				if len(user.Following) > 0 {
					for _, following := range user.Following {
						str += fmt.Sprintf("('%s','%s'),", util.GetUuidFill(following.Oid, 32), util.GetUuidFill(user.ID.Oid, 32))
					}
				}
			}
			str = str[:len(str)-1]
			str += ";"
		}
	case "uuid":
		{
			u := "1476142dea1a4fd6b23b92bb907"
			log(
				util.GetUuidFill(u, 32),
			)

		}
	}
	return
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

// func ReduceDup(t User) string {
// 	set := mapset.NewSet()
// 	for _, item := range t.Followers {
// 		set.Add(item.Oid)
// 	}
// 	g := set.ToSlice()
// 	return g[0].(string)
// }

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
