package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "log"
    "time"
)

func DBConnect() *gorm.DB {
    var err error

	dsn := "root:@/finalbtpn?charset=utf8&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Error when connecting to database: %s", err)
	}

	log.Println("Database connected")

	sqlDB, err := db.DB()

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnMaxLifetime(time.Hour)

    return db
}