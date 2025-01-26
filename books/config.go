package books

import (
	"fmt"

	"log"

	"gorm.io/gorm"

	_ "LibManMicroServ/config"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
)

var db *gorm.DB

func connectDB() {
	host := viper.GetString("DATABASE.HOST")
	port := viper.GetString("DATABASE.PORT")
	password := viper.GetString("DATABASE.PASSWORD")

	dbName := viper.GetString("DBNAME.BOOKS")
	user := dbName + viper.GetString("DATABASE.USER")

	dsn := "host=" + host + " user=" + user + " password=" + password
	dsn += " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Kolkata"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal("Failed to connect to Books database: ", err)
		panic(err)
	}

	db = database

	err = db.AutoMigrate(&Book{})
	if err != nil {
		log.Fatal("Failed to migrate Books database: ", err)
		panic(err)
	}

	fmt.Println("Connected to Books database")
}

func init() {
	connectDB()
}
