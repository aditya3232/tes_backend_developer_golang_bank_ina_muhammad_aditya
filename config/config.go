package config

import (
	"log"

	"github.com/spf13/viper"
)

var CONFIG = load()

type envConfigs struct {
	DEBUG int `mapstructure:"DEBUG"`

	APP_PORT string `mapstructure:"APP_PORT"`
	APP_HOST string `mapstructure:"APP_HOST"`

	GOOGLE_CLIENT_ID     string `mapstructure:"GOOGLE_CLIENT_ID"`
	GOOGLE_CLIENT_SECRET string `mapstructure:"GOOGLE_CLIENT_SECRET"`
	GOOGLE_REDIRECT_URL  string `mapstructure:"GOOGLE_REDIRECT_URL"`
	OAUTH_STATE_STRING   string `mapstructure:"OAUTH_STATE_STRING"`

	DB_HOST    string `mapstructure:"DB_HOST"`
	DB_PORT    string `mapstructure:"DB_PORT"`
	DB_USER    string `mapstructure:"DB_USER"`
	DB_PASS    string `mapstructure:"DB_PASS"`
	DB_NAME    string `mapstructure:"DB_NAME"`
	DB_CHARSET string `mapstructure:"DB_CHARSET"`
	DB_LOC     string `mapstructure:"DB_LOC"`
}

func load() (config *envConfigs) {
	// Tell viper the path/location of your env file. If it is root just add "."
	viper.AddConfigPath("../tes_backend_developer_golang_bank_ina_muhammad_aditya/config")

	// Tell viper the name of your file
	viper.SetConfigName(".env")

	// Tell viper the type of your file
	viper.SetConfigType("env")

	// Viper reads all the variables from env file and log error if any found
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	// Viper unmarshals the loaded env varialbes into the struct
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return
}
