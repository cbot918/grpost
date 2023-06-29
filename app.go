package main

import (
	"context"
	"fmt"

	"github.com/cbot918/grpost/server"
	"github.com/cbot918/grpost/server/util"
)

const (
	port       = ":5010"
	staticPath = "ui/build"
)

type Config struct {
	PORT       string
	StaticPath string
}

func NewConfig() *Config {
	return &Config{
		PORT: port,
	}
}

func main() {
	// load config
	cfg, err := util.LoadConfig(".", "app", "env")
	if err != nil {
		fmt.Println("failed to load config")
	}

	// db setup
	db := util.GetQueryInstance("postgres", cfg.DSN)

	ctx := context.Background()
	user, err := db.ListUsers(ctx)
	if err != nil {
		fmt.Println("query user failed")
		panic(err)
	}
	fmt.Println(user)

	// server listen
	grpost := server.New(cfg.UI_PATH, db)
	grpost.Server.Listen(cfg.PORT)
}
