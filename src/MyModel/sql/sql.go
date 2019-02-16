package sql

import (
	"database/sql"
	"errors"
	"fmt"

	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	Ip       string
	Port     string
	Username string
	Password string
	Schema   string
	Type     string
	LogLevel string
	ConnPtr  *sql.DB
}

type Transaction struct {
	Txn      *sql.Tx
	LogLevel string
}

func (obj *Database) Connect() (err error) {
	conn_str := ""
	if obj.Type == "postgres" {
		conn_str = obj.Type + "://" + obj.Username + ":" + obj.Password +
			"@" + obj.Ip + ":" + obj.Port + "/" + obj.Schema + "?sslmode=disable"
	} else if obj.Type == "oci8" {
		conn_str = obj.Username + "/" + obj.Password +
			"@" + obj.Ip + ":" + obj.Port + "/" + obj.Schema
	} else if obj.Type == "mssql" {
		conn_str = fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s;encrypt=disable", obj.Ip, obj.Username, obj.Password, obj.Port, obj.Schema)
	} else {
		err = errors.New("Unsupported DB type")
		return
	}

	obj.ConnPtr, err = sql.Open(obj.Type, conn_str)
	if err != nil {
		log.Println(obj.LogLevel, "Error", err)
		err = errors.New("DB connect fail")
		return
	}
	if obj.Type == "postgres" {
		err = obj.ConnPtr.Ping()
		if err != nil {
			log.Println(obj.LogLevel, "Error", err)
			err = errors.New("DB ping fail")
			return
		}
	}
	return
}

func (obj *Database) DisConnect() (err error) {
	err = obj.ConnPtr.Close()
	if err != nil {
		log.Println(obj.LogLevel, "Error", err)
		err = errors.New("DB Disconnect fail")
		return
	}
	return
} //verify_DB_Connection_fn()

func (obj *Database) Query(query string, args ...interface{}) (row *sql.Rows, err error) {
	if obj.Type == "postgres" {
		err = obj.ConnPtr.Ping()
		if err != nil {
			log.Println(obj.LogLevel, "Error", err)
			err = errors.New("DB ping fail")
			return
		}
	}
	log.Println(obj.LogLevel, "Debug", "Query : ", query)
	log.Println(obj.LogLevel, "Debug", "Query Args : ", args)

	row, err = obj.ConnPtr.Query(query, args...)
	if err != nil {
		log.Println(obj.LogLevel, "Error", err)
		err = errors.New(query + "DB query fail")
		return
	}
	return
}

func (obj *Database) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	if obj.Type == "postgres" {
		err = obj.ConnPtr.Ping()
		if err != nil {
			log.Println("Error", "Error", err)
			err = errors.New("DB ping fail")
			return
		}
	}
	log.Println(obj.LogLevel, "Debug", "Query : ", query)
	log.Println(obj.LogLevel, "Debug", "Query Args : ", args)

	result, err = obj.ConnPtr.Exec(query, args...)

	if err != nil {
		log.Println(obj.LogLevel, "Error", err)
		//		err = errors.New(query + "DB exec fail")
		return
	}
	return
}

func Scan(row *sql.Rows) (cols []string, data [][]string, err error) {
	cols, err = row.Columns()
	if err != nil {
		err = errors.New("DB get columns fail")
		return
	}

	tmp_byte := make([][]byte, len(cols))
	tmp := make([]interface{}, len(cols))
	for i, _ := range tmp_byte {
		tmp[i] = &tmp_byte[i] // Put pointers to each string in the interface slice
	}
	for row.Next() {
		err = row.Scan(tmp...)
		if err != nil {
			err = errors.New("DB row scan fail")
			return
		}
		rawResult := make([]string, len(cols))
		for i, _ := range tmp_byte {
			rawResult[i] = string(tmp_byte[i]) // Put pointers to each string in the interface slice
		}
		data = append(data, rawResult)
	}
	return
}

func Close(row *sql.Rows) (err error) {
	if row != nil {
		row.Close()
	}
	return
}
