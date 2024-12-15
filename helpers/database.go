package helpers

import (
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupMySQL() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", GetEnv("DB_USER", ""), GetEnv("DB_PASS", ""), GetEnv("DB_HOST", "127.0.0.1"), GetEnv("DB_PORT", "3306"), GetEnv("DB_NAME", ""))

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database", err)
	}

	logrus.Info("Connected to database")

}
