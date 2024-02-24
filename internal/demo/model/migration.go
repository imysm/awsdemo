package model

import (
	"gorm.io/gorm"
	"log"
)

func Migrate(db *gorm.DB) {
	var err error

	// 迁移模式，创建表
	log.Printf("create table %s \n", (&Test{}).TableName())
	err = db.AutoMigrate(&Test{})
	if err != nil {
		log.Fatal("failed to create table")
	}
}
