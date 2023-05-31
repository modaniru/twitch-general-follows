package main

import (
	"github.com/modaniru/twitch-general-follows/src/internal"
)

func main() {
	server := internal.NewServer("config.toml")
	server.Start()
}
