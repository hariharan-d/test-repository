package ssb

import (
	"errors"
	"log"
	"strings"

	"MyModel/db"

	sqlx "database/sql"

	"MyModel/sql"

	"github.com/astaxie/beego"
)

type Contact struct {
	ID      string
	Name    string
	Email   string
	Address string
	Mobile  string
	BookID  string
}

func GetContactDetails(ContactName, Email, id, BookId string) (table []Contact, err error) {
	var row *sqlx.Rows

	row, err = db.Db.Query(`SELECT id, name, email, address, mobile, book_id 
	FROM lookup.contacts 
	where ($1='' OR LOWER(name) like '%' || $1 || '%') AND ($2='' OR LOWER(email) like '%' || $2 || '%') 
	AND ($3='' OR id = $3::integer) AND (book_id::text=$4)`, strings.ToLower(ContactName), strings.ToLower(Email), id, BookId)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get org details")
		return
	}
	defer sql.Close(row)
	_, data, err := sql.Scan(row)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		err = errors.New("Unable to get org details")
		return
	}
	if len(data) <= 0 {
		err = errors.New("No record found.")
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}

	for i := range data {
		var a Contact
		a.ID = data[i][0]
		a.Name = data[i][1]
		a.Email = data[i][2]
		a.Address = data[i][3]
		a.Mobile = data[i][4]
		a.BookID = data[i][5]

		table = append(table, a)
	}
	return
}
