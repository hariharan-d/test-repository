package db

import (
	"log"

	"MyModel/sql"

	"github.com/astaxie/beego"
)

var Db sql.Database

func Init() (err error) {
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Trying to connect DB")
	Db.Ip = beego.AppConfig.String("DBIP")
	Db.Port = beego.AppConfig.String("DBPort")
	Db.Type = beego.AppConfig.String("DBType")
	Db.Schema = beego.AppConfig.String("DBName")
	Db.Username = beego.AppConfig.String("DBUsername")
	Db.Password = beego.AppConfig.String("DBPassword")

	err = Db.Connect()
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", "DB Connect fail")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Info", "DB Connected successfully")
	return
}
