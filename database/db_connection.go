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

// table sql 
// CREATE TABLE IF NOT EXISTS `users` (
//     `id` INT AUTO_INCREMENT PRIMARY KEY, 
//     `username` VARCHAR(128) NOT NULL,
//     `email` VARCHAR(128) UNIQUE NOT NULL,
//     `password` VARCHAR(128) NOT NULL,
//     `created_at` DATETIME,
//     `updated_at` DATETIME
// );

// CREATE TABLE IF NOT EXISTS `photo` (
//     `id` INT AUTO_INCREMENT PRIMARY KEY, 
//     `title` VARCHAR(255) NOT NULL,
//     `caption` VARCHAR(255) NOT NULL,
//     `user_id` INT NOT NULL,
//     `photo_url` VARCHAR(255) NOT NULL,
//     FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
// );