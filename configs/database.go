package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	db, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/go-jwt?charset=utf8&parseTime=True"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect database")
	}

	fmt.Println("Database Connected")
	DB = db
}
