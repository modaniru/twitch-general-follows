package main

import (
	"github.com/modaniru/twitch-general-follows/src/internal"
	"github.com/modaniru/twitch-general-follows/src/twitch"
	"github.com/spf13/viper"
	"log"
)

func main() {
	// Init config.yaml
	viper.AddConfigPath("./")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
	service := internal.NewService(*twitch.NewQueries(viper.GetString("client_id"), viper.GetString("client_secret")))
	server := internal.NewServer(*service)
	server.Start(viper.GetInt("port"))
}
