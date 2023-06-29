package util

import (
	"database/sql"
	"fmt"

	sqlcdb "github.com/cbot918/grpost/db/sqlc"
	_ "github.com/lib/pq"
)

func GetQueryInstance(dbtype string, dburl string) *sqlcdb.Queries {
	// db_conn_pool instance
	pool, err := sql.Open(dbtype, dburl)
	if err != nil {
		fmt.Println("sql Open failed")
		panic(err)
	}

	err = pool.Ping()
	if err != nil {
		fmt.Println("db ping failed")
		panic(err)
	}

	// sqlc query instance
	queries := sqlcdb.New(pool)
	return queries
}
