package common

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func Connect() (*sql.DB, error) {
	user := Config.DbUser
	password := Config.DbPassword
	host := Config.DbHost
	port := Config.DbPort
	dbname := Config.DbName
	driver := Config.DbDriver

	if driver == "mysql" {
		URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, port, dbname)
		db, err := sql.Open(driver, URL)

		if err != nil {
			return nil, err
		}

		return db, nil
	} else {
		return nil, errors.New("Driver is not supported")
	}
}