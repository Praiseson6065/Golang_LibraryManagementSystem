package config

import (
	
	"os"
	"log"
	"github.com/joho/godotenv"
)

type envConfigs struct {
	SecretKey       string `mapstructure:"SECRET_KEY"`
	G_CLIENT_ID     string `mapstructure:"G_CLIENT_ID"`
	G_CLIENT_SECRET string `mapstructure:"G_CLIENT_SECRET"`
	G_REDIRECT      string `mapstructure:"G_REDIRECT"`
	STRIPE_key      string `mapstructure:"STRIPE_key"`
	STRIPE_P        string `mapstructure:"STRIPE_P"`
}



func EnvConfigs() (envConfigs){
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var EnvConfig envConfigs
	EnvConfig.SecretKey = os.Getenv("SecretKey")
	EnvConfig.G_CLIENT_ID = os.Getenv("G_CLIENT_ID")
	EnvConfig.G_CLIENT_SECRET = os.Getenv("G_CLIENT_SECRET")
	EnvConfig.STRIPE_key = os.Getenv("STRIPE_key")
	EnvConfig.STRIPE_P = os.Getenv("STRIPE_P")
	
	return EnvConfig

}
