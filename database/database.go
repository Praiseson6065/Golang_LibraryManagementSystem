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

const  G_CLIENT_ID = "608859335609-vaqkdeb1n4b5ofrc92m2h9q77udfl013.apps.googleusercontent.com"
const 	G_CLIENT_SECRET = "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"

func DbConnect() (*sql.DB, error) {
	return sql.Open(DbType, "postgres://"+DbUserName+":"+DbPassword+"@localhost:"+DbPort+"/"+DbName)
}
func DbGormConnect() (*gorm.DB, error) {
	dbURL := "postgres://" + DbUserName + ":" + DbPassword + "@localhost:" + DbPort + "/" + DbName
	
	return gorm.Open(postgres.Open(dbURL), &gorm.Config{})
}
