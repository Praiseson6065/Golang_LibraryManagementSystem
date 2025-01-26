package payments

import (
	"fmt"

	"log"

	"gorm.io/gorm"

	_ "LibManMicroServ/config"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
)

var db *gorm.DB

var stripeKey string

func connectDB() {
	host := viper.GetString("DATABASE.HOST")
	port := viper.GetString("DATABASE.PORT")
	password := viper.GetString("DATABASE.PASSWORD")

	dbName := viper.GetString("DBNAME.PAYMENTS")
	user := dbName + viper.GetString("DATABASE.USER")

	dsn := "host=" + host + " user=" + user + " password=" + password
	dsn += " dbname=" + dbName + " port=" + port + " sslmode=disable TimeZone=Asia/Kolkata"

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		log.Fatal("Failed to connect to payments database: ", err)
		panic(err)
	}

	db = database
	stripeKey = viper.GetString("STRIPE.SECRET_KEY ")
	err = db.AutoMigrate(&Payment{}, &Book{})
	if err != nil {
		log.Fatal("Failed to migrate payments database: ", err)
		panic(err)
	}

	fmt.Println("Connected to Payments database")
}

func init() {
	connectDB()

}
