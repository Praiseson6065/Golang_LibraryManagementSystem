package config

import (
	"log"

	"github.com/spf13/viper"
)

type envConfigs struct {
	SecretKey       string `mapstructure:"SECRET_KEY"`
	G_CLIENT_ID     string `mapstructure:"G_CLIENT_ID"`
	G_CLIENT_SECRET string `mapstructure:"G_CLIENT_SECRET"`
	G_REDIRECT      string `mapstructure:"G_REDIRECT "`
	STRIPE_key      string `mapstructure:"STRIPE_key"`
	STRIPE_P        string `mapstructure:"STRIPE_P"`
}

var EnvConfigs *envConfigs

func InitEnvConfigs() {
	EnvConfigs = loadEnvVariables()
}
func loadEnvVariables() (config *envConfigs) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}
	return
}
