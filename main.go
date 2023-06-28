package main

import (
	"github.com/cbot918/grpost/internal"
)

const (
	port = ":5010"
)

type Config struct {
	PORT string
}

func NewConfig() *Config {
	return &Config{
		PORT: port,
	}
}

func main() {

	cfg := NewConfig()

	app := internal.New(cfg.PORT, "ui/build")

	app.Run()
}
