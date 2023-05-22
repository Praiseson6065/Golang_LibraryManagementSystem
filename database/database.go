package database
import(
	"database/sql"
)
const DbName = "lib"
const DbType = "postgres"
const DbUserName = "lib"
const DbPassword = "lib"
const DbPort = "5432"

func DbConnect() (*sql.DB,error){
	return sql.Open(DbType, "postgres://"+DbUserName+":"+DbPassword+"@localhost:"+DbPort+"/"+DbName)
}
