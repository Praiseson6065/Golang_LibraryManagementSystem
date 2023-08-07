package database

import (
	"database/sql"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseConfigs struct {
	DbName     string `mapstructure:"DbName"`
	DbType     string `mapstructure:"DbType"`
	DbUserName string `mapstructure:"DbUserName"`
	DbPassword string `mapstructure:"DbPassword"`
	DbPort     string `mapstructure:"DbPort"`
}

var databaseConfig *databaseConfigs

func InitDatabseConfig() {
	databaseConfig = loadDatabaseVariables()
}
func loadDatabaseVariables() (config *databaseConfigs) {
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
func DbConnect() (*sql.DB, error) {
	return sql.Open(databaseConfig.DbType, "postgres://"+databaseConfig.DbUserName+":"+databaseConfig.DbPassword+"@localhost:"+databaseConfig.DbPort+"/"+databaseConfig.DbName)
}
func DbGormConnect() (*gorm.DB, error) {
	dbURL := "postgres://" + databaseConfig.DbUserName + ":" + databaseConfig.DbPassword + "@localhost:" + databaseConfig.DbPort + "/" + databaseConfig.DbName
	return gorm.Open(postgres.Open(dbURL), &gorm.Config{})
}
