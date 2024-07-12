package database

import (
	"database/sql"
	"fmt"
	"log"
)

func LocalDBConnect() (*sql.DB, error) {
	log.Println("LocalDBConnect")
	connString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", "root", "root", "192.168.2.5", 3306, "sureshkrishnanv")
	db, err := sql.Open("mysql", connString)
	if err != nil {
		log.Println("Open Connection Failed : ", err.Error())
		return db, err
	}
	log.Println("LocalDBConnect-")
	return db, nil
}
