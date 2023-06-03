package main

import (
	"github.com/modaniru/twitch-general-follows/src/internal"
	"github.com/modaniru/twitch-general-follows/src/twitch"
)

func main() {
	config := internal.NewConfiguration("config.toml")
	service := internal.NewService(*twitch.NewQueries(config.ClientId, config.ClientSecret))
	server := internal.NewServer(*service)
	server.Start(config.Port)
}
