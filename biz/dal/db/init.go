package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/czhi-bin/mini-tiktok-backend/pkg/constants"
)

var DB *gorm.DB

func Init() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN), &gorm.Config{})
	
	if err != nil {
		panic(err)
	}

}