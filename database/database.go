package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseConfigs struct {
	DbName     string `mapstructure:"DbName"`
	DbType     string `mapstructure:"DbType"`
	DbUserName string `mapstructure:"DbUserName"`
	DbPassword string `mapstructure:"DbPassword"`
	DbPort     string `mapstructure:"DbPort"`
	DbHost     string `mapstructure:"DbHost"`
}

func DatabaseConfigs() databaseConfigs {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	var databaseConfig databaseConfigs
	databaseConfig.DbName = os.Getenv("DbName")
	databaseConfig.DbType = os.Getenv("DbType")
	databaseConfig.DbUserName = os.Getenv("DbUserName")
	databaseConfig.DbPassword = os.Getenv("DbPassword")
	databaseConfig.DbPort = os.Getenv("DbPort")
	databaseConfig.DbHost = os.Getenv("DbHost")
	return databaseConfig

}

func DbConnect() (*sql.DB, error) {
	return sql.Open(DatabaseConfigs().DbType, "postgres://"+DatabaseConfigs().DbUserName+":"+DatabaseConfigs().DbPassword+"@"+DatabaseConfigs().DbHost+":"+DatabaseConfigs().DbPort+"/"+DatabaseConfigs().DbName)
}
func DbGormConnect() (*gorm.DB, error) {
	dbURL := "postgres://" + DatabaseConfigs().DbUserName + ":" + DatabaseConfigs().DbPassword + "@" + DatabaseConfigs().DbHost + ":" + DatabaseConfigs().DbPort + "/" + DatabaseConfigs().DbName
	return gorm.Open(postgres.Open(dbURL), &gorm.Config{})
}
