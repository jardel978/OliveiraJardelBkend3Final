package configs

import (
	"OliveiraJardelBkend3Final/internal/migrations"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

var gormDb *gorm.DB

func StartGormDb() {
	cnn, err := gorm.Open(mysql.Open(cfg.DB.getURI()), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	config, _ := cnn.DB()
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
	config.SetConnMaxLifetime(time.Hour)

	err = migrations.RunMigrations(cnn)
	if err != nil {
		log.Fatal(err)
	}

	gormDb = cnn
}

func GetGormDB() *gorm.DB {
	return gormDb
}
