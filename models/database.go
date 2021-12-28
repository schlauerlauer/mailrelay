package models

import (
	// "gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	log "github.com/sirupsen/logrus"
)

var (
	DB *gorm.DB
)

func local(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
	log.Info("Database connected.")
	return db
}

// func sql(config Sqlconfig) gorm.DB {
// 	dsn := config.Username + ":" + config.Password + "@tcp(" + config.Hostname + ":" + config.Port + ")/" + config.Database + "?charset=utf8mb4&parseTime=true&loc=Local&tls=skip-verify"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		PrepareStmt: true,
// 	})
// 	return db
// }

func migrate(db gorm.DB) {
	db.AutoMigrate(
		&Address{},
		&List{},
		&Mail{},
		&Permission{},
	)
	log.Info("Database migration sucessful.")
}

func ConnectDatabase(config DatabaseConfig) {
	if (config.Sqlite == "") {
		log.Fatal("No database path provided.")
	}
	DB = local(config.Sqlite)
	migrate(*DB)
}