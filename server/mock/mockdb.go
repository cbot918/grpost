package mock

import (
	"fmt"

	"github.com/cbot918/grpost/server/util"
	_ "github.com/lib/pq"
)

func MockDb() {
	fmt.Println("hihi")

	config, err := util.LoadConfig(".", "app", "env")
	if err != nil {
		fmt.Println("failed to load config")
	}
	fmt.Println(config.DSN)

	// connPoll, err := sql.Open("postgres", )

	// db.New()
}

func main() {

	MockDb()
}
