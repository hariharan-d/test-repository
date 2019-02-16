package user

import (
	"errors"
	"log"

	"MyModel/db"

	"MyModel/sql"

	"github.com/astaxie/beego"
)

func Authenticate(uname, pass string) (id string, err error) {
	row, err := db.Db.Query("select id, password from lookup.\"user\" where username=$1 limit 1", uname)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to authenticate user")
		return
	}
	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to authenticate user")
		return
	}
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "Query Data - ", data, "\nData len - ", len(data))
	if len(data) <= 0 {
		err = errors.New("User not registered")
		return
	}

	if pass != data[0][1] {
		err = errors.New("User password incorrect")
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "data[0][3]", data[0][3])
		return
	} else {
		id = data[0][0]
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "Login SUccess")
	}

	return
}
