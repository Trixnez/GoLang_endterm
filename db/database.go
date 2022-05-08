package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	connStr := "host = 127.0.0.1 port = 5432 user = adil password = adil dbname = market sslmode = disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic("Не удается подключится к базе")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Card{})
	db.AutoMigrate(&Type{})
	db.AutoMigrate(&Product{})
	db.AutoMigrate(&Shop{})
	DB = db
}
