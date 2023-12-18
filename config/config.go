package config

import (
	"fmt"

	log "github.com/aditya3232/latihan-go-ms-one.git/log"
	"github.com/spf13/viper"
)

var CONFIG = load()

type envConfigs struct {
	DEBUG int `mapstructure:"DEBUG"`

	JWT_KEY string `mapstructure:"JWT_KEY"`

	DB_1_HOST    string `mapstructure:"DB_1_HOST"`
	DB_1_PORT    string `mapstructure:"DB_1_PORT"`
	DB_1_USER    string `mapstructure:"DB_1_USER"`
	DB_1_PASS    string `mapstructure:"DB_1_PASS"`
	DB_1_NAME    string `mapstructure:"DB_1_NAME"`
	DB_1_CHARSET string `mapstructure:"DB_1_CHARSET"`
	DB_1_LOC     string `mapstructure:"DB_1_LOC"`

	APP_PORT string `mapstructure:"APP_PORT"`
	API_KEY  string `mapstructure:"API_KEY"`
	APP_HOST string `mapstructure:"APP_HOST"`
}

func load() (config *envConfigs) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath("../latihan-go-ms-one/config")

	// Tell viper the name of your file
	viper.SetConfigName(".env")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Panic(fmt.Sprintf("Error reading env file: %s", err))

	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Panic(fmt.Sprintf("Error loaded env variables: %s", err))

	}

	return
}
