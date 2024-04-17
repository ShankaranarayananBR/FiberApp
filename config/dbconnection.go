package config

import (
	"fmt"
	"log"
	"os"

	"github.com/ShankaranarayananBR/FiberApp/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	godotenv.Load()

	dbhost := os.Getenv("MYSQL_HOST")
	dbuser := os.Getenv("MYSQL_USER")
	dbpassword := os.Getenv("MSQL_PASSWORD")
	dbname := os.Getenv("MYSQL_DBNAME")

	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbname)
	var db, err = gorm.Open(mysql.Open(connection), &gorm.Config{})
	if err != nil {
		log.Printf("Error while connection to DB:%v", err)
	}

	DB = db
	fmt.Println("db connected successfully")

	AutoMigrate(db)

}

func AutoMigrate(connection *gorm.DB) {
	connection.Debug().AutoMigrate(
		&model.Cashier{},
		&model.Category{},
		&model.Payment{},
		&model.Discount{},
		&model.Order{},
		&model.PaymentType{},
	)
}
