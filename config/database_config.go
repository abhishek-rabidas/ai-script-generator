package config

import (
	"ai-script-generator/model"
	"fmt"
	"github.com/charmbracelet/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func SetupDatabase() {
	var (
		host   = os.Getenv("DB_HOST")
		user   = os.Getenv("DB_USER")
		pass   = os.Getenv("DB_PASSWORD")
		dbname = os.Getenv("DB_NAME")
		dbport = os.Getenv("DB_PORT")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, dbport, dbname)

	log.Info("Establishing db connection", "dsn", dsn)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db

	migrate()
}

func migrate() {
	err := DB.AutoMigrate(model.Account{}, model.Profile{})
	if err != nil {
		log.Error(err)
	}
}
