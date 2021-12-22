package model

import (
	"LinkedTo-go/conf"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

type Model interface {
	DataStatistic | Map | MapStatistic | News | Policy | Specialist | User
}

func init() {
	connectDatabase()
	migrate()
}

func connectDatabase() {
	dbArgs := conf.LoginInfo["username"] + ":" + conf.LoginInfo["password"] +
		"@(localhost)/" + conf.LoginInfo["db_name"] + "?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dbArgs), &gorm.Config{})
	if err != nil {
		logrus.Panic(err)
	}
}

func migrate() {
	err := DB.AutoMigrate(&DataStatistic{}, &Map{}, &MapStatistic{}, &News{}, &Policy{}, &Specialist{}, &User{})
	if err != nil {
		log.Fatal(err)
	}
}
