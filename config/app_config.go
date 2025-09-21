package config

import (
	"log"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.AddConfigPath("./config")
	viper.SetConfigFile("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err := viper.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

}
