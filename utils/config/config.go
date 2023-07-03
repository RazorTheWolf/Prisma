package config

import (
	"github.com/spf13/viper"
	"log"
)

// Configurations exports all the configs
type Configurations struct {
	Discord DiscordConfigurations
	Server  ServerConfigurations
}

type DiscordConfigurations struct {
	ClientId     string
	ClientSecret string
	RedirectURI  string
	Scope        string
}
type ServerConfigurations struct {
	Port string
}

var Configuration Configurations

func Config() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&Configuration)
	if err != nil {
		log.Printf("Unable to decode into struct, %v", err)

	}
}
