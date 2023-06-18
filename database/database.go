package database

import (
	"database/sql"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const DbName = "lib"
const DbType = "postgres"
const DbUserName = "lib"
const DbPassword = "lib"
const DbPort = "5432"

func DbConnect() (*sql.DB, error) {
	return sql.Open(DbType, "postgres://"+DbUserName+":"+DbPassword+"@localhost:"+DbPort+"/"+DbName)
}
func DbGormConnect() (*gorm.DB, error) {
	dbURL := "postgres://" + DbUserName + ":" + DbPassword + "@localhost:" + DbPort + "/" + DbName
	
	return gorm.Open(postgres.Open(dbURL), &gorm.Config{})
}
