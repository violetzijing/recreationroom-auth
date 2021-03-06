package util

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"sync"
)

const (
	maxIdleConns = 10
	maxOpenConns = 100
)

var handle *gorm.DB
var once sync.Once

func ORM() *gorm.DB {
	once.Do(func() {
		dsn := "admin:password@(localhost:3306)/recreationroom?charset=utf8mb4"
		db, err := gorm.Open("mysql", dsn)
		if err != nil {
			log.Printf("[util.db] gorm open failed")
			panic(err)
		}
		if err := db.DB().Ping(); err != nil {
			log.Printf("[util.db] Ping database failed")
			panic(err)
		}
		log.Print("[util.db] Database connection established")
		db.DB().SetMaxIdleConns(maxIdleConns)
		db.DB().SetMaxOpenConns(maxOpenConns)
		handle = db
	})
	return handle
}
