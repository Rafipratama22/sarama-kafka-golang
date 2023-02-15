package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"example-goka/app/shared/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type MySQL struct {
	DB  *gorm.DB
	Sql *sql.DB
}

func SetupMySQL() *MySQL {
	fmt.Println("Setup... MySQL")

	url := config.DB.MySQL.Datasourcename

	log.Println(url, "url")
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	db, err := gorm.Open(mysql.Open(url), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	sql, err := db.DB()
	if err != nil {
		panic(err)
	}

	sql.SetConnMaxLifetime(5 * time.Minute)
	sql.SetConnMaxIdleTime(5 * time.Minute)
	sql.SetMaxIdleConns(40)
	sql.SetMaxOpenConns(40)

	return &MySQL{
		db,
		sql,
	}
}
