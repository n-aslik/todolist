package database

import (
	"fmt"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conn *gorm.DB

func ConnectDB() error {
	connstr := "host=localhost port=5432 user=postgres password=@@sl8998 dbname=test1 sslmode=disable"
	db, err := gorm.Open(postgres.Open(connstr), &gorm.Config{})
	if err != nil {
		return err
	}
	fmt.Println("Connected to db")
	conn = db
	return nil
}

func CloseDB(db *gorm.DB) error {
	// err := conn.
	// if err != nil {
	// 	return err
	// }
	return nil

}
func GetconnectDB() *gorm.DB {
	return conn
}
