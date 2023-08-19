package database

import (
	"database/sql"

	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type databaseConfigs struct {
	DbName     string `mapstructure:"DbName"`
	DbType     string `mapstructure:"DbType"`
	DbUserName string `mapstructure:"DbUserName"`
	DbPassword string `mapstructure:"DbPassword"`
	DbPort     string `mapstructure:"DbPort"`
	DbHost     string
}



func DatabaseConfigs()(databaseConfigs) {
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
